package download

import (
"inc/lib"
  "fmt"
  "net/http"
   "encoding/json"
  "io/ioutil"
  "strings"
  "os"
  "bytes"
	"regexp"
  "os/exec"
)

func init() {
	lib.NewCommands(
    &lib.ICommand{
		Name:     "(tt|tiktok|tiktoknowm)",
		As:       []string{"tiktok"},
		Tags:     "downloader",
		IsPrefix: true,
		IsQuerry: true,
		IsWaitt:  true,
		Exec: func(client *lib.Event, m *lib.IMessage) {

      
      if !strings.Contains(m.Querry, "tiktok") {
          m.Reply("Itu bukan link tiktok")
        return
      }  


type Stats struct {
	LikeCount    string `json:"likeCount"`
	ShareCount   int    `json:"shareCount"`
	PlayCount    string `json:"playCount"`
}

type Video struct {
	NoWatermark string `json:"noWatermark"`
	Watermark   string `json:"watermark"`
	Cover       string `json:"cover"`
	DynamicCover string `json:"dynamic_cover"`
	OriginCover string `json:"origin_cover"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	Duration    int    `json:"duration"`
	Ratio       string `json:"ratio"`
}
      
type Image struct {
  URL    string `json:"url"`
  Width  int    `json:"width"`
  Height int    `json:"height"`
}

type Music struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	CoverHD     string `json:"cover_hd"`
	CoverLarge  string `json:"cover_large"`
	CoverMedium string `json:"cover_medium"`
	CoverThumb  string `json:"cover_thumb"`
	Duration    int    `json:"duration"`
	PlayURL     string `json:"play_url"`
}

type Author struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	UniqueID    string `json:"unique_id"`
	Signature   string `json:"signature"`
	Avatar      string `json:"avatar"`
	AvatarThumb string `json:"avatar_thumb"`
}

type TikTokVideo struct {
	ID            int64  `json:"id"`
	Title         string `json:"title"`
	URL           string `json:"url"`
	CreatedAt     string `json:"created_at"`
	Stats         Stats  `json:"stats"`
	Video         Video  `json:"video"`
  Image        []Image `json:"images"`
	Music         Music  `json:"music"`
	Author        Author `json:"author"`
}
        regex := regexp.MustCompile(`(https?:\/\/[^\s]+)`)
 newLink := regex.FindStringSubmatch(m.Querry) 


	resp, err := http.Get("https://api.tiklydown.eu.org/api/download?url="+newLink[0])
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var data TikTokVideo
	err = json.Unmarshal(body, &data)
	if err != nil {
    fmt.Println(err)
    }

      
      //IMAGES
      for i, Images := range data.Image {
        fmt.Printf("Image %d:\n", i+1)

        resp, err := http.Get(Images.URL)
        if err != nil {
          fmt.Println("Error:", err)
          return
        }
        defer resp.Body.Close()

        data, err := ioutil.ReadAll(resp.Body)
      if err != nil {
        fmt.Println("Error reading response body:", err)
        return
      }
  
        randomPng := "./" + lib.GetRandomString(5) + ".png"
        randomWebp := "./" + lib.GetRandomString(5) + ".webp"

        if err := os.WriteFile(randomWebp, data, 0600); err != nil {
            fmt.Printf("Failed to save image: %v", err)
            return
        }

          // Run ffmpeg command
          cmd := exec.Command("ffmpeg", "-i", randomWebp, randomPng)
          var out bytes.Buffer
          var stderr bytes.Buffer
          cmd.Stdout = &out
          cmd.Stderr = &stderr
          err = cmd.Run()

          // Check error
          if err != nil {
            fmt.Println("Error:", err)
            return
          }

      
        url, err := lib.UploadV2(randomPng)
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
          bytes, err := client.GetBytes(url)
          if err != nil {
             fmt.Println("Error:", err)
            return
          }
          client.SendImage(m.From, bytes, "", m.ID)
        os.Remove(randomPng)
        os.Remove(randomWebp)
      }



      //VIDEOS
    if len(data.Video.NoWatermark) > 0 {
      bytes, err := client.GetBytes(data.Video.NoWatermark)
      if err != nil {
        m.Reply(err.Error())
        return
      }
      client.SendVideo(m.From, bytes, " ", m.ID)
    } 
		},
	})
}

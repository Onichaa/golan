package download

import (
	"inc/lib"
  "fmt"
  "net/http"
   "encoding/json"
 // "net/url"
 // "strconv"
  "io/ioutil"
 // "time"
  "strings"
 // "os"
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
	CommentCount string `json:"commentCount"`
	ShareCount   int    `json:"shareCount"`
	PlayCount    string `json:"playCount"`
	SaveCount    string `json:"saveCount"`
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
	Music         Music  `json:"music"`
	Author        Author `json:"author"`
}
        
      
	tikTokURL := "https://api.tiklydown.eu.org/api/download?url="+m.Querry

	resp, err := http.Get(tikTokURL)
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

	var tikTokVideo TikTokVideo
	err = json.Unmarshal(body, &tikTokVideo)
	if err != nil {
    fmt.Println(err)
    }
      //fmt.Println(tikTokVideo.Video.NoWatermark)
      bytes, err := client.GetBytes(tikTokVideo.Video.NoWatermark)
			if err != nil {
				m.Reply(err.Error())
				return
			}
			client.SendVideo(m.From, bytes, " ", m.ID)
      /*
      
      type TikTokData struct {
        Creator       string `json:"creator"`
        Code          int    `json:"code"`
        Msg           string `json:"msg"`
        ProcessedTime float64 `json:"processed_time"`
        Data          struct {
          ID              string `json:"id"`
          Region          string `json:"region"`
          Title           string `json:"title"`
          Cover           string `json:"cover"`
          OriginCover     string `json:"origin_cover"`
          Duration        int    `json:"duration"`
          Play            string `json:"play"`
          WmPlay          string `json:"wmplay"`
          HdPlay          string `json:"hdplay"`
          Size            int    `json:"size"`
          WmSize          int    `json:"wm_size"`
          HdSize          int    `json:"hd_size"`
          Music           string `json:"music"`
          MusicInfo       struct {
            ID       string `json:"id"`
            Title    string `json:"title"`
            Play     string `json:"play"`
            Cover    string `json:"cover"`
            Author   string `json:"author"`
            Original bool   `json:"original"`
            Duration int    `json:"duration"`
            Album    string `json:"album"`
          } `json:"music_info"`
          PlayCount     int `json:"play_count"`
          DiggCount     int `json:"digg_count"`
          CommentCount  int `json:"comment_count"`
          ShareCount    int `json:"share_count"`
          DownloadCount int `json:"download_count"`
          CollectCount  int `json:"collect_count"`
          CreateTime    int `json:"create_time"`

          Author              struct {
            ID        string `json:"id"`
            UniqueID  string `json:"unique_id"`
            Nickname  string `json:"nickname"`
            Avatar    string `json:"avatar"`
          } `json:"author"`
          Images    []string `json:"images"`
        } `json:"data"`
      }

        url := "https://skizo.tech/api/tiktok?url="+url.QueryEscape(m.Querry)+"&apikey="+os.Getenv("KEY")

      response, err := http.Get(url)
      if err != nil {
        fmt.Println("Error:", err)
        return
      }
      defer response.Body.Close()
    

      body, err := ioutil.ReadAll(response.Body)
      if err != nil {
        fmt.Println("Error:", err)
        return
      }
      

      var tiktokData TikTokData
      err = json.Unmarshal(body, &tiktokData)
      if err != nil {
        fmt.Println("Error:", err)
        return
      }
      

      if tiktokData.Data.Duration == 0 {
        for _, i := range tiktokData.Data.Images {
          lib.Sleep(2 * time.Second)

          bytes, err := client.GetBytes(i)
          if err != nil {
            m.Reply(err.Error())
            return
          }
          client.SendImage(m.From, bytes, "nih", m.ID) 
        }
        
      } else { 
      
          teks := `*TIKTOK NO WATERMARK*

𖦹 *ID:* ` + tiktokData.Data.ID + `
𖦹 *Author:* ` + tiktokData.Data.Author.UniqueID + `
𖦹 *Region:* ` + tiktokData.Data.Region + `
𖦹 *Judul:* ` + tiktokData.Data.Title + `
𖦹 *Durasi:* ` + strconv.Itoa(tiktokData.Data.Duration) + `
𖦹 *Music:* ` + tiktokData.Data.Music + `
𖦹 *Info Musik:*
  - *Judul:* ` + tiktokData.Data.MusicInfo.Title + `
  - *Author:* ` + tiktokData.Data.MusicInfo.Author + `
𖦹 *Jumlah Komentar:* ` + strconv.Itoa(tiktokData.Data.CommentCount) + `
𖦹 *Jumlah Share:* ` + strconv.Itoa(tiktokData.Data.ShareCount) + `
𖦹 *Didownload:* ` + strconv.Itoa(tiktokData.Data.DownloadCount) + ` kali`

			bytes, err := client.GetBytes(tiktokData.Data.Play)
			if err != nil {
				m.Reply(err.Error())
				return
			}
			client.SendVideo(m.From, bytes, teks, m.ID)
        }
        */
		},
	})
}

package download

import (
  "inc/lib"
  "net/http"
  "encoding/json"
 // "net/url"
  "fmt"
  "strings"
  	"io/ioutil"
	"os"
)

func init() {
  lib.NewCommands(&lib.ICommand{
    Name:     "(ig|instagram)",
    As:       []string{"instagram"},
    Tags:     "downloader",
    IsPrefix: true,
    IsQuerry: true,
    IsWaitt:  true,
    Exec: func(client *lib.Event, m *lib.IMessage) {

      if !strings.Contains(m.Querry, "instagram") {
          m.Reply("Itu bukan link instagram")
        return
      }
    

type Result struct {
	Status int    `json:"status"`
	Media  []string `json:"media"`
}

type ApiResponse struct {
	Status   int      `json:"status"`
	Creator  string   `json:"creator"`
	Response Result `json:"result"`
}


	apiUrl := "https://api.arifzyn.tech/download/instagram?url="+m.Querry+"&apikey="+os.Getenv("KEY")
	resp, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var apiResponse ApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, mediaUrl := range apiResponse.Response.Media {
		fmt.Println("Media URL:", mediaUrl)
    
    bytes, err := client.GetBytes(mediaUrl)
          if err != nil {
             fmt.Println("Error:", err)
            return
          }
          client.SendVideo(m.From, bytes, "", m.ID)
    client.SendImage(m.From, bytes, "", m.ID)
    
             
	}


	

      
      result, err := lib.Instagram(m.Querry)
      if err != nil {
        fmt.Println("Error:", err)
        return
      }

      fmt.Println(result)

      types := []string{}
      image := []string{}
      urls := []string{}

      for _, value := range result {
          typess := value["type"]
          types = append(types, typess)

          if typess == "video" {
              url := value["url"]
              urls = append(urls, url)
            bytes, err := client.GetBytes(urls[0])
            if err != nil {
              m.Reply(err.Error())
              return
            }
            client.SendVideo(m.From, bytes, "ini", m.ID)
          } else if typess == "image" {
              img := value["url"]
              image = append(image, img)
            bytes, err := client.GetBytes(image[0])
            if err != nil {
              m.Reply(err.Error())
              return
            }
            client.SendImage(m.From, bytes, "ini", m.ID)
          }
      }

      /*      
      resp, err := http.Get("https://skizo.tech/api/igdl?url="+url.QueryEscape(m.Querry)+"&apikey=batu")

      if strings.Contains(m.Querry, "https://www.instagram.com/reel/") {
      type respon struct {
        Caption string   `json:"caption"`
        Media   []string `json:"media"`
      }
      if err != nil {
          fmt.Println("Error:", err)
          return
        }
        defer resp.Body.Close()
        var data respon
        err = json.NewDecoder(resp.Body).Decode(&data)
        if err != nil {
          fmt.Println("Error:", err)
          return
        }
        // Mengambil media
          caption := data.Caption
        media := data.Media
        for _, url := range media {
         
          bytes, err := client.GetBytes(url)
          if err != nil {
             fmt.Println("Error:", err)
            return
          }
          client.SendVideo(m.From, bytes, caption, m.ID)
             
        }

       } else if strings.Contains(m.Querry, "https://www.instagram.com/p/") {
        type respon struct {
        Caption string   `json:"caption"`
        Media   []string `json:"media"`
      }
      if err != nil {
          fmt.Println("Error:", err)
          return
        }
        defer resp.Body.Close()
        var data respon
        err = json.NewDecoder(resp.Body).Decode(&data)
        if err != nil {
          m.Reply(err.Error())
          return
        }
        // Mengambil media
          caption := data.Caption
        media := data.Media
        for _, ur := range media {
          bytes, err := client.GetBytes(ur)
          if err != nil {
            m.Reply(err.Error())
            return
          }
          client.SendVideo(m.From, bytes, caption, m.ID)
           client.SendImage(m.From, bytes, caption, m.ID)
             
        }   

      } else if strings.Contains(m.Querry, "https://www.instagram.com/stories/") {
        type respon struct {
          Caption string   `json:"caption"`
          Media   []string `json:"media"`
        }
        if err != nil {
            fmt.Println("Error:", err)
            return
          }
          defer resp.Body.Close()
          var data respon
          err = json.NewDecoder(resp.Body).Decode(&data)
          if err != nil {
            m.Reply(err.Error())
            return
          }
          // Mengambil media
            caption := data.Caption
          media := data.Media
          for _, ur := range media {
            bytes, err := client.GetBytes(ur)
            if err != nil {
              m.Reply(err.Error())
              return
            }
            client.SendVideo(m.From, bytes, caption, m.ID)
             client.SendImage(m.From, bytes, caption, m.ID)

          }   
          
      }
      */
    },
  })
}

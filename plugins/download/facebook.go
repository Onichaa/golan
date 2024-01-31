package download

import (
  "inc/lib"
  "fmt"
  //"net/http"
   //"encoding/json"
  //"io/ioutil"
  "strings"
)

func init() {
  lib.NewCommands(
    &lib.ICommand{
    Name:     "(fb|facebook)",
    As:       []string{"facebook"},
    Tags:     "downloader",
    IsPrefix: true,
    IsQuerry: true,
    IsWaitt:  true,
    Exec: func(client *lib.Event, m *lib.IMessage) {

      if !strings.Contains(m.Querry, "facebook") {
          m.Reply("Itu bukan link facebook")
        return
      }

      
      resp, err := lib.Fb(m.Querry)
      if err != nil {
        fmt.Println("Error:", err)
        return
      }
      
      if resp != nil {
        bytes, err := client.GetBytes(resp.HD)
        if err != nil {
          fmt.Println("Error:", err)
          return
        }
        client.SendVideo(m.From, bytes, resp.Desc, m.ID)
      }
      
      /*
      
      type Media struct {
        URL        string `json:"url"`
        Quality    string `json:"quality"`
        Extension string `json:"extension"`
        Size       int    `json:"size"`
        FormattedSize string `json:"formattedSize"`
        VideoAvailable bool `json:"videoAvailable"`
        AudioAvailable bool `json:"audioAvailable"`
        Chunked      bool `json:"chunked"`
        Cached       bool `json:"cached"`
      }

      type Response struct {
        Creator   string   `json:"creator"`
        URL       string   `json:"url"`
        Title     string   `json:"title"`
        Thumbnail string   `json:"thumbnail"`
        Duration int      `json:"duration"`
        Source    string   `json:"source"`
        Medias    []Media `json:"medias"`
      }


      
        resp, err := http.Get("https://skizo.tech/api/facebook?url="+ m.Querry +"&apikey=batu")
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
      

        var result Response
        err = json.Unmarshal(body, &result)
        if err != nil {
          fmt.Println("Error:", err)
          return
        }
      
      
      bytes, err := client.GetBytes(result.Medias[1].URL)
      if err != nil {
        fmt.Println("Error:", err)
        return
      }
      client.SendVideo(m.From, bytes, " ", m.ID)
      */
    },
  })
}

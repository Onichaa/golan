package download

import (
  "inc/lib"
  "fmt"
  "encoding/json"
)

func init() {
  lib.NewCommands(
    &lib.ICommand{
    Name:     "(cp|capcut)",
    As:       []string{"capcut"},
    Tags:     "downloader",
    IsPrefix: true,
    IsQuerry: true,
    IsWaitt:  true,
    Exec: func(client *lib.Event, m *lib.IMessage) {

      type Result struct {
        Code        int    `json:"code"`
        Title       string `json:"title"`
        Description string `json:"description"`
        Usage       string `json:"usage"`
        OriginalVideoUrl string `json:"originalVideoUrl"`
        CoverUrl    string `json:"coverUrl"`
        AuthorUrl   string `json:"authorUrl"`
      }
      
      data, err := lib.Capcutdl(m.Querry)
      if err != nil {
        fmt.Println("Error:", err)
        return
      }

      var result Result
      err = json.Unmarshal(data, &result)
      if err != nil {
        fmt.Println(err)
      }
      
      teks := `*CAPCUT DOWNLOADER*

𖦹 *Title:* ` + result.Title + `
𖦹 *Description:* ` + result.Description + `
𖦹 *Usage:* ` + result.Usage + `

Reply/balas video ini dengan ketik *.toaudio* untuk menjadikan video ke audio`


      bytes, err := client.GetBytes("https://ssscap.net"+result.OriginalVideoUrl)
      if err != nil {
        m.Reply(err.Error())
        return
      }
      client.SendVideo(m.From, bytes, teks, m.ID)
      
     
    },
  })
}
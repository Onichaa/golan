package ai

import (
  "inc/lib"
  "fmt"
  "net/http"
  "io"
   "regexp"
)

func init() {
  lib.NewCommands(
    &lib.ICommand{
    Name:     "(px|pixiv)",
    As:       []string{"pixiv"},
    Tags:     "ai",
    IsPrefix: true,
    IsQuerry: true,
    IsWaitt:  true,
    Exec: func(client *lib.Event, m *lib.IMessage) {


      data, err := lib.Pixiv(m.Querry)
      if err != nil {
        fmt.Println("Error:", err)
        return
      }

      regex := regexp.MustCompile(`(https?:\/\/[^\s]+)`)
       newLink := regex.FindStringSubmatch(data) 

      iniclient := &http.Client {
        }
      req1, err := http.NewRequest("GET", newLink[0], nil)
      req1.Header.Add("referer", "https://www.pixiv.net/")
        resps, err := iniclient.Do(req1)

      if err != nil {
        return
      }
      defer resps.Body.Close()

      imageData, err := io.ReadAll(resps.Body)
      if err != nil {
        return
      }

      client.SendImage(m.From, imageData, data, m.ID)
    },
  })
}
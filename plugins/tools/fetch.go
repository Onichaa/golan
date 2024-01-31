package tools

import (
  "fmt"
  "inc/lib"
  "inc/lib/api"
  "inc/lib/typings"
  "net/http"
  "strings"
  "io/ioutil"
  "bytes"
  "encoding/json"
)

func init() {
  lib.NewCommands(&lib.ICommand{
    Name:     "(get|fetch)",
    As:       []string{"get", "fetch"},
    Tags:     "tools",
    IsPrefix: true,
    IsQuerry: true,
    Exec: func(client *lib.Event, m *lib.IMessage) {


      resp, err := http.Get(m.Querry)
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

      mime := resp.Header.Get("Content-Type")

      if mime == "" {
        m.Reply("No Content-Type")
      }

      if strings.Contains(mime, "video") {
        client.SendVideo(m.From, body, m.Querry, m.ID)   
      } else if strings.Contains(mime, "webp") {
        s := api.StickerApi(&typings.Sticker{
          File: body,
          Tipe: func() typings.MediaType {
            if m.IsImage || m.IsQuotedImage || m.IsQuotedSticker {
              return typings.IMAGE
            } else if m.IsVideo || m.IsQuotedVideo {
              return typings.VIDEO
            } else {
              return typings.TEKS
            }

          }(),
        }, &typings.MetadataSticker{
          Author:    m.PushName,
          Pack:      "https://s.id/ryuubot",
          KeepScale: true,
          Removebg:  "true",
          Circle: func() bool {
            if m.Querry == "-c" {
              return true
            } else {
              return false
            }
          }(),
        })

        client.SendSticker(m.From, s.Build(), m.ID)
      } else if strings.Contains(mime, "image") {
        client.SendImage(m.From, body, m.Querry, m.ID)
      } else if strings.Contains(mime, "audio") {
        client.SendAudio(m.From, body, false, m.ID)
      } else if strings.Contains(mime, "json") {
        formattedJSON := formatJSON(body)
        m.Reply(formattedJSON)
      } else if strings.Contains(mime, "text") {
        m.Reply(string(body))
      }
      
    },
  })
}


func formatJSON(data []byte) string {
  var prettyJSON bytes.Buffer
  err := json.Indent(&prettyJSON, data, "", "  ")
  if err != nil {
    return "Error formatting JSON"
  }
  return prettyJSON.String()
}

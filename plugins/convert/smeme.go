package convert

import (
  "inc/lib"
"fmt"
  "os"
  "inc/lib/api"
  "inc/lib/typings"
)

func init() {
  lib.NewCommands(&lib.ICommand{
    Name:     "(sm|smeme)",
    As:       []string{"smeme"},
    Tags:     "convert",
    IsPrefix: true,
    IsMedia:  true,
     IsWaitt:  true,
     IsQuerry: true,
    Exec: func(client *lib.Event, m *lib.IMessage) {

      byte, _ := client.WA.Download(m.Media)

    randomJpgImg := "./" + lib.GetRandomString(5) + ".jpg"
    if err := os.WriteFile(randomJpgImg, byte, 0600); err != nil {
        fmt.Printf("Failed to save image: %v", err)
        return
    }
    //log.Printf("Saved image in %s", randomJpgImg)
    url, err := lib.Upload(randomJpgImg)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
       res := "https://api.memegen.link/images/custom/-/" +m.Querry+ ".jpg?background="+url
      
      bytes, err := client.GetBytes(res)
      if err != nil {
         fmt.Println("Error:", err)
        return
      }
      s := api.StickerApi(&typings.Sticker{
        File: bytes,
        Tipe: func() typings.MediaType {
          if m.IsImage || m.IsQuotedImage || m.IsQuotedSticker {
            return typings.IMAGE
          } else {
            return typings.VIDEO
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
      os.Remove(randomJpgImg)

    },
  })
}
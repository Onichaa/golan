package convert

import (
  "inc/lib"
//"fmt"
  "inc/lib/api"
  "inc/lib/typings"
  "math/rand"
  //"time"
  "github.com/disintegration/imaging"
  "github.com/fogleman/gg"
  "net/http"
  "strings"
  "bytes"
  "encoding/json"
  "encoding/base64"
  "strconv"
)

func init() {
  lib.NewCommands(&lib.ICommand{
    Name:     "qc",
    As:       []string{"qc"},
    Tags:     "convert",
    IsPrefix: true,
    IsMedia:  false,
     IsWaitt:  true,
     IsQuerry: true,
    Exec: func(client *lib.Event, m *lib.IMessage) {

      type From struct {
        ID   int    `json:"id"`
        Name string `json:"name"`
        Photo struct {
          URL string `json:"url"`
        } `json:"photo"`
      }

      type JsonResponse struct {
        Ok    bool   `json:"ok"`
        Result struct {
          Image string `json:"image"`
        } `json:"result"`
      }
      
      type Message struct {
        Entities []interface{} `json:"entities"`
        Media    struct {
          URL string `json:"url"`
        } `json:"media"`
        Avatar   bool      `json:"avatar"`
        From     From      `json:"from"`
        Text     string    `json:"text"`
        ReplyMessage interface{} `json:"replyMessage"`
      }

      teks := m.Querry
      name := m.PushName
      avatar := "https://telegra.ph/file/89c1638d9620584e6e140.png"
      id := pickRandom([]int{0, 4, 5, 3, 2, 7, 5, 9, 8, 1, 6, 10, 9, 7, 5, 3, 1, 2, 4, 6, 8, 0, 10})
     // anu := "http://example.com/anu.png"

      jsonStr := `{
        "type": "quote",
        "format": "png",
        "backgroundColor": "#FFFFFF",
        "width": 512,
        "height": 768,
        "scale": 2,
        "messages": [
          {
            "entities": [],
            "avatar": true,
            "from": {
              "id": ` + strconv.Itoa(id) + `,
              "name": "` + name + `",
              "photo": {
                "url": "` + avatar + `"
              }
            },
            "text": "` + teks + `",
            "replyMessage": {}
          }
        ]
      }`

      var jsonData JsonResponse
      resp, err := http.Post("https://quoteapi-ld81.onrender.com/generate", "application/json", strings.NewReader(jsonStr))
      if err != nil {
        // handle error
      }
      defer resp.Body.Close()

      err = json.NewDecoder(resp.Body).Decode(&jsonData)
      if err != nil {
        // handle error
      }

      if !jsonData.Ok {
        // handle error
      }

      buffer, err := base64.StdEncoding.DecodeString(jsonData.Result.Image)
      if err != nil {
        // handle error
      }

      dc := gg.NewContext(512, 768)
      img, err := imaging.Decode(bytes.NewReader(buffer))
      if err != nil {
        // handle error
      }
      dc.DrawImage(img, 0, 0)

     
      /*
      warna := []string{"000000", "FFFFFF", "999999", "c27ba0", "bcbcbc", "1f2c34"}
      rand.Seed(time.Now().UnixNano())
      index := rand.Intn(len(warna))
      hasil := warna[index]
      res := "https://skizo.tech/api/qc?text="+ m.Querry +"&username="+m.PushName+"&avatar=https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_960_720.png?q=60&apikey=batu&hex="+hasil
      bytes, err := client.GetBytes(res)
      if err != nil {
         fmt.Println("Error:", err)
        return
      }
      */

      
      s := api.StickerApi(&typings.Sticker{
        File: buffer,
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
      
    },
  })
}
func pickRandom(arr []int) int {
  return arr[rand.Intn(len(arr))]
}
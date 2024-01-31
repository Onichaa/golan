package convert

import (
  "inc/lib"
"fmt"
  "strings"
  "net/url"
  "net/http"
  "inc/lib/api"
  "inc/lib/typings"
   "encoding/json"
)

func init() {
  lib.NewCommands(&lib.ICommand{
    Name:     "(emojimix|emoji)",
    As:       []string{"emojimix"},
    Tags:     "convert",
    IsPrefix: true,
    IsMedia:  false,
     IsWaitt:  true,
     IsQuerry: true,
    Exec: func(client *lib.Event, m *lib.IMessage) {


emojis := strings.Split(m.Querry, "+")
if len(emojis) != 2 {
    m.Reply("Example: .emojimix😅+🤔")
    return
}

emoji1 := emojis[0]
emoji2 := emojis[1]

url := fmt.Sprintf("https://tenor.googleapis.com/v2/featured?key=AIzaSyAyimkuYQYF_FXVALexPuGQctUWRURdCYQ&contentfilter=high&media_filter=png_transparent&component=proactive&collection=emoji_kitchen_v5&q=%s_%s", url.QueryEscape(emoji1), url.QueryEscape(emoji2))

resp, err := http.Get(url)
if err != nil {
    fmt.Println("Error:", err)
    return
}
defer resp.Body.Close()

var data struct {
    Results []struct {
        URL string `json:"url"`
    } `json:"results"`
}

err = json.NewDecoder(resp.Body).Decode(&data)
if err != nil {
    fmt.Println("Error:", err)
    return
}

for _, res := range data.Results {
  bytes, err := client.GetBytes(res.URL)
  if err != nil {
     fmt.Println("Error:", err)
    return
  }
  
  s := api.StickerApi(&typings.Sticker{
    File: bytes,
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
  
}

},
})
}

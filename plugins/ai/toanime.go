package ai

import (
  "inc/lib"
"fmt"
)

func init() {
  lib.NewCommands(&lib.ICommand{
    Name:     "(toanime|jadianime)",
    As:       []string{"toanime"},
    Tags:     "ai",
    IsPrefix: true,
    IsMedia:  true,
     IsWaitt:  true,
    Exec: func(client *lib.Event, m *lib.IMessage) {

        m.Reply("Processing 2 minutes...")
        
      data, _ := client.WA.Download(m.Media)

      url, err := client.UploadImage(data)
      if err != nil {
          fmt.Println("Error:", err)
          return
      }
      res := "https://skizo.tech/api/toanime?url=" + url + "&apikey=batu" 

      bytes, err := client.GetBytes(res)
      if err != nil {
         fmt.Println("Error:", err)
        return
      }
      client.SendImage(m.From, bytes, "nihhh", m.ID)

    },
  })
}

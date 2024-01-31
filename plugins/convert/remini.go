package convert

import (
  "inc/lib"
"fmt"
)

func init() {
  lib.NewCommands(&lib.ICommand{
    Name:     "(hd|remini)",
    As:       []string{"remini"},
    Tags:     "convert",
    IsPrefix: true,
    IsMedia:  true,
     IsWaitt:  true,
    Exec: func(client *lib.Event, m *lib.IMessage) {
      
        m.Reply("fitur ini tidak digunakan lgi, erorrrrrrrrrr")
        
      data, _ := client.WA.Download(m.Media)

      url, err := client.UploadImage(data)
      if err != nil {
          fmt.Println("Error:", err)
          return
      }
      res := "https://skizo.tech/api/remini?url=" + url + "&apikey=kertas" 
      
      bytes, err := client.GetBytes(res)
      if err != nil {
         fmt.Println("Error:", err)
          m.Reply("Error")
        return
      }
      client.SendImage(m.From, bytes, "nihhh", m.ID)

    },
  })
}

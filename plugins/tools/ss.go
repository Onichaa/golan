package tools

import (
  "inc/lib"
  "fmt"
  "net/url"

)

func init() {
  lib.NewCommands(&lib.ICommand{
    Name:     "(ss|ssweb)",
    As:       []string{"screnshot"},
    Tags:     "tools",
    IsPrefix: true,
     IsQuerry: true,
   IsWaitt:  true,
    Exec: func(client *lib.Event, m *lib.IMessage) {

      res := "https://api.apiflash.com/v1/urltoimage?access_key=185eff3aa9fe4e3c8e30bda63b1fb9cf&wait_until=page_loaded&url=" + url.QueryEscape(m.Querry)
      bytes, err := client.GetBytes(res)
      if err != nil {
         fmt.Println("Error:", err)
        return
      }
client.SendImage(m.From, bytes, "Screenshot Web", m.ID)
    },
  })
}

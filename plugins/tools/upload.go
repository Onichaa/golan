package tools

import (
  "inc/lib"
"fmt"
)

func init() {
  lib.NewCommands(&lib.ICommand{
    Name:     "(up|upload)",
    As:       []string{"upload"},
    Tags:     "tools",
    IsPrefix: true,
    IsMedia:  true,
     IsWaitt:  true,
    Exec: func(client *lib.Event, m *lib.IMessage) {

      data, _ := client.WA.Download(m.Media)

      url, err := client.UploadImage(data)
      if err != nil {
         fmt.Println(err)
          return
      }
    m.Reply(url)

    },
  })
}

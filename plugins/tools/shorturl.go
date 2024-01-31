package tools

import (
  "inc/lib"
  "fmt"
)

func init() {
  lib.NewCommands(&lib.ICommand{
    Name:     "(shorturl|short)",
    As:       []string{"shorturl"},
    Tags:     "tools",
    IsPrefix: true,
    IsQuerry: true,
    IsWaitt:  true,
    Exec: func(client *lib.Event, m *lib.IMessage) {

      res, err := lib.ShortUrl(m.Querry)
      if err != nil {
        fmt.Println("Error:", err)
        return
      }
         m.Reply(res)

    },
  })
}
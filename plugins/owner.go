package plugins

import (
  "inc/lib"
)

func init() {
  lib.NewCommands(&lib.ICommand{
    Name:     "(owner|pemilik)",
    As:       []string{"owner"},
    Tags:     "main",
    IsPrefix: true,
    Exec: func(client *lib.Event, m *lib.IMessage) {   
      client.SendContact(m.From, "628388024064", "Takim", m.ID)

    },
  })
}

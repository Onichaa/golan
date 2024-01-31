package info

import (
  "fmt"
  "inc/lib"
  "time"
)

func init() {
  lib.NewCommands(&lib.ICommand{
    Name:     "runtime",
    As:       []string{"runtime"},
    Tags:     "info",
    IsPrefix: true,
    Exec: func(client *lib.Event, m *lib.IMessage) {

elapsed := time.Since(lib.GetUptime())
days := int(elapsed.Hours()) / 24
hours := int(elapsed.Hours()) % 24
minutes := int(elapsed.Minutes()) % 60
seconds := int(elapsed.Seconds()) % 60
        m.Reply(fmt.Sprintf("Bot aktif selama: \n%02d Day(s) %02d Hour(s) %02d Minute(s) and %02d Seconds!", days, hours, minutes, seconds))
      
        },
        })
      }

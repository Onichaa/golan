package info

import (
  "inc/lib"
  "math/rand"
  "time"
)

func init() {
  lib.NewCommands(&lib.ICommand{
    Name:     "bot",
    As:       []string{"ping"},
    Tags:     "info",
   //IsPrefix: true,
    Exec: func(client *lib.Event, m *lib.IMessage) {   

      warna := []string{"Eum", "knpa", "haa",  "botaktif", "haaaa capekkk", "napa ayng", "napa bebbbkuuu", "napa ay", "knpa atuh manggil aku", "knpan lgiiiii", "hayoyo apa pulakkk", "huu ayang npa manggill"}
      rand.Seed(time.Now().UnixNano())
      index := rand.Intn(len(warna))
      hasil := warna[index]
      
      m.Reply(hasil)
    },
  })
}

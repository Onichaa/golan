package convert

import (
  "inc/lib"
  "fmt"
  "io/ioutil"
  "os"
)

func init() {
  lib.NewCommands(&lib.ICommand{
    Name:     "(toaudio|audio)",
    As:       []string{"toaudio"},
    Tags:     "convert",
    IsPrefix: true,
    IsMedia:  true,
     IsWaitt:  true,
    // IsQuerry: true,
    Exec: func(client *lib.Event, m *lib.IMessage) {

       ran := "./" + lib.GetRandomString(5) + ".mp3"
      
      byte, _ := client.WA.Download(m.Media)
        whatsappAudio, err := lib.ToAudio(byte, "mp4")
        if err != nil {
          fmt.Println("Error:", err)
          return
        }

        err = ioutil.WriteFile(ran, whatsappAudio, 0644)
        if err != nil {
          fmt.Println("Error:", err)
          return
        }
      url, err := lib.Upload(ran)
      if err != nil {
          fmt.Println("Error:", err)
          return
      }
      fmt.Println(url)
      bytes, err := client.GetBytes(url)
      if err != nil {
        m.Reply(err.Error())
        return
      }
      client.SendAudio(m.From, bytes, false, m.ID)
     // client.SendDocument(m.From, bytes, fmt.Sprintf("%s.mp3", "audio"), " ", m.ID)
      os.Remove(ran)
    },
  })
}
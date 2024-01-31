package convert

import (
  "inc/lib"
  "fmt"
  "os"
  "bytes"
  "os/exec"

)

func init() {
  lib.NewCommands(&lib.ICommand{
    Name:     "(toimage|toimg)",
    As:       []string{"toimage"},
    Tags:     "convert",
    IsPrefix: true,
    IsMedia:  true,
     IsWaitt:  true,
    Exec: func(client *lib.Event, m *lib.IMessage) {

      byte, _ := client.WA.Download(m.Media)

       ran := "./" + lib.GetRandomString(5) + ".png"
       randomJpgImg := "./" + lib.GetRandomString(5) + ".jpg"
      
    if err := os.WriteFile(randomJpgImg, byte, 0600); err != nil {
        fmt.Printf("Failed to save image: %v", err)
        return
    }
 
      // Run ffmpeg command
      cmd := exec.Command("ffmpeg", "-i", randomJpgImg, ran)
      var out bytes.Buffer
      var stderr bytes.Buffer
      cmd.Stdout = &out
      cmd.Stderr = &stderr
      err := cmd.Run()

      // Check error
      if err != nil {
        fmt.Println("Error:", err)
        return
      }
      
    //log.Printf("Saved image in %s", randomJpgImg)
    url, err := lib.UploadV2(ran)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
      bytes, err := client.GetBytes(url)
      if err != nil {
         fmt.Println("Error:", err)
        return
      }
      client.SendImage(m.From, bytes, "nih", m.ID)

      os.Remove(ran)
      os.Remove(randomJpgImg)
    },
  })
}
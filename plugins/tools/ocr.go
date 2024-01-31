package tools

import (
  "inc/lib"
  "fmt"
  "io/ioutil"
  "os"
  "os/exec"

)

func init() {
  lib.NewCommands(&lib.ICommand{
    Name:     "ocr",
    As:       []string{"ocr"},
    Tags:     "tools",
    IsPrefix: true,
    IsMedia: true, 
    Exec: func(client *lib.Event, m *lib.IMessage) {     
      byte, _ := client.WA.Download(m.Media)

         randomJpgImg := "./" + lib.GetRandomString(5) + ".jpg"

      if err := os.WriteFile(randomJpgImg, byte, 0600); err != nil {
          fmt.Printf("Failed to save image: %v", err)
          return
      }

      cmd := exec.Command("tesseract", randomJpgImg, "ocr")
      err := cmd.Run()
      if err != nil {
        fmt.Println(err)
        return
      }

      txt, err := ioutil.ReadFile("./ocr.txt")
      if err != nil {
        fmt.Println(err)
        return
      }

      encmedia := string(txt)
      m.Reply(encmedia)

       os.Remove(randomJpgImg)

    },
  })
}

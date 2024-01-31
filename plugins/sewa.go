package plugins

import (
  "inc/lib"
    "fmt"
)

func init() {
  lib.NewCommands(&lib.ICommand{
    Name:     "(sewa|sewabot)",
    As:       []string{"sewabot"},
    Tags:     "main",
    IsPrefix: true,
    Exec: func(client *lib.Event, m *lib.IMessage) {  

       teks := `*PRICE BOT FOR GROUP*\n\n- 7 hari [ 3 ribu ]\n- 14 hari [ 5 ribu ]
Silahkan ketik .owner untuk mengetahui pemilik bot ini`
      res := "https://telegra.ph//file/024a6a1fe28189278a5f7.jpg" 

      bytes, err := client.GetBytes(res)
      if err != nil {
         fmt.Println("Error:", err)
        return
      }
      client.SendImage(m.From, bytes, teks, m.ID)
      

    },
  })
}

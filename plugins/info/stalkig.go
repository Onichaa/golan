package info

import (
  "inc/lib"
  "fmt"
  "net/http"
   "encoding/json"
  "io/ioutil"
  "strconv"
)

func init() {
  lib.NewCommands(
    &lib.ICommand{
    Name:     "(stalkig|igstalk)",
    As:       []string{"igstalk"},
    Tags:     "info",
    IsPrefix: true,
    IsQuerry: true,
    IsWaitt:  true,
    Exec: func(client *lib.Event, m *lib.IMessage) {

      type User struct {
        PhotoProfile string `json:"photo_profile"`
        Username     string `json:"username"`
        Fullname     string `json:"fullname"`
        Posts        int    `json:"posts"`
        Followers    int    `json:"followers"`
        Following    int    `json:"following"`
        Bio          string `json:"bio"`
      }

      resp, err := http.Get("https://skizo.tech/api/igstalk?user="+ m.Querry +"&apikey=batu")
      if err != nil {
        fmt.Println("Error:", err)
        return
      }
      defer resp.Body.Close()

      body, err := ioutil.ReadAll(resp.Body)
      if err != nil {
        fmt.Println("Error:", err)
        return
      }

      var user User
      err = json.Unmarshal(body, &user)
      if err != nil {
        fmt.Println("Error:", err)
        return
      }
      
      teks := `*INSTAGRAM STALK*

𖦹 *Username:* ` + user.Username + `
𖦹 *Fullname:* ` + user.Fullname + `
𖦹 *Posts:* ` + strconv.Itoa(user.Posts) + `
𖦹 *Followers:* ` + strconv.Itoa(user.Followers) + `
𖦹 *Following:* ` + strconv.Itoa(user.Following) + `
𖦹 *Bio:* ` + user.Bio + `
      `

      bytes, err := client.GetBytes(user.PhotoProfile)
      if err != nil {
        m.Reply(err.Error())
        return
      }
     
      client.SendImage(m.From, bytes, teks, m.ID)
    },
  })
}
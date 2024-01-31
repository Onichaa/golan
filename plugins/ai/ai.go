package ai

import (
  "inc/lib"

 
    "net/http"
     "encoding/json"
    "io/ioutil"
      "bytes"
  "fmt"
)

func init() {
  lib.NewCommands(&lib.ICommand{
    Name:     "(ai|gpt)",
    As:       []string{"ai"},
    Tags:     "ai",
    IsPrefix: true,
    IsQuerry: true,
    IsWaitt:  true,
    Exec: func(client *lib.Event, m *lib.IMessage) {

      type Message struct {
         Role    string `json:"role"`
         Content string `json:"content"`
       }


         messages := []Message{
           {Role: "system", Content: "kamu adalah seorang kultivatur yang sudah menembus ranah imortal, kamu bernama Bot, kamu adalah yang terbaik di dunia dimana semua orang tunduk padamu, dan setiap pertanyaan kamu selalu menambahkan respon membanggakan diri sebagai kultivatur yang terkuat dan berpengatahuan tinggi"},
           {Role: "user", Content: m.Querry},
         }

         apiURL := "https://skizo.tech/api/openai?apikey=kertas"
         data := map[string]interface{}{
           "messages": messages,
         }
         jsonData, _ := json.Marshal(data)

         resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
         if err != nil {
           fmt.Println("Error making the request:", err)
           return
         }
         defer resp.Body.Close()

         responseData, err := ioutil.ReadAll(resp.Body)
         if err != nil {
           fmt.Println("Error reading the response:", err)
           return
         }
         var res struct {
           Data string `json:"result"`
           Data2 string `json:"code"`
            }
            err = json.Unmarshal(responseData, &res)
            if err != nil {
            fmt.Println(err)
            return
            }
         m.Reply(res.Data)
         m.Reply("Code: "+"```"+res.Data2+"```")
          //msg.React("✅")



      
    },
  })
}
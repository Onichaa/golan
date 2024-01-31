package owner

import (
  "fmt"
  "inc/lib"
  "os/exec"
  "os"
  "syscall"
)

func init() {
  lib.NewCommands(&lib.ICommand{
    Name:     "(r|restart)",
    As:       []string{"restart"},
    Tags:     "owner",
    IsPrefix: true,
    IsOwner:  true,
    Exec: func(client *lib.Event, m *lib.IMessage) {

          
      dir, err := os.Getwd()
      if err != nil {
        fmt.Println(err)
      }

         m.Reply("sukses restarting")
      cmd := exec.Command("go", "run", dir+"/main.go")
      cmd.Stdout = os.Stdout
      cmd.Stderr = os.Stderr
      cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
      err = cmd.Start()
      if err != nil {
        fmt.Println(err)
      }

      // Menunggu proses selesai
      err = cmd.Wait()
      if err != nil {
        fmt.Println(err)
      }
       
      
    },
  })
}





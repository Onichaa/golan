package lib

import (
	//"inc/lib/helpers"
    "fmt"
    "net/http"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
    "github.com/gorilla/websocket"
)



var upgrader = websocket.Upgrader{
  CheckOrigin: func(r *http.Request) bool {
    return true
  },
}


type IHandler struct {
	Container *store.Device
}

func NewHandler(container *sqlstore.Container) *IHandler {
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
        fmt.Println(err)
	}
	return &IHandler{
		Container: deviceStore,
	}
}

func (h *IHandler) Client(jbot ...bool) *whatsmeow.Client {
	clientLog := waLog.Stdout("lient", "ERROR", true)
	client := whatsmeow.NewClient(h.Container, clientLog)
	client.AddEventHandler(RegisterHandler(client, jbot...))
	return client
}


func RegisterHandler(client *whatsmeow.Client, jbot ...bool) func(evt interface{}) {
	return func(evt interface{}) {
		sock := NewClient(client)
		switch v := evt.(type) {
		case *events.Message:
      m := NewSmsg(v, sock, jbot...)
      //if !helpers.Public && !m.IsOwner {
        //return
     // }
       go Get(sock, m)
		
		}
	}
}


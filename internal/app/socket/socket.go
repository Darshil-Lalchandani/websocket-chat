package socket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var Clients = make(map[*websocket.Conn]string)
var broadcaster = make(chan ChatMessage)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request, id string) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("error upgrading connections")
	}
	defer ws.Close()
	Clients[ws] = id
	log.Print("All clients", len(Clients))
	for _, v := range Clients {
		log.Print(v)
	}

	for {
		var msg ChatMessage
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Print(err)
			delete(Clients, ws)
			break
		}
		broadcaster <- msg
	}
}

func HandleMessages() {
	for {
		msg := <-broadcaster
		log.Print("Message received", msg)
	}
}

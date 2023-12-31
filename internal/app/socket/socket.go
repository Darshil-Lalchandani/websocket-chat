package socket

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[string]Socket)
var receiver = make(chan ChatMessage)

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
	clients[id] = Socket{
		Conn: ws,
	}
	ws.WriteJSON(ChatMessage{Message: "Ack"})
	log.Print("New connection received, id: ", id)

	for {
		var msg ChatMessage
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Print(err)
			for id, socket := range clients {
				if socket.Conn == ws {
					delete(clients, id)
				}
			}
			break
		}
		receiver <- msg
	}
}

func HandleMessages() {
	for {
		msg := <-receiver
		log.Print("Message received: ", msg)
	}
}

func sendMessage(w http.ResponseWriter, r *http.Request, text string) {
	for _, c := range clients {
		msg := ChatMessage{
			Message: text,
		}
		c.Conn.WriteJSON(&msg)
		break
	}
	res, _ := json.Marshal(text)
	w.Write(res)
}

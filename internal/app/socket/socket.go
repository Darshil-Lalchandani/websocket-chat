package socket

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var Clients = make(map[string]Socket)
var broadcaster = make(chan ChatMessage)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request, id string, geo string) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("error upgrading connections")
	}
	defer ws.Close()
	timeNow := time.Now().UTC()
	Clients[id] = Socket{
		Conn:       ws,
		UsedCount:  0,
		Geo:        geo,
		LastUsedAt: time.Time{},
		CreatedAt:  timeNow,
	}
	ws.WriteJSON(ChatMessage{Message: "Ack"})
	log.Print("All clients", len(Clients))
	for id, v := range Clients {
		log.Print(id, v)
	}

	for {
		var msg ChatMessage
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Print(err)
			for id, socket := range Clients {
				if socket.Conn == ws {
					delete(Clients, id)
				}
			}
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

func querySocket(w http.ResponseWriter, r *http.Request, text string) {
	for _, c := range Clients {
		msg := ChatMessage{
			Message: text,
		}
		c.Conn.WriteJSON(&msg)
		break
	}
	res, _ := json.Marshal(text)
	w.Write(res)
}

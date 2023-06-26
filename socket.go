package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]string)
var broadcaster = make(chan ChatMessage)
var requests = make(map[string]ChatMessage)

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
	clients[ws] = id
	log.Print("All clients", len(clients))
	for _, v := range clients {
		log.Print(v)
	}

	for {
		var msg ChatMessage
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Print(err)
			delete(clients, ws)
			break
		}
		broadcaster <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcaster
		log.Print("Message received", msg)
	}
}

func querySocket(w http.ResponseWriter, r *http.Request, text string) {
	for c, id := range clients {
		msg := ChatMessage{
			Message: text,
		}
		c.WriteJSON(&msg)
		requests[id] = msg
		break
	}
	for id, cm := range requests {
		log.Print(id, cm)
	}
	res, _ := json.Marshal(text)
	w.Write(res)
}

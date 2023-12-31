package socket

import (
	"github.com/gorilla/websocket"
)

type ChatMessage struct {
	Message string `json:"text"`
}

type Socket struct {
	Conn *websocket.Conn
}

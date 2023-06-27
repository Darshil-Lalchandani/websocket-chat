package socket

import (
	"time"

	"github.com/gorilla/websocket"
)

type ChatMessage struct {
	Message   string `json:"text"`
	RequestId string `json:"request_id"`
}

type Socket struct {
	Conn       *websocket.Conn
	UsedCount  int
	Geo        string
	IP         string
	LastUsedAt time.Time
	CreatedAt  time.Time
}

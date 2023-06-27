package socket

import (
	"container/list"
	"fmt"
	"log"
)

type SocketPool struct {
	Pool *list.List
}

func InitiatePool() SocketPool {
	pool := list.New()
	socketPool := SocketPool{
		Pool: pool,
	}
	return socketPool
}

func PushFront(id string, sp SocketPool) {
	element := sp.Pool.PushFront(id)
	log.Print("Added element to front of list", element.Value, id, "list length ", sp.Pool.Len())
}

// gets the front element and pushes to back
func GetSocketId(id string, sp SocketPool, clients map[string]Socket) (string, error) {
	front := sp.Pool.Front()
	if front == nil {
		log.Print("No element found in the list")
		return "", fmt.Errorf("no element found in the list")
	}
	sp.Pool.Remove(front)
	frontValue := front.Value
	log.Print("Found element in the list with ID ", frontValue)
	if clients[frontValue.(string)].UsedCount > 50 {
		log.Print("50 used count exceeded for ID:", frontValue)
	} else {
		sp.Pool.PushBack(frontValue)
		log.Print("Pushed to the back of list ", frontValue)
	}
	return frontValue.(string), nil
}

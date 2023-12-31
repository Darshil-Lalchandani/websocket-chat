package main

import (
	"log"
	"net/http"
	"os"

	"example.com/websocket-chat/internal/app/socket"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	port := os.Getenv("PORT")
	log.Print("server starting at localhost:4444")

	http.Handle("/", http.FileServer(http.Dir("./public")))
	socket.MountSocketRoutes()

	go socket.HandleMessages()
	http.ListenAndServe(":"+port, nil)
}

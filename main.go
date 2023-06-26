package main

import (
	"log"
	"net/http"
	"os"

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
	http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		handleConnections(w, r, id)
	})
	http.HandleFunc("/forwardRequest", func(w http.ResponseWriter, r *http.Request) {
		text := r.URL.Query().Get("text")
		querySocket(w, r, text)
	})
	go handleMessages()
	http.ListenAndServe(":"+port, nil)
}

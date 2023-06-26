package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/sellerapp-com/scraper-net/internal/app/redirector"
	"github.com/sellerapp-com/scraper-net/internal/app/socket"
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
	redirector.MountRedirectionRoutes()

	go socket.HandleMessages()
	http.ListenAndServe(":"+port, nil)
}

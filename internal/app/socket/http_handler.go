package socket

import "net/http"

func MountSocketRoutes() {
	http.HandleFunc("/websocket/connect", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		handleConnections(w, r, id)
	})
	http.HandleFunc("/websocket/send", func(w http.ResponseWriter, r *http.Request) {
		text := r.URL.Query().Get("message")
		toID := r.URL.Query().Get("id")
		sendMessage(w, r, text, toID)
	})
}

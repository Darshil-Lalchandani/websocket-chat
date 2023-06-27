package socket

import "net/http"

func MountSocketRoutes() {
	http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		geo := r.URL.Query().Get("geo")
		handleConnections(w, r, id, geo)
	})
	http.HandleFunc("/forwardRequest", func(w http.ResponseWriter, r *http.Request) {
		text := r.URL.Query().Get("text")
		querySocket(w, r, text)
	})
}

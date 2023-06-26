package socket

import "net/http"

func MountSocketRoutes() {
	http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		handleConnections(w, r, id)
	})
}

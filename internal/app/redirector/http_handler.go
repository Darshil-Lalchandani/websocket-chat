package redirector

import "net/http"

func MountRedirectionRoutes() {
	http.HandleFunc("/forwardRequest", func(w http.ResponseWriter, r *http.Request) {
		text := r.URL.Query().Get("text")
		querySocket(w, r, text)
	})
}

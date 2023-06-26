package redirector

import (
	"encoding/json"
	"net/http"

	"github.com/sellerapp-com/scraper-net/internal/app/socket"
)

func querySocket(w http.ResponseWriter, r *http.Request, text string) {
	for c := range socket.Clients {
		msg := socket.ChatMessage{
			Message: text,
		}
		c.WriteJSON(&msg)
		break
	}
	res, _ := json.Marshal(text)
	w.Write(res)
}

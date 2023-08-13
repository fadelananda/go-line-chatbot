package healthcheck

import (
	"net/http"

	util "github.com/fadelananda/go-line-chatbot/internal/utils"
)

func broadcastLineMessageHandler(w http.ResponseWriter, r *http.Request) {
	util.RespondWithJSON(w, 200, struct{}{})
}

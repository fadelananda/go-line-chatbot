package api

import (
	"net/http"

	"github.com/fadelananda/go-line-chatbot/internal/utils"
)

func ServerLivenessHealthcheck(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, 200, struct{}{})
}

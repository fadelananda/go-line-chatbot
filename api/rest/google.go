package rest

import (
	"fmt"
	"net/http"

	"github.com/fadelananda/go-line-chatbot/internal/service"
	"github.com/fadelananda/go-line-chatbot/internal/utils"
	"github.com/go-chi/chi"
)

type googleServiceHandler struct {
	service *service.GoogleService
}

func InitGoogleRESTHandler(router chi.Router, service *service.GoogleService) {
	h := &googleServiceHandler{
		service: service,
	}

	googleRouter := chi.NewRouter()
	googleRouter.Get("/", h.OauthCallbackHandler)

	router.Mount("/", googleRouter)
}

func (h *googleServiceHandler) OauthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	jwtToken := r.URL.Query().Get("state")

	if code != "" && jwtToken != "" {
		err := h.service.HandleOauthCallback(code, jwtToken)
		if err != nil {
			fmt.Println(err)
		}

		utils.RespondWithJSON(w, 200, map[string]interface{}{
			"message":   "authentication successful",
			"code":      code,
			"jwt_token": jwtToken,
		})
	} else {
		utils.RespondWithError(w, 400, "invalid URL")
	}

}

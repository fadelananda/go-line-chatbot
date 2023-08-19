package rest

import (
	"fmt"
	"net/http"

	"github.com/fadelananda/go-line-chatbot/internal/middleware"
	"github.com/fadelananda/go-line-chatbot/internal/service"
	"github.com/fadelananda/go-line-chatbot/internal/utils"
	"github.com/go-chi/chi"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type lineServiceHandler struct {
	service *service.LineService
}

func InitLineRESTHandler(router chi.Router, service *service.LineService) {
	h := lineServiceHandler{
		service: service,
	}

	lineRouter := chi.NewRouter()
	lineRouter.Use(middleware.DecodeLineRequestBody(h.service.LineClient.GetLineBotInstance()))
	lineRouter.Post("/callback", h.LineBotWebhookHandler)

	router.Mount("/line", lineRouter)
}

func (h *lineServiceHandler) LineBotWebhookHandler(w http.ResponseWriter, r *http.Request) {
	events, ok := r.Context().Value(middleware.ContextKey("lineEvents")).([]*linebot.Event)
	if !ok {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("err")
		return
	}

	h.service.HandleWebhookEvents(events)

	utils.RespondWithJSON(w, 200, struct{}{})
}

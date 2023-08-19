package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/fadelananda/go-line-chatbot/internal/utils"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type ContextKey string

func DecodeLineRequestBody(lineBot *linebot.Client) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			events, err := lineBot.ParseRequest(r)
			if err != nil {
				fmt.Println(err)
			}

			utils.LogInfo("decoding line request", map[string]interface{}{
				"parsed_request": events,
			})

			ctx := context.WithValue(r.Context(), ContextKey("lineEvents"), events)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

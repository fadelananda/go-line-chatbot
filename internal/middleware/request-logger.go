package middleware

import (
	"io"
	"net/http"

	"github.com/fadelananda/go-line-chatbot/internal/utils"
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestBody, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		r.Body.Close()

		utils.LogInfo("Request received", map[string]interface{}{
			"method":  r.Method,
			"URL":     r.RequestURI,
			"headers": r.Header,
			"body":    requestBody,
		})

		next.ServeHTTP(w, r)
	})
}

package middleware

import (
	"bytes"
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
		defer r.Body.Close()

		utils.LogInfo("request received", map[string]interface{}{
			"method":  r.Method,
			"URL":     r.RequestURI,
			"headers": r.Header,
			"body":    requestBody,
		})

		// Restore the request body for subsequent handlers
		r.Body = io.NopCloser(bytes.NewBuffer(requestBody))

		next.ServeHTTP(w, r)
	})
}

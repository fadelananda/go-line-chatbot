package healthcheck

import (
	"net/http"

	"github.com/go-chi/chi"
)

func NewHealthCheckRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/readiness", broadcastLineMessageHandler)

	return r
}

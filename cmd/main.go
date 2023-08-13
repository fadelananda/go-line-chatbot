package main

import (
	"log"
	"net/http"
	"os"

	"github.com/fadelananda/go-line-chatbot/internal/app/healthcheck"
	"github.com/fadelananda/go-line-chatbot/internal/utils"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// init router
	readinessRouter := healthcheck.NewHealthCheckRouter()
	router.Mount("/healthcheck", readinessRouter)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	utils.InitLogger()
	utils.Logger.Info("Server starting at port ", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

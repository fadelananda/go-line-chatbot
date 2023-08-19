package main

import (
	"log"
	"net/http"
	"os"

	"github.com/fadelananda/go-line-chatbot/api"
	"github.com/fadelananda/go-line-chatbot/api/rest"
	"github.com/fadelananda/go-line-chatbot/internal/client"
	"github.com/fadelananda/go-line-chatbot/internal/middleware"
	"github.com/fadelananda/go-line-chatbot/internal/service"
	"github.com/fadelananda/go-line-chatbot/internal/utils"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")

	// init client
	lineClient, err := client.NewLineClient()
	if err != nil {
		utils.LogError("unable to initialize line bot client", err, nil)
		os.Exit(1)
	}
	googleCalendarClient, err := client.NewGoogleCalendarClient()
	if err != nil {
		utils.LogError("unable to initialize google calendar client", err, nil)
		os.Exit(1)
	}

	// init service
	lineService := service.NewLineService(lineClient, googleCalendarClient)
	googleCalendarService := service.NewGoogleService(googleCalendarClient)

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	router.Use(middleware.LogRequest)

	// init REST handler
	rest.InitLineRESTHandler(router, lineService)
	rest.InitGoogleRESTHandler(router, googleCalendarService)

	// init router
	readinessRouter := api.NewHealthCheckRouter()
	router.Mount("/healthcheck", readinessRouter)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	utils.InitLogger()
	utils.LogInfo("Server starting at port 3000", nil)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fadelananda/go-line-chatbot/api"
	"github.com/fadelananda/go-line-chatbot/api/rest"
	"github.com/fadelananda/go-line-chatbot/internal/client"
	"github.com/fadelananda/go-line-chatbot/internal/middleware"
	"github.com/fadelananda/go-line-chatbot/internal/repository"
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
	awsClient, err := client.NewAWSClient()
	if err != nil {
		utils.LogError("unable to initialize aws client", err, nil)
		os.Exit(1)
	}

	// init repository
	userRepository := repository.NewUserRepository(awsClient)

	// init service
	lineService := service.NewLineService(lineClient, googleCalendarClient, awsClient)
	googleCalendarService := service.NewGoogleService(googleCalendarClient, userRepository)

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

	iddleConnsClosed := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)
		signal.Notify(s, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
		<-s

		utils.LogInfo("shutting down server...", nil)
		time.Sleep(1 * time.Second)

		ctx, cancel := context.WithTimeout(context.Background(), 60&time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Println("err")
		}

		close(iddleConnsClosed)
	}()

	utils.InitLogger()
	utils.LogInfo("Server starting at port 3000", nil)
	err = server.ListenAndServe()
	if err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("fatal http server failed to start:", err)
		}
	}

	<-iddleConnsClosed
	utils.LogInfo("server stopped", nil)
}

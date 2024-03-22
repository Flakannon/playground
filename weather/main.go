package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/playground/weather/database/weatherDB"
	"github.com/playground/weather/restapi"
)

func main() {
	var app restapi.App
	app.WeatherClient = weatherDB.WeatherDBClientMock{}

	engine := restapi.NewEngine(app)

	server := http.Server{
		Addr:    ":8080",
		Handler: engine,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Listen and serve failed")
		}
	}()

	quit := make(chan os.Signal, 2)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Print("Shutdown")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Print("Exited properly")
}

package main

import (
	"Relocate/accommodation/internal/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "accommodation-service", log.LstdFlags)
	accommodationHandler := handlers.NewAccommodation(logger)

	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/", accommodationHandler.ServeHTTP)

	server := &http.Server{
		Addr:         ":4000",
		Handler:      serverMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <-sigChan
	logger.Println("Recieved terminate, gracefull shutdown", sig)

	timeoutContext, err := context.WithTimeout(context.Background(), 30*time.Second)
	if err != nil {
		logger.Fatal(err)
	}
	server.Shutdown(timeoutContext)
}

package app

import (
	"Relocate/accommodation/internal/handlers"
	"Relocate/accommodation/pkg/db"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Run() {
	logger := log.New(os.Stdout, "accommodation-service", log.LstdFlags)

	err := godotenv.Load("/opt/app/.env")
	if err := godotenv.Load(); err != nil {
		logger.Fatal(fmt.Sprintf("Error loading .env file - %#v", err))
	}

	dbConnection, err := db.NewPostgresDB(db.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_DBNAME"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})

	if err != nil {
		logger.Fatal(fmt.Sprintf("DB connection error - %#v --- %#v", dbConnection, err))
	}

	logger.Println("db connected")

	accommodationHandler := handlers.NewAccommodation(logger, dbConnection)

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
			logger.Fatal(fmt.Sprintf("Start server error -  %#v", err))
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <-sigChan
	logger.Println("Recieved terminate, gracefull shutdown", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext)
}

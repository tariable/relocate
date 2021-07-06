package main

import (
	"Relocate/user/internal/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "accommodation-service", log.LstdFlags)
	helloHandler := handlers.NewHelloHandler(logger)
	byeHandler := handlers.NewByeHandler(logger)

	servMx := http.NewServeMux()
	servMx.HandleFunc("/hello", helloHandler.ServeHTTP)
	servMx.HandleFunc("/byeee", byeHandler.ServeHTTP)

	http.ListenAndServe(":4001", servMx)
}

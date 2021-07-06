package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type HelloHandler struct {
	logger *log.Logger
}

func NewHelloHandler(logger *log.Logger) *HelloHandler {
	return &HelloHandler{
		logger,
	}
}

func (h *HelloHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.logger.Println("Hello World")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello user - %s", d)
}

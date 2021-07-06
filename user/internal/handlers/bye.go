package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ByeHandler struct {
	logger *log.Logger
}

func NewByeHandler(logger *log.Logger) *ByeHandler {
	return &ByeHandler{
		logger,
	}
}

func (h *ByeHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.logger.Println("Bye World")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Bye user - %s", d)
}

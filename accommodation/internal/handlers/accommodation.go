package handlers

import (
	"Relocate/accommodation/internal/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Accommodation struct {
	logger *log.Logger
}

func NewAccommodation(logger *log.Logger) *Accommodation {
	return &Accommodation{logger}
}

func (accommodation *Accommodation) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		accommodation.GetAccommodations(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		accommodation.CreateAccommodation(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		// get id from URL
		regex := regexp.MustCompile("/([0-9])")
		ids := regex.FindAllStringSubmatch(r.URL.Path, -1)
		if len(ids) != 1 || len(ids[0]) != 2 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		accommodationId, err := strconv.Atoi(ids[0][1])
		if err != nil {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		accommodation.UpdateAccommodation(accommodationId, rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (accommodation *Accommodation) GetAccommodations(rw http.ResponseWriter, r *http.Request) {
	accommodation.logger.Println("Handle GET Accommodations")

	accommodations := data.GetAccommodations()

	err := accommodations.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal accommodation list", http.StatusInternalServerError)
	}
}

func (accommodation *Accommodation) CreateAccommodation(rw http.ResponseWriter, r *http.Request) {
	accommodation.logger.Println("Handle POST Accommodation")

	newAccommodation := &data.Accommodation{}
	err := newAccommodation.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
	accommodation.logger.Println("Accommodation data from request: %#v", newAccommodation)
	data.StoreAccommodation(newAccommodation)
}

func (accommodation *Accommodation) UpdateAccommodation(id int, rw http.ResponseWriter, r *http.Request) {
	accommodation.logger.Println("Handle PATCH Accommodation")

	newAccommodation := &data.Accommodation{}
	err := newAccommodation.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	accommodation.logger.Println("Accommodation data from request: %#v", newAccommodation)
	err = data.UpdateAccommodation(id, newAccommodation)
	if err != nil {
		http.Error(rw, "Accommodation not found", http.StatusNotFound)
		return
	}
}

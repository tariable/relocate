package data

import (
	"encoding/json"
	"fmt"
	"io"
)

type Accommodation struct {
	Id          int     `json:"id"`
	Status      string  `json:"status"`
	Description string  `json:"description"`
	Space       int64   `json:"space"`
	Price       int64   `json:"price"`
	Deposit     int64   `json:"deposit"`
	Floor       int64   `json:"floor"`
	MaxFloor    int64   `json:"max_floor"`
	RoomCount   int64   `json:"room_count"`
	Images      []Image `json:"images"`
}

type Image struct {
	Path string `json:"path"`
}

type Accommodations []*Accommodation

func (accommodation *Accommodation) FromJSON(reader io.Reader) error {
	e := json.NewDecoder(reader)
	return e.Decode(accommodation)
}

func (accommodations *Accommodations) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(accommodations)
}

func GetAccommodations() Accommodations {
	return AccommodationList
}

func StoreAccommodation(accommodation *Accommodation) {
	accommodation.Id = getNextID()
	AccommodationList = append(AccommodationList, accommodation)
}

func getNextID() int {
	return len(AccommodationList) + 1
}

func UpdateAccommodation(id int, accommodation *Accommodation) error {
	pos, err := getAccommodation(id)
	if err != nil {
		return err
	}
	accommodation.Id = id
	AccommodationList[pos] = accommodation

	return nil
}

func getAccommodation(id int) (int, error) {
	for index, acc := range AccommodationList {
		if acc.Id == id {
			return index, nil
		}
	}
	return -1, fmt.Errorf("Accommodation not found")
}

var AccommodationList = Accommodations{
	{
		Id:          1,
		Status:      "actual",
		Description: "nice flat on Baumanskaya",
	},
	{
		Id:          2,
		Status:      "closed",
		Description: "piece of shit on Kurskaya",
	},
}

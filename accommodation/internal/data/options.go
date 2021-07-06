package data

type Options struct {
	RentTerm          string `json:"rent_term"`
	WcType            string `json:"wc_type"`
	BathType          string `json:"bath_type"`
	Balcony           bool   `json:"balcony"`
	IsAllowedPets     bool   `json:"is_allowed_pets"`
	IsAllowedSmoking  bool   `json:"is_allowed_smoking"`
	IsAllowedChildren bool   `json:"is_allowed_children"`
	Elevator          bool   `json:"elevator"`
	Conditioner       bool   `json:"conditioner"`
	Fridge            bool   `json:"fridge"`
	DishWasher        bool   `json:"dish_washer"`
	Laundry           bool   `json:"laundry"`
	Internet          bool   `json:"internet"`
	Tv                bool   `json:"tv"`
}

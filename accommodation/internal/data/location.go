package data

type Location struct {
	Id       int64    `json:"id"`
	Long     float64  `json:"long"`
	Lat      float64  `json:"lat"`
	Street   string   `json:"street"`
	House    int64    `json:"house"`
	Building string   `json:"building"`
	Flat     int64    `json:"flat"`
	Subway   []Subway `json:"subway"`
}

type Subway struct {
	StationName   string `json:"station_name"`
	LineColor     string `json:"line_color"`
	TransportType string `json:"transport_type"`
	TransportTime int64  `json:"transport_time"`
}

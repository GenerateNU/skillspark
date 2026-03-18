package models

type GeocodeAddressInput struct {
	Body struct {
		Address string `json:"address" doc:"Text address to geocode" minLength:"1"`
	}
}

type GeocodeAddressOutput struct {
	Body struct {
		Latitude  float64 `json:"latitude" doc:"Latitude of the geocoded address"`
		Longitude float64 `json:"longitude" doc:"Longitude of the geocoded address"`
	}
}

package models

import "time"

type GeocodeCache struct {
	Address    string    `json:"address" db:"address"`
	RawAddress string    `json:"raw_address" db:"raw_address"`
	Latitude   float64   `json:"latitude" db:"latitude"`
	Longitude  float64   `json:"longitude" db:"longitude"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

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

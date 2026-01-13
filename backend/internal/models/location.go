package models

import (
	"time"

	"github.com/google/uuid"
)

// The database model for a location
type Location struct {
	ID               uuid.UUID `json:"id" db:"id"`
	Latitude         float64   `json:"latitude" db:"latitude"`
	Longitude        float64   `json:"longitude" db:"longitude"`
	StreetNumber     string    `json:"street_number" db:"street_number"`
	StreetName       string    `json:"street_name" db:"street_name"`
	SecondaryAddress *string   `json:"secondary_address" db:"secondary_address"`
	City             string    `json:"city" db:"city"`
	State            string    `json:"state" db:"state"`
	PostalCode       string    `json:"postal_code" db:"postal_code"`
	Country          string    `json:"country" db:"country"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
}

type GetLocationByIDInput struct {
	ID uuid.UUID `path:"id"`
}

type CreateLocationInput struct {
	Body struct {
		Latitude float64 `json:"latitude" db:"latitude" doc:"Latitude of the location" minimum:"-90" maximum:"90"`

		Longitude float64 `json:"longitude" db:"longitude" doc:"Longitude of the location" minimum:"-180" maximum:"180"`

		StreetNumber string `json:"street_number" db:"street_number" doc:"Street number of the location" minLength:"1" maxLength:"20"`

		StreetName string `json:"street_name" db:"street_name" doc:"Street name of the location" minLength:"2" maxLength:"100"`

		SecondaryAddress *string `json:"secondary_address" db:"secondary_address" doc:"Secondary address of the location" minLength:"0" maxLength:"100"`

		City string `json:"city" db:"city" doc:"City of the location" minLength:"2" maxLength:"100"`

		State string `json:"state" db:"state" doc:"State of the location" minLength:"2" maxLength:"50"`

		PostalCode string `json:"postal_code" db:"postal_code" doc:"Postal code of the location" minLength:"3" maxLength:"20"`

		Country string `json:"country" db:"country" doc:"Country of the location" minLength:"2" maxLength:"100"`
	}
}

type GetLocationByIDOutput struct {
	Body *Location `json:"body"` // Huma will serialize this as JSON
}

type CreateLocationOutput struct {
	Body *Location `json:"body"`
}

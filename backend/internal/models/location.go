package models

import (
	"time"

	"github.com/google/uuid"
)

// The database model for a location
type Location struct {
	ID           uuid.UUID `json:"id" db:"id"`
	Latitude     float64   `json:"latitude" db:"latitude"`
	Longitude    float64   `json:"longitude" db:"longitude"`
	AddressLine1 string    `json:"address_line1" db:"address_line1"`
	AddressLine2 *string   `json:"address_line2" db:"address_line2"`
	Subdistrict  string    `json:"subdistrict" db:"subdistrict"`
	District     string    `json:"district" db:"district"`
	Province     string    `json:"province" db:"province"`
	PostalCode   string    `json:"postal_code" db:"postal_code"`
	Country      string    `json:"country" db:"country"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type GetLocationByIDInput struct {
	ID uuid.UUID `path:"id"`
}

type CreateLocationInput struct {
	Body struct {
		Latitude float64 `json:"latitude" db:"latitude" doc:"Latitude of the location" minimum:"-90" maximum:"90"`

		Longitude float64 `json:"longitude" db:"longitude" doc:"Longitude of the location" minimum:"-180" maximum:"180"`

		AddressLine1 string `json:"address_line1" db:"address_line1" doc:"Primary address line of the location" minLength:"5" maxLength:"200"`

		AddressLine2 *string `json:"address_line2,omitempty" db:"address_line2" doc:"Secondary address line of the location" minLength:"5" maxLength:"200"`

		Subdistrict string `json:"subdistrict" db:"subdistrict" doc:"Subdistrict of the location" minLength:"2" maxLength:"100"`

		District string `json:"district" db:"district" doc:"District of the location" minLength:"2" maxLength:"100"`

		Province string `json:"province" db:"province" doc:"Province/State of the location" minLength:"2" maxLength:"100"`

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

type GetAllLocationsInput struct {
	Page  int `query:"page" minimum:"1" default:"1"`
	Limit int `query:"limit" minimum:"1" maximum:"100" default:"100"`
}

type GetAllLocationsOutput struct {
	Body []Location `json:"body" doc:"List of all locations in the database"`
}

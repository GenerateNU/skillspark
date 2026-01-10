package models

import (
	"time"

	"github.com/google/uuid"
)

// The database model for a location
type Location struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Latitude  float64   `json:"latitude" db:"latitude"`
	Longitude float64   `json:"longitude" db:"longitude"`
	Address   string    `json:"address" db:"address"`
	City      string    `json:"city" db:"city"`
	State     string    `json:"state" db:"state"`
	ZipCode   string    `json:"zip_code" db:"zip_code"`
	Country   string    `json:"country" db:"country"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type GetLocationByIDInput struct {
	ID uuid.UUID `path:"id"`
}

type CreateLocationInput struct {
	Body struct {
		Latitude float64 `json:"latitude" db:"latitude" doc:"Latitude of the location" validate:"required,gte=-90,lte=90"`

		Longitude float64 `json:"longitude" db:"longitude" doc:"Longitude of the location" validate:"required,gte=-180,lte=180"`

		Address string `json:"address" db:"address" doc:"Street address of the location" validate:"required,min=3,max=255"`

		City string `json:"city" db:"city" doc:"City of the location" validate:"required,min=2,max=100"`

		State string `json:"state" db:"state" doc:"State of the location" validate:"required,min=2,max=50"`

		ZipCode string `json:"zip_code" db:"zip_code" doc:"ZIP code of the location" validate:"required,min=3,max=20"`

		Country string `json:"country" db:"country" doc:"Country of the location" validate:"required,min=2,max=100"`
	}
}

type GetLocationByIDOutput struct {
	Body *Location `json:"body"` // Huma will serialize this as JSON
}

type CreateLocationOutput struct {
	Body *Location `json:"body"`
}

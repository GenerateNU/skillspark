package models

import (
	"time"

	"github.com/google/uuid"
)

// The database model for a location
type Location struct {
	ID             uuid.UUID `json:"id" db:"id"`
	Latitude       float64   `json:"latitude" db:"latitude"`
	Longitude      float64   `json:"longitude" db:"longitude"`
	OrganizationID uuid.UUID `json:"organization_id" db:"organization_id"`
	Address        string    `json:"address" db:"address"`
	City           string    `json:"city" db:"city"`
	State          string    `json:"state" db:"state"`
	ZipCode        string    `json:"zip_code" db:"zip_code"`
	Country        string    `json:"country" db:"country"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

type GetLocationByIDInput struct {
	ID uuid.UUID `path:"id"`
}

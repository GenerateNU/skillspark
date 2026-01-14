package models

import (
	"time"

	"github.com/google/uuid"
)

type School struct {
	ID         uuid.UUID `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	LocationID uuid.UUID `json:"location_id" db:"location_id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

type GetAllSchoolsOutput struct {
	Body []School `json:"body" doc:"List of all schools in the database"`
}

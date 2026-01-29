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

type GetAllSchoolsInput struct {
	Page  int `query:"page" minimum:"1" default:"1"`
	Limit int `query:"limit" minimum:"1" maximum:"100" default:"100"`
}

type CreateSchoolInput struct {
	Body struct {
		Name       string    `json:"name" db:"name" minLength:"1" maxLength:"200"`
		LocationID uuid.UUID `json:"location_id" db:"location_id"`
	}
}

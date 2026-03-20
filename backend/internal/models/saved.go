package models

import (
	"time"

	"github.com/google/uuid"
)

type Saved struct {
	ID         uuid.UUID `json:"id" db:"id"`
	GuardianID uuid.UUID `json:"guardian_id" db:"guardian_id"`
	Event      Event     `json:"event" db:"-"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

type CreateSavedInput struct {
	Body struct {
		GuardianID uuid.UUID `json:"guardian_id" db:"guardian_id" doc:"ID of the guardian that saved this."`
		EventID    uuid.UUID `json:"event_id" db:"event_id" doc:"ID of this saved event."`
	}
}

type CreateSavedOutput struct {
	Body Saved
}

type DeleteSavedInput struct {
	ID uuid.UUID `path:"id"`
}

type DeleteSavedOutput struct {
	Body struct {
		Message string `json:"message" doc:"Success message"`
	} `json:"body"`
}

type GetSavedInput struct {
	ID       uuid.UUID `path:"id"`
	Page     int       `query:"page" minimum:"1" default:"1" doc:"Page number (starts at 1)"`
	PageSize int       `query:"page_size" minimum:"1" maximum:"100" default:"10" doc:"Number of items per page"`
}

type GetSavedOutput struct {
	Body []Saved `json:"body" doc:"List of saved event occurrences"`
}

package models

import (
	"time"

	"github.com/google/uuid"
)

type Saved struct {
	ID                uuid.UUID `json:"id" db:"id"`
	GuardianID        uuid.UUID `json:"guardian_id" db:"guardian_id"`
	EventOccurrenceID uuid.UUID `json:"event_occurrence_id" db:"event_occurrence_id"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
}

type CreateSavedInput struct {
	Body struct {
		GuardianID        uuid.UUID `json:"guardian_id" db:"guardian_id" doc:"ID of the guardian that saved this."`
		EventOccurrenceID uuid.UUID `json:"event_occurrence_id" db:"event_occurrence_id" doc:"ID of the event occurrence of this saved event."`
	}
}

type DeleteSavedInput struct {
	ID uuid.UUID `path:"id"`
}

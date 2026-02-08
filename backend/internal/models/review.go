package models

import (
	"time"

	"github.com/google/uuid"
)

type Review struct {
	ID             uuid.UUID `json:"id" db:"id" doc:"Unique review identifier"`
	RegistrationID uuid.UUID `json:"registration_id" db:"registration_id" doc:"ID of the linked registration"`
	GuardianID     uuid.UUID `json:"guardian_id" db:"guardian_id" doc:"ID of the guardian"`
	Description    string    `json:"description" db:"description" doc:"The review text"`
	Categories     []string  `json:"categories" db:"categories" doc:"Review categories for this review, can be one of fun, engaging, interesting or informative."`
	CreatedAt      time.Time `json:"created_at" db:"created_at" doc:"Timestamp when registration was created"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at" doc:"Timestamp when registration was last updated"`
}

type CreateReviewInput struct {
	Body struct {
		RegistrationID uuid.UUID `json:"registration_id" db:"registration_id" doc:"ID of the linked registration"`
		GuardianID     uuid.UUID `json:"guardian_id" db:"guardian_id" doc:"ID of the guardian"`
		Description    string    `json:"description" db:"description" doc:"The review text"`
		Categories     []string  `json:"categories" db:"categories" doc:"Review categories for this review, can be one of fun, engaging, interesting or informative."`
	}
}

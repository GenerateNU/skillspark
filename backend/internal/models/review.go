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

type CreateReviewDBBody struct {
	RegistrationID uuid.UUID `json:"registration_id" db:"registration_id" doc:"ID of the linked registration"`
	GuardianID     uuid.UUID `json:"guardian_id" db:"guardian_id" doc:"ID of the guardian"`
	Description_EN string    `json:"description_en" db:"description_en" doc:"The review text"`
	Description_TH string    `json:"description_th" db:"description_th" doc:"The review text"`
	Categories_EN  []string  `json:"categories_en" db:"categories_en" doc:"Review categories for this review, can be one of fun, engaging, interesting or informative."`
	Categories_TH  []string  `json:"categories_th" db:"categories_th" doc:"Review categories for this review, can be one of fun, engaging, interesting or informative."`
}

type CreateReviewDBInput struct {
	Body CreateReviewDBBody
}

type CreateReviewInput struct {
	Body struct {
		RegistrationID uuid.UUID `json:"registration_id" db:"registration_id" doc:"ID of the linked registration"`
		GuardianID     uuid.UUID `json:"guardian_id" db:"guardian_id" doc:"ID of the guardian"`
		Description    string    `json:"description" db:"description" doc:"The review text"`
		Categories     []string  `json:"categories" db:"categories" doc:"Review categories for this review, can be one of fun, engaging, interesting or informative."`
	}
}

type GetReviewsByGuardianIDInput struct {
	ID       uuid.UUID `path:"id"`
	Page     int       `query:"page" minimum:"1" default:"1" doc:"Page number (starts at 1)"`
	PageSize int       `query:"page_size" minimum:"1" maximum:"100" default:"10" doc:"Number of items per page"`
}

type GetReviewsByEventIDInput struct {
	ID       uuid.UUID `path:"id"`
	Page     int       `query:"page" minimum:"1" default:"1" doc:"Page number (starts at 1)"`
	PageSize int       `query:"page_size" minimum:"1" maximum:"100" default:"10" doc:"Number of items per page"`
}

type DeleteReviewInput struct {
	ID uuid.UUID `path:"id"`
}

type ReviewsOutput struct {
	Body []Review `json:"body" doc:"List of reviews"`
}

type DeleteReviewOutput struct {
	Body struct {
		Message string `json:"message" doc:"Success message"`
	} `json:"body"`
}

type CreateReviewOutput struct {
	Body Review
}

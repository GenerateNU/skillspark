package models

import (
	"time"

	"github.com/google/uuid"
)

// The database model for an event
type Event struct {
	ID               uuid.UUID `json:"id" db:"id"`
	Title            string    `json:"title" db:"title"`
	Description      string    `json:"description" db:"description"`
	OrganizationID   uuid.UUID `json:"organization_id" db:"organization_id"`
	AgeRangeMin      *int      `json:"age_range_min" db:"age_range_min"`
	AgeRangeMax      *int      `json:"age_range_max" db:"age_range_max"`
	Category         []string  `json:"category" db:"category"`
	HeaderImageS3Key *string   `json:"header_image_s3_key" db:"header_image_s3_key"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
}

type CreateEventInput struct {
	Body struct {
		Title string `json:"title" db:"title" doc:"Title of the event" minLength:"2" maxLength:"100"`

		Description string `json:"description" db:"description" doc:"Description of the event" minLength:"2" maxLength:"200"`

		OrganizationID uuid.UUID `json:"organization_id" db:"organization_id" doc:"ID of the hosting organization"`

		AgeRangeMin *int `json:"age_range_min" db:"age_range_min" doc:"Minimum age for the event" minimum:"0" maximum:"100"`

		AgeRangeMax *int `json:"age_range_max" db:"age_range_max" doc:"Max age for the event" minimum:"0" maximum:"100"`

		Category []string `json:"category" db:"category" doc:"Category of the event"`

		HeaderImageS3Key *string `json:"header_image_s3_key" db:"header_image_s3_key" doc:"S3 key for the header image"`
	}
}

type CreateEventOutput struct {
	Body *Event `json:"body"`
}

type PatchEventInput struct {
	ID   uuid.UUID `path:"id"`
	Body struct {
		Title string `json:"title" db:"title" doc:"Title of the event" minLength:"2" maxLength:"100"`

		Description string `json:"description" db:"description" doc:"Description of the event" minLength:"2" maxLength:"200"`

		OrganizationID uuid.UUID `json:"organization_id" db:"organization_id" doc:"ID of the hosting organization"`

		AgeRangeMin *int `json:"age_range_min" db:"age_range_min" doc:"Minimum age for the event" minimum:"0" maximum:"100"`

		AgeRangeMax *int `json:"age_range_max" db:"age_range_max" doc:"Max age for the event" minimum:"0" maximum:"100"`

		Category []string `json:"category" db:"category" doc:"Category of the event"`

		HeaderImageS3Key *string `json:"header_image_s3_key" db:"header_image_s3_key" doc:"S3 key for the header image"`
	}
}

type PatchEventOutput struct {
	Body *Event `json:"body"`
}

type DeleteEventInput struct {
	ID uuid.UUID `path:"id"`
}

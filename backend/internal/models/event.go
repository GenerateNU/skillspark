package models

import (
	"mime/multipart"
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

		// changed from HeaderImageS3Key to File datatype
		HeaderImageS3Key *string `json:"header_image_s3_key" db:"header_image_s3_key" doc:"S3 key for the header image"`
	}
}

// CreateEventRouteInput is the multipart form input for creating an event with an image
type CreateEventRouteInput struct {
	RawBody multipart.Form
}

// CreateEventFormData holds the parsed form data for creating an event
type CreateEventFormData struct {
	Title             string
	Description       string
	OrganizationID    uuid.UUID
	AgeRangeMin       *int
	AgeRangeMax       *int
	Category          []string
	HeaderImage       multipart.File        // The actual file
	HeaderImageHeader *multipart.FileHeader // File metadata (filename, size, etc.)
}

type CreateEventOutput struct {
	Body *Event `json:"body"`
}

type UpdateEventBody struct {
	Title            *string    `json:"title,omitempty" db:"title" doc:"Title of the event" minLength:"2" maxLength:"100"`
	Description      *string    `json:"description,omitempty" db:"description" doc:"Description of the event" minLength:"2" maxLength:"200"`
	OrganizationID   *uuid.UUID `json:"organization_id,omitempty" db:"organization_id" doc:"ID of the hosting organization"`
	AgeRangeMin      *int       `json:"age_range_min,omitempty" db:"age_range_min" doc:"Minimum age for the event" minimum:"0" maximum:"100"`
	AgeRangeMax      *int       `json:"age_range_max,omitempty" db:"age_range_max" doc:"Max age for the event" minimum:"0" maximum:"100"`
	Category         *[]string  `json:"category,omitempty" db:"category" doc:"Category of the event"`
	HeaderImageS3Key *string    `json:"header_image_s3_key,omitempty" db:"header_image_s3_key" doc:"S3 key for the header image"`
}

type UpdateEventInput struct {
	ID   uuid.UUID `path:"id"`
	Body struct {
		UpdateEventBody
	}
}

type UpdateEventOutput struct {
	Body *Event `json:"body"`
}

type DeleteEventInput struct {
	ID uuid.UUID `path:"id"`
}

type DeleteEventOutput struct {
	Body struct {
		Message string `json:"message" doc:"Success message"`
	} `json:"body"`
}

// get by event id
type GetEventOccurrencesByEventIDInput struct {
	ID uuid.UUID `path:"event_id" doc:"ID of an event"`
}

type GetEventOccurrencesByEventIDOutput struct {
	Body []EventOccurrence `json:"body" doc:"List of event occurrences in the database that match the event ID"`
}

package models

import (
	"time"

	"github.com/danielgtaylor/huma/v2"
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
	PresignedURL     *string   `json:"presigned_url"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
}

type CreateDBBody struct {
	Title_EN       string    `json:"title_en,omitempty" db:"title_en" doc:"Title of the event in english" minLength:"2" maxLength:"100"`
	Title_TH       *string   `json:"title_th,omitempty" db:"title_th" doc:"Title of the event in thai" minLength:"2" maxLength:"100"`
	Description_EN string    `json:"description_en,omitempty" db:"description_en" doc:"Description of the event in english" minLength:"2" maxLength:"200"`
	Description_TH *string   `json:"description_th,omitempty" db:"description_th" doc:"Description of the event in thai" minLength:"2" maxLength:"200"`
	OrganizationID uuid.UUID `json:"organization_id,omitempty" db:"organization_id" doc:"ID of the hosting organization"`
	AgeRangeMin    *int      `json:"age_range_min,omitempty" db:"age_range_min" doc:"Minimum age for the event" minimum:"0" maximum:"100"`
	AgeRangeMax    *int      `json:"age_range_max,omitempty" db:"age_range_max" doc:"Max age for the event" minimum:"0" maximum:"100"`
	Category       []string  `json:"category,omitempty" db:"category" doc:"Category of the event"`
}

type UpdateDBBody struct {
	Title_EN       *string    `json:"title_en,omitempty" db:"title_en" doc:"Title of the event in english" minLength:"2" maxLength:"100"`
	Title_TH       *string    `json:"title_th,omitempty" db:"title_th" doc:"Title of the event in thai" minLength:"2" maxLength:"100"`
	Description_EN *string    `json:"description_en,omitempty" db:"description_en" doc:"Description of the event in english" minLength:"2" maxLength:"200"`
	Description_TH *string    `json:"description_th,omitempty" db:"description_th" doc:"Description of the event in thai" minLength:"2" maxLength:"200"`
	OrganizationID *uuid.UUID `json:"organization_id,omitempty" db:"organization_id" doc:"ID of the hosting organization"`
	AgeRangeMin    *int       `json:"age_range_min,omitempty" db:"age_range_min" doc:"Minimum age for the event" minimum:"0" maximum:"100"`
	AgeRangeMax    *int       `json:"age_range_max,omitempty" db:"age_range_max" doc:"Max age for the event" minimum:"0" maximum:"100"`
	Category       *[]string  `json:"category,omitempty" db:"category" doc:"Category of the event"`
}

type CreateEventDBInput struct {
	Body CreateDBBody
}

type UpdateEventDBInput struct {
	ID   uuid.UUID `path:"id"`
	Body UpdateDBBody
}

// ----------------------------

type CreateEventInput struct {
	Body CreateEventBody
}

// CreateEventRouteInput is the multipart form input for creating an event with an image
type CreateEventRouteInput struct {
	RawBody huma.MultipartFormFiles[CreateEventFormData]
}

// CreateEventFormData holds the parsed form data for creating an event
// make agerangemin, max and headerimage optional somehow
type CreateEventFormData struct {
	Title          string        `form:"title" required:"true" minLength:"2" maxLength:"100"`
	Description    string        `form:"description" required:"true" minLength:"2" maxLength:"200"`
	OrganizationID uuid.UUID     `form:"organization_id" required:"true"`
	AgeRangeMin    int           `form:"age_range_min"`
	AgeRangeMax    int           `form:"age_range_max"`
	Category       []string      `form:"category"`
	HeaderImage    huma.FormFile `form:"header_image" contentType:"image/png,image/jpeg"`
}

type UpdateEventFormData struct {
	Title          string        `form:"title" minLength:"2" maxLength:"100"`
	Description    string        `form:"description" minLength:"2" maxLength:"200"`
	OrganizationID uuid.UUID     `form:"organization_id"`
	AgeRangeMin    int           `form:"age_range_min"`
	AgeRangeMax    int           `form:"age_range_max"`
	Category       []string      `form:"category"`
	HeaderImage    huma.FormFile `form:"header_image" contentType:"image/png,image/jpeg"`
}

type CreateEventOutput struct {
	Body *Event `json:"body"`
}

type UpdateEventRouteInput struct {
	ID      uuid.UUID `path:"id"`
	RawBody huma.MultipartFormFiles[UpdateEventFormData]
}

type CreateEventBody struct {
	Title          string    `json:"title,omitempty" db:"title" doc:"Title of the event" minLength:"2" maxLength:"100"`
	Description    string    `json:"description,omitempty" db:"description" doc:"Description of the event" minLength:"2" maxLength:"200"`
	OrganizationID uuid.UUID `json:"organization_id,omitempty" db:"organization_id" doc:"ID of the hosting organization"`
	AgeRangeMin    *int      `json:"age_range_min,omitempty" db:"age_range_min" doc:"Minimum age for the event" minimum:"0" maximum:"100"`
	AgeRangeMax    *int      `json:"age_range_max,omitempty" db:"age_range_max" doc:"Max age for the event" minimum:"0" maximum:"100"`
	Category       []string  `json:"category,omitempty" db:"category" doc:"Category of the event"`
}

type UpdateEventBody struct {
	Title          *string    `json:"title,omitempty" db:"title" doc:"Title of the event" minLength:"2" maxLength:"100"`
	Description    *string    `json:"description,omitempty" db:"description" doc:"Description of the event" minLength:"2" maxLength:"200"`
	OrganizationID *uuid.UUID `json:"organization_id,omitempty" db:"organization_id" doc:"ID of the hosting organization"`
	AgeRangeMin    *int       `json:"age_range_min,omitempty" db:"age_range_min" doc:"Minimum age for the event" minimum:"0" maximum:"100"`
	AgeRangeMax    *int       `json:"age_range_max,omitempty" db:"age_range_max" doc:"Max age for the event" minimum:"0" maximum:"100"`
	Category       *[]string  `json:"category,omitempty" db:"category" doc:"Category of the event"`
}

type UpdateEventInput struct {
	ID   uuid.UUID `path:"id"`
	Body UpdateEventBody
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
	AcceptLanguage string    `header:"Accept-Language"`
	ID             uuid.UUID `path:"event_id" doc:"ID of an event"`
}

type GetEventOccurrencesByEventIDOutput struct {
	AcceptLanguage string            `header:"Accept-Language"`
	Body           []EventOccurrence `json:"body" doc:"List of event occurrences in the database that match the event ID"`
}

package models

import (
	"time"

	"github.com/google/uuid"
)

type EventOccurrenceStatus string

const (
	EventOccurrenceStatusScheduled EventOccurrenceStatus = "scheduled"
	EventOccurrenceStatusCancelled EventOccurrenceStatus = "cancelled"
)

type EventOccurrence struct {
	ID           uuid.UUID             `json:"id" db:"id"`
	ManagerId    *uuid.UUID            `json:"manager_id" db:"manager_id"`
	Event        Event                 `json:"event" db:"-"`
	Location     Location              `json:"location" db:"-"`
	StartTime    time.Time             `json:"start_time" db:"start_time"`
	EndTime      time.Time             `json:"end_time" db:"end_time"`
	MaxAttendees int                   `json:"max_attendees" db:"max_attendees"`
	Language     string                `json:"language" db:"language"`
	CurrEnrolled int                   `json:"curr_enrolled" db:"curr_enrolled"`
	Price        int                   `json:"price" db:"price" doc:"Price in cents (e.g., 10000 = ฿100)"`
	CreatedAt    time.Time             `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time             `json:"updated_at" db:"updated_at"`
	Status       EventOccurrenceStatus `json:"status" db:"status" doc:"Current status of the event occurrence" enum:"scheduled,cancelled"`
}

type GetAllEventOccurrencesInput struct {
	Page  int `query:"page" minimum:"1" default:"1"`
	Limit int `query:"limit" minimum:"1" maximum:"100" default:"100"`
}

type GetAllEventOccurrencesOutput struct {
	Body []EventOccurrence `json:"body" doc:"List of all event occurrences in the database"`
}

type GetEventOccurrenceByIDInput struct {
	ID uuid.UUID `path:"id" doc:"ID of an event occurrence"`
}

type GetEventOccurrenceByIDOutput struct {
	Body *EventOccurrence `json:"body" doc:"Event occurrence in the database that matches the ID"`
}

type CreateEventOccurrenceInput struct {
	Body struct {
		ManagerId    *uuid.UUID `json:"manager_id,omitempty" doc:"ID of a manager in the database"`
		EventId      uuid.UUID  `json:"event_id" doc:"ID of an event in the database"`
		LocationId   uuid.UUID  `json:"location_id" doc:"ID of a location in the database"`
		StartTime    time.Time  `json:"start_time" doc:"Start time of the event occurrence"`
		EndTime      time.Time  `json:"end_time" doc:"End time of the event occurrence"`
		MaxAttendees int        `json:"max_attendees" doc:"Maximum number of attendees" minimum:"1" maximum:"100"`
		Language     string     `json:"language" doc:"Primary language used for the event occurrence" minLength:"2" maxLength:"30"`
		Price        int        `json:"price" doc:"Price in cents (e.g., 10000 = ฿100)" minimum:"0"`
	} `json:"body" doc:"New event occurrence to add"`
}

type CreateEventOccurrenceOutput struct {
	Body *EventOccurrence `json:"body" doc:"Created event occurrence"`
}

type UpdateEventOccurrenceInput struct {
	ID   uuid.UUID `path:"id" doc:"ID of the event occurrence to update"`
	Body struct {
		ManagerId    *uuid.UUID `json:"manager_id,omitempty" doc:"ID of a manager in the database"`
		EventId      *uuid.UUID `json:"event_id,omitempty" doc:"ID of an event in the database"`
		LocationId   *uuid.UUID `json:"location_id,omitempty" doc:"ID of a location in the database"`
		StartTime    *time.Time `json:"start_time,omitempty" doc:"Start time of the event occurrence"`
		EndTime      *time.Time `json:"end_time,omitempty" doc:"End time of the event occurrence"`
		MaxAttendees *int       `json:"max_attendees,omitempty" doc:"Maximum number of attendees" minimum:"1" maximum:"100"`
		Language     *string    `json:"language,omitempty" doc:"Primary language used for the event occurrence" minLength:"2" maxLength:"30"`
		CurrEnrolled *int       `json:"curr_enrolled,omitempty" doc:"Number of students currently enrolled in the event occurrence" minimum:"0" maximum:"100"`
		Price        *int       `json:"price,omitempty" doc:"Price in cents" minimum:"0"`
	} `json:"body" doc:"Event occurrence fields to update"`
}

type UpdateEventOccurrenceOutput struct {
	Body *EventOccurrence `json:"body" doc:"Updated event occurrence"`
}

type CancelEventOccurrenceInput struct {
	ID uuid.UUID `path:"id" doc:"ID of an event occurrence"`
}

type CancelEventOccurrenceOutput struct {
	Body struct {
		Message string `json:"message" doc:"Success message"`
	} `json:"body"`
}
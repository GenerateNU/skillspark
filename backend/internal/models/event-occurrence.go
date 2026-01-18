package models

import (
	"time"

	"github.com/google/uuid"
)

// database model for a specific instance of an event 
type EventOccurrence struct {
	ID 				uuid.UUID 	`json:"id" db:"id"`
	ManagerId		uuid.UUID 	`json:"manager(id)"`   
	EventId 		uuid.UUID 	`json:"event_id"`
 	LocationId 		uuid.UUID 	`json:"location_id" db:"location_id"`
   	StartTime 		time.Time 	`json:"start_time" db:"start_time"`
   	EndTime 		time.Time 	`json:"end_time" db:"end_time"`
  	MaxAttendees 	int			`json:"max_attendees" db:"max_attendees"`
   	Language 		string		`json:"language" db:"language"`  
	CurrEnrolled 	int			`json:"curr_enrolled" db:"curr_enrolled"`
  	CreatedAt 		time.Time 	`json:"created_at" db:"created_at"`
  	UpdatedAt  		time.Time 	`json:"updated_at1" db:"updated_at"`
}

// get all
type GetAllEventOccurrencesOutput struct {
	Body []EventOccurrence `json:"body" doc:"List of all event occurrences in the database"`
}

// get by event occurrence id
type GetEventOccurrenceByIDInput struct {
	ID uuid.UUID `path:"id" doc:"ID of an event occurrence"`
}

type GetEventOccurrencesByIDOutput struct {
	Body []EventOccurrence `json:"body" doc:"List of event occurrences in the database from an ID"`
}

// get by event id
type GetEventOccurrenceByEventIDInput struct {
	EventID uuid.UUID `path:"event_id" doc:"ID of an event"`
}

type GetEventOccurrencesByEventIDOutput struct {
	Body []EventOccurrence `json:"body" doc:"List of event occurrences in the database from an event ID"`
}

// post
type PostEventOccurrencesInput struct {
	Body []EventOccurrence `json:"body" doc:"List of new event occurrences to add"`
}
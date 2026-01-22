package models

import (
	"time"

	"github.com/google/uuid"
)

// database model for a specific instance of an event 
// stores full type information for Event and Location
type EventOccurrence struct {
	ID 				uuid.UUID 	`json:"id" db:"id"`
	ManagerId		*uuid.UUID 	`json:"manager_id" db:"manager_id"`
	Event	 		Event 		`json:"event" db:"-"`
 	Location 		Location 	`json:"location" db:"-"`
   	StartTime 		time.Time 	`json:"start_time" db:"start_time"`
   	EndTime 		time.Time 	`json:"end_time" db:"end_time"`
  	MaxAttendees 	int			`json:"max_attendees" db:"max_attendees"`
   	Language 		string		`json:"language" db:"language"`  
	CurrEnrolled 	int			`json:"curr_enrolled" db:"curr_enrolled"`
  	CreatedAt 		time.Time 	`json:"created_at" db:"created_at"`
  	UpdatedAt  		time.Time 	`json:"updated_at" db:"updated_at"`
}

// get all
type GetAllEventOccurrencesInput struct {
	Page int `query:"page" minimum:"1" default:"1"`
	Limit int `query:"limit" minimum:"1" maximum:"100" default:"100"`
}

type GetAllEventOccurrencesOutput struct {
	Body []EventOccurrence `json:"body" doc:"List of all event occurrences in the database"`
}

// get by event occurrence id
type GetEventOccurrenceByIDInput struct {
	ID uuid.UUID `path:"id" doc:"ID of an event occurrence"`
}

type GetEventOccurrenceByIDOutput struct {
	Body *EventOccurrence `json:"body" doc:"Event occurrence in the database that matches the ID"`
}

// get by event id
type GetEventOccurrencesByEventIDInput struct {
	ID uuid.UUID `path:"event_id" doc:"ID of an event"`
}

type GetEventOccurrencesByEventIDOutput struct {
	Body []EventOccurrence `json:"body" doc:"List of event occurrences in the database that match the event ID"`
}

// post
type CreateEventOccurrenceInput struct {
	Body struct {
		ManagerId *uuid.UUID `json:"manager_id, omitempty" doc:"ID of a manager in the database"`
		EventId uuid.UUID `json:"event_id" doc:"ID of an event in the database"`
		LocationId uuid.UUID `json:"location_id" doc:"ID of a location in the database"`
		StartTime time.Time `json:"start_time" doc:"Start time of the event occurrence"`
		EndTime time.Time `json:"end_time" doc:"End time of the event occurrence"`
		MaxAttendees int `json:"max_attendees" doc:"Maximum number of attendees" minimum:"1" maximum:"100"`
		Language string	`json:"language" doc:"Primary language used for the event occurrence" minLength:"2" maxLength:"30"`  
	} `json:"body" doc:"New event occurrence to add"`
}

type CreateEventOccurrenceOutput struct {
	Body *EventOccurrence `json:"body" doc:"Created event occurrence"`
}
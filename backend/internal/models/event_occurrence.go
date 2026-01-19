package models

import (
	"time"

	"github.com/google/uuid"
)

// database model for a specific instance of an event 
type EventOccurrence struct {
	ID 				uuid.UUID 	`json:"id" db:"id"`
	ManagerId		uuid.UUID 	`json:"manager_id" db:"manager_id"`
	Event	 		Event 		`json:"event"`
 	Location 		Location 	`json:"location"`
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
	ID uuid.UUID `path:"id" doc:"ID of an event"`
}

type GetEventOccurrencesByEventIDOutput struct {
	Body []EventOccurrence `json:"body" doc:"List of event occurrences in the database that match the event ID"`
}

// post
type CreateEventOccurrenceInput struct {
	Body struct {
		ManagerId uuid.UUID `json:"manager_id"`
		EventId uuid.UUID `json:"event"`
		LocationId uuid.UUID `json:"location"`
		StartTime time.Time `json:"start_time"`
		EndTime time.Time `json:"end_time"`
		MaxAttendees int `json:"max_attendees"`
		Language string	`json:"language"`  
	} `json:"body" doc:"New event occurrence to add"`
}

type CreateEventOccurrenceOutput struct {
	Body *EventOccurrence `json:"body"`
}
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
		Event struct {
			Title string
			Description string
			OrganizationID uuid.UUID
			AgeRangeMin int
			AgeRangeMax int
			Category string
			HeaderImageS3Key string
		} `json:"event"`
		Location struct {
			Latitude float64 `json:"latitude" db:"latitude" doc:"Latitude of the location" minimum:"-90" maximum:"90"`
			Longitude float64 `json:"longitude" db:"longitude" doc:"Longitude of the location" minimum:"-180" maximum:"180"`
			AddressLine1 string `json:"address_line1" db:"address_line1" doc:"Primary address line of the location" minLength:"5" maxLength:"200"`
			AddressLine2 *string `json:"address_line2" db:"address_line2" doc:"Secondary address line of the location" minLength:"5" maxLength:"200"`
			Subdistrict string `json:"subdistrict" db:"subdistrict" doc:"Subdistrict of the location" minLength:"2" maxLength:"100"`
			District string `json:"district" db:"district" doc:"District of the location" minLength:"2" maxLength:"100"`
			Province string `json:"province" db:"province" doc:"Province/State of the location" minLength:"2" maxLength:"100"`
			PostalCode string `json:"postal_code" db:"postal_code" doc:"Postal code of the location" minLength:"3" maxLength:"20"`
			Country string `json:"country" db:"country" doc:"Country of the location" minLength:"2" maxLength:"100"`
		} `json:"location"`
		StartTime time.Time `json:"start_time"`
		EndTime time.Time `json:"end_time"`
		MaxAttendees int `json:"max_attendees"`
		Language string	`json:"language"`  
		CurrEnrolled int `json:"curr_enrolled"`
	} `json:"body" doc:"New event occurrence to add"`
}

type CreateEventOccurrenceOutput struct {
	Body *EventOccurrence `json:"body"`
}
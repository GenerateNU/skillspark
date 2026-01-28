package models

import (
	"time"

	"github.com/google/uuid"
)

type Registration struct {
	ID                  uuid.UUID          `json:"id" db:"id" doc:"Unique registration identifier"`
	ChildID             uuid.UUID          `json:"child_id" db:"child_id" doc:"ID of the registered child"`
	GuardianID          uuid.UUID          `json:"guardian_id" db:"guardian_id" doc:"ID of the child's guardian"`
	EventOccurrenceID   uuid.UUID          `json:"event_occurrence_id" db:"event_occurrence_id" doc:"ID of the event occurrence"`
	Status              RegistrationStatus `json:"status" db:"status" doc:"Current status of the registration" enum:"registered,cancelled"`
	CreatedAt           time.Time          `json:"created_at" db:"created_at" doc:"Timestamp when registration was created"`
	UpdatedAt           time.Time          `json:"updated_at" db:"updated_at" doc:"Timestamp when registration was last updated"`
	EventName           string             `json:"event_name" db:"event_name" doc:"Name of the event (joined from events table)"`
	OccurrenceStartTime time.Time          `json:"occurrence_start_time" db:"occurrence_start_time" doc:"Start time of the event occurrence"`
}

type RegistrationStatus string

const (
	RegistrationStatusRegistered RegistrationStatus = "registered"
	RegistrationStatusCancelled  RegistrationStatus = "cancelled"
)

func (rs RegistrationStatus) IsValid() bool {
	return rs == RegistrationStatusRegistered || rs == RegistrationStatusCancelled
}

type CreateRegistrationInput struct {
	Body struct {
		ChildID           uuid.UUID          `json:"child_id" doc:"ID of the child to register" format:"uuid" required:"true"`
		GuardianID        uuid.UUID          `json:"guardian_id" doc:"ID of the guardian registering the child" format:"uuid" required:"true"`
		EventOccurrenceID uuid.UUID          `json:"event_occurrence_id" doc:"ID of the event occurrence to register for" format:"uuid" required:"true"`
		Status            RegistrationStatus `json:"status" doc:"Initial status of the registration" default:"registered" enum:"registered,cancelled"`
	} `json:"body"`
}

type CreateRegistrationOutput struct {
	Body Registration `json:"body" doc:"The newly created registration with full details"`
}

type UpdateRegistrationInput struct {
	ID   uuid.UUID `path:"id" format:"uuid" doc:"Registration ID to update" required:"true"`
	Body struct {
		ChildID           *uuid.UUID          `json:"child_id,omitempty" doc:"Updated child ID (optional)" format:"uuid"`
		GuardianID        *uuid.UUID          `json:"guardian_id,omitempty" doc:"Updated guardian ID (optional)" format:"uuid"`
		EventOccurrenceID *uuid.UUID          `json:"event_occurrence_id,omitempty" doc:"Updated event occurrence ID (optional)" format:"uuid"`
		Status            *RegistrationStatus `json:"status,omitempty" doc:"Updated registration status (optional)" enum:"registered,cancelled"`
	} `json:"body"`
}

type UpdateRegistrationOutput struct {
	Body Registration `json:"body" doc:"The updated registration with full details"`
}

type GetRegistrationByIDInput struct {
	ID uuid.UUID `path:"id" format:"uuid" doc:"Registration ID to retrieve" required:"true"`
}

type GetRegistrationByIDOutput struct {
	Body Registration `json:"body" doc:"The requested registration with full details"`
}

type GetRegistrationsByChildIDInput struct {
	ChildID uuid.UUID `path:"child_id" format:"uuid" doc:"Child ID to retrieve registrations for" required:"true"`
}

type GetRegistrationsByChildIDOutput struct {
	Body struct {
		Registrations []Registration `json:"registrations" doc:"List of registrations for the child"`
	} `json:"body"`
}

type GetRegistrationsByGuardianIDInput struct {
	GuardianID uuid.UUID `path:"guardian_id" format:"uuid" doc:"Guardian ID to retrieve registrations for" required:"true"`
}

type GetRegistrationsByGuardianIDOutput struct {
	Body struct {
		Registrations []Registration `json:"registrations" doc:"List of registrations for the guardian"`
	} `json:"body"`
}

type DeleteRegistrationInput struct {
	ID uuid.UUID `path:"id" format:"uuid" doc:"Registration ID to delete" required:"true"`
}

type DeleteRegistrationOutput struct {
	Body Registration `json:"body" doc:"The deleted registration details"`
}
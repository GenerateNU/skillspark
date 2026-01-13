package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

// EventType represents the type of event
type EventType string

const (
	EventTypeTutor    EventType = "tutor"
	EventTypeActivity EventType = "activity"
)

type Event struct {
	ID             uuid.UUID      `json:"id" db:"id"`
	Title          string         `json:"title" db:"title"`
	Description    string         `json:"description" db:"description"`
	LocationID     uuid.UUID      `json:"location_id" db:"location_id"`
	OrganizationID uuid.UUID      `json:"organization_id" db:"organization_id"`
	MaxAttendees   int            `json:"max_attendees" db:"max_attendees"`
	CurrEnrolled   int            `json:"curr_enrolled" db:"curr_enrolled"`
	Type           EventType      `json:"type" db:"type"`
	Tags           pq.StringArray `json:"tags" db:"tags"`
	CreatedAt      time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at" db:"updated_at"`
}

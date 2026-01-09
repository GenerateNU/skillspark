package models

import (
	"time"

	"github.com/google/uuid"
)

type EventOccurrence struct {
	ID         uuid.UUID  `json:"id" db:"id"`
	ProviderID *uuid.UUID `json:"provider_id" db:"provider_id"`
	EventID    uuid.UUID  `json:"event_id" db:"event_id"`
	StartTime  time.Time  `json:"start_time" db:"start_time"`
	EndTime    time.Time  `json:"end_time" db:"end_time"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`
}
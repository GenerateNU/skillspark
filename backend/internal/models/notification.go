package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type NotificationType string

const (
	NotificationTypeSMS   NotificationType = "sms"
	NotificationTypeEmail NotificationType = "email"
)

type Notification struct {
	ID             uuid.UUID        `json:"id" db:"id"`
	RegistrationID uuid.UUID        `json:"registration_id" db:"registration_id"`
	Type           NotificationType `json:"type" db:"type"`
	Payload        json.RawMessage  `json:"payload" db:"payload"`
	CreatedAt      time.Time        `json:"created_at" db:"created_at"`
}

type NotificationPayload struct {
	Message string `json:"message"`
	To      string `json:"to"`                // Email or Phone
	Subject string `json:"subject,omitempty"` // For Email
}

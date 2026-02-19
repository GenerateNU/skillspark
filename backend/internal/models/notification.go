package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// NotificationType represents the type of notification
type NotificationType string

const (
	NotificationTypeEmail NotificationType = "email"
	NotificationTypePush  NotificationType = "push"
	NotificationTypeBoth  NotificationType = "both"
)

// NotificationStatus represents the status of a notification
type NotificationStatus string

const (
	NotificationStatusPending NotificationStatus = "pending"
	NotificationStatusSent    NotificationStatus = "sent"
	NotificationStatusFailed  NotificationStatus = "failed"
)

// Notification represents a scheduled notification in the database
type Notification struct {
	ID                 uuid.UUID        `json:"id" db:"id"`
	NotificationType   NotificationType  `json:"notification_type" db:"notification_type"`
	RecipientEmail     *string          `json:"recipient_email,omitempty" db:"recipient_email"`
	RecipientPushToken *string          `json:"recipient_push_token,omitempty" db:"recipient_push_token"`
	Subject            *string          `json:"subject,omitempty" db:"subject"`
	Body               string           `json:"body" db:"body"`
	Metadata           json.RawMessage  `json:"metadata,omitempty" db:"metadata"`
	ScheduledFor       time.Time        `json:"scheduled_for" db:"scheduled_for"`
	SentAt             *time.Time       `json:"sent_at,omitempty" db:"sent_at"`
	Status             NotificationStatus `json:"status" db:"status"`
	CreatedAt          time.Time        `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time        `json:"updated_at" db:"updated_at"`
}

// NotificationMessage represents the payload structure sent to SQS
// This is what the Lambda function will receive
type NotificationMessage struct {
	NotificationType   NotificationType `json:"notification_type"`
	RecipientEmail     *string          `json:"recipient_email,omitempty"`
	RecipientPushToken *string          `json:"recipient_push_token,omitempty"`
	Subject            *string          `json:"subject,omitempty"`
	Body               string           `json:"body"`
	Metadata           json.RawMessage `json:"metadata,omitempty"`
}

// CreateScheduledNotificationInput is used internally to create a scheduled notification
type CreateScheduledNotificationInput struct {
	NotificationType   NotificationType
	RecipientEmail     *string
	RecipientPushToken *string
	Subject            *string
	Body               string
	Metadata           json.RawMessage
	ScheduledFor       time.Time
}

// SendNotificationInput is used internally to send an immediate notification
type SendNotificationInput struct {
	NotificationType   NotificationType
	RecipientEmail     *string
	RecipientPushToken *string
	Subject            *string
	Body               string
	Metadata           json.RawMessage
}


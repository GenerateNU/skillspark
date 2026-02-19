package main

import "encoding/json"

// NotificationType represents the type of notification
type NotificationType string

const (
	NotificationTypeEmail NotificationType = "email"
	NotificationTypePush  NotificationType = "push"
	NotificationTypeBoth  NotificationType = "both"
)

// NotificationMessage represents the payload structure sent to SQS
// This matches the backend models.NotificationMessage structure
type NotificationMessage struct {
	NotificationType   NotificationType `json:"notification_type"`
	RecipientEmail     *string          `json:"recipient_email,omitempty"`
	RecipientPushToken *string          `json:"recipient_push_token,omitempty"`
	Subject            *string          `json:"subject,omitempty"`
	Body               string           `json:"body"`
	Metadata           json.RawMessage  `json:"metadata,omitempty"`
}


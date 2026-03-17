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


// SQSEvent represents an SQS event
type SQSEvent struct {
	Records []SQSEventRecord `json:"Records"`
}

// SQSEventRecord represents a single SQS record
type SQSEventRecord struct {
	MessageID     string `json:"messageId"`
	ReceiptHandle string `json:"receiptHandle"`
	Body          string `json:"body"`
	Attributes    map[string]interface{} `json:"attributes,omitempty"`
}

// SQSEventResponse represents the response from the Lambda handler
// Supports partial batch failures for SQS event source mapping
type SQSEventResponse struct {
	BatchItemFailures []BatchItemFailure `json:"batchItemFailures,omitempty"`
}

// BatchItemFailure represents a failed message in the batch
type BatchItemFailure struct {
	ItemIdentifier string `json:"itemIdentifier"`
}

// Handler processes SQS events
type Handler struct {
	processor *NotificationProcessor
}


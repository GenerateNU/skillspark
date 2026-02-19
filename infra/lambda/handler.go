package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
)

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

// NewHandler creates a new SQS event handler
func NewHandler(processor *NotificationProcessor) *Handler {
	return &Handler{
		processor: processor,
	}
}

// Handle processes an SQS event and returns a response with any batch item failures
func (h *Handler) Handle(ctx context.Context, event SQSEvent) (SQSEventResponse, error) {
	var failures []BatchItemFailure

	slog.Info("Processing SQS event",
		"record_count", len(event.Records),
	)

	// Process each record in the batch
	for _, record := range event.Records {
		if err := h.processRecord(ctx, record); err != nil {
			slog.Error("Failed to process record",
				"message_id", record.MessageID,
				"error", err,
			)

			// Add to batch item failures for partial batch failure support
			failures = append(failures, BatchItemFailure{
				ItemIdentifier: record.MessageID,
			})
		}
	}

	// Return response with any failures
	// If there are failures, SQS will retry only those messages
	// If empty, all messages will be deleted from the queue
	response := SQSEventResponse{
		BatchItemFailures: failures,
	}

	if len(failures) > 0 {
		slog.Warn("Some messages failed processing",
			"failed_count", len(failures),
			"total_count", len(event.Records),
		)
	} else {
		slog.Info("All messages processed successfully",
			"total_count", len(event.Records),
		)
	}

	return response, nil
}

// processRecord processes a single SQS record
func (h *Handler) processRecord(ctx context.Context, record SQSEventRecord) error {
	// Parse the message body (JSON string)
	var message NotificationMessage
	if err := json.Unmarshal([]byte(record.Body), &message); err != nil {
		return fmt.Errorf("failed to unmarshal message body: %w", err)
	}

	// Process the notification
	if err := h.processor.ProcessNotification(ctx, message); err != nil {
		return fmt.Errorf("failed to process notification: %w", err)
	}

	return nil
}


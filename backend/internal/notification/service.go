package notification

import (
	"context"
	"fmt"
	"skillspark/internal/models"
	"skillspark/internal/sqs_client"
	"skillspark/internal/storage"
)

type Service struct {
	repo     *storage.Repository
	sqsClient sqs_client.SQSInterface
}

func NewService(repo *storage.Repository, sqsClient sqs_client.SQSInterface) *Service {
	return &Service{
		repo:     repo,
		sqsClient: sqsClient,
	}
}

// SendNotification sends an immediate notification to SQS
func (s *Service) SendNotification(ctx context.Context, input *models.SendNotificationInput) error {
	// Validate input
	if err := validateNotificationInput(input.NotificationType, input.RecipientEmail, input.RecipientPushToken); err != nil {
		return err
	}

	// Create notification message for SQS
	message := models.NotificationMessage{
		NotificationType:   input.NotificationType,
		RecipientEmail:     input.RecipientEmail,
		RecipientPushToken: input.RecipientPushToken,
		Subject:            input.Subject,
		Body:               input.Body,
		Metadata:           input.Metadata,
	}

	// Send to SQS
	if err := s.sqsClient.SendMessage(ctx, message); err != nil {
		return fmt.Errorf("failed to send notification to SQS: %w", err)
	}

	return nil
}

// ScheduleNotification schedules a notification to be sent at a later time
func (s *Service) ScheduleNotification(ctx context.Context, input *models.CreateScheduledNotificationInput) (*models.Notification, error) {
	// Validate input
	if err := validateNotificationInput(input.NotificationType, input.RecipientEmail, input.RecipientPushToken); err != nil {
		return nil, err
	}

	// Create scheduled notification in database
	notification, err := s.repo.Notification.CreateScheduledNotification(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to schedule notification: %w", err)
	}

	return notification, nil
}

// validateNotificationInput validates that the notification has the required recipients
func validateNotificationInput(notificationType models.NotificationType, email *string, pushToken *string) error {
	switch notificationType {
	case models.NotificationTypeEmail:
		if email == nil || *email == "" {
			return fmt.Errorf("recipient_email is required for email notifications")
		}
	case models.NotificationTypePush:
		if pushToken == nil || *pushToken == "" {
			return fmt.Errorf("recipient_push_token is required for push notifications")
		}
	case models.NotificationTypeBoth:
		if email == nil || *email == "" {
			return fmt.Errorf("recipient_email is required for email notifications")
		}
		if pushToken == nil || *pushToken == "" {
			return fmt.Errorf("recipient_push_token is required for push notifications")
		}
	default:
		return fmt.Errorf("invalid notification_type: %s", notificationType)
	}
	return nil
}


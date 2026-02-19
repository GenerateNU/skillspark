package notification

import (
	"context"
	"fmt"
	"log/slog"
	"skillspark/internal/models"
	"skillspark/internal/sqs_client"
	"skillspark/internal/storage"
	"time"
)

type Scheduler struct {
	repo      *storage.Repository
	sqsClient sqs_client.SQSInterface
	interval  time.Duration
}

func NewScheduler(repo *storage.Repository, sqsClient sqs_client.SQSInterface) *Scheduler {
	return &Scheduler{
		repo:      repo,
		sqsClient: sqsClient,
		interval:  5 * time.Minute,
	}
}

// StartScheduler starts the background goroutine that checks for scheduled notifications
func (s *Scheduler) StartScheduler(ctx context.Context) {
	ticker := time.NewTicker(s.interval)
	defer ticker.Stop()

	// Run immediately on start
	s.checkScheduledNotifications(ctx)

	// Then run on interval
	for {
		select {
		case <-ctx.Done():
			slog.Info("Scheduler stopped")
			return
		case <-ticker.C:
			s.checkScheduledNotifications(ctx)
		}
	}
}

// checkScheduledNotifications queries the database for due notifications and forwards them to SQS
func (s *Scheduler) checkScheduledNotifications(ctx context.Context) {
	slog.Info("Checking for scheduled notifications")

	// Get pending notifications that are due
	notifications, err := s.repo.Notification.GetPendingNotifications(ctx)
	if err != nil {
		slog.Error("Failed to get pending notifications", "error", err)
		return
	}

	if len(notifications) == 0 {
		slog.Info("No pending notifications found")
		return
	}
	

	slog.Info("Found pending notifications", "count", len(notifications))

	// Process each notification
	for _, notification := range notifications {
		if err := s.processNotification(ctx, notification); err != nil {
			slog.Error("Failed to process notification", "id", notification.ID, "error", err)
			// Update status to failed
			_, updateErr := s.repo.Notification.UpdateNotificationStatus(ctx, notification.ID, models.NotificationStatusFailed)
			if updateErr != nil {
				slog.Error("Failed to update notification status to failed", "id", notification.ID, "error", updateErr)
			}
			continue
		}

		// Update status to sent (even though it's just queued, we mark it as sent to prevent reprocessing)
		// The actual sending will be handled by Lambda
		_, err := s.repo.Notification.UpdateNotificationStatus(ctx, notification.ID, models.NotificationStatusSent)
		if err != nil {
			slog.Error("Failed to update notification status", "id", notification.ID, "error", err)
			// Continue processing other notifications even if status update fails
		} else {
			slog.Info("Notification forwarded to SQS", "id", notification.ID)
		}
	}
}

// processNotification forwards a notification to SQS
func (s *Scheduler) processNotification(ctx context.Context, notification models.Notification) error {
	// Create notification message for SQS
	message := models.NotificationMessage{
		NotificationType:   notification.NotificationType,
		RecipientEmail:     notification.RecipientEmail,
		RecipientPushToken: notification.RecipientPushToken,
		Subject:            notification.Subject,
		Body:               notification.Body,
		Metadata:           notification.Metadata,
	}

	// Send to SQS
	if err := s.sqsClient.SendMessage(ctx, message); err != nil {
		return fmt.Errorf("failed to send notification to SQS: %w", err)
	}

	return nil
}


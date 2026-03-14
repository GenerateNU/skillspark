package jobs

import (
	"context"
	"fmt"
	"log/slog"
	"skillspark/internal/models"
	"log"
)

func (j *JobScheduler) SendScheduledNotificationsJob() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("CapturePaymentsJob panicked: %v", r)
		}
	}()
	
	ctx := context.Background()

	// Get pending notifications that are due
	notifications, err := j.repo.Notification.GetPendingNotifications(ctx)
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
		if err := j.processNotification(ctx, notification); err != nil {
			slog.Error("Failed to process notification", "id", notification.ID, "error", err)
			// Update status to failed
			_, updateErr := j.repo.Notification.UpdateNotificationStatus(ctx, notification.ID, models.NotificationStatusFailed)
			if updateErr != nil {
				slog.Error("Failed to update notification status to failed", "id", notification.ID, "error", updateErr)
			}
			continue
		}

		// Update status to sent (even though it's just queued, we mark it as sent to prevent reprocessing)
		// The actual sending will be handled by Lambda
		_, err := j.repo.Notification.UpdateNotificationStatus(ctx, notification.ID, models.NotificationStatusSent)
		if err != nil {
			slog.Error("Failed to update notification status", "id", notification.ID, "error", err)
			// Continue processing other notifications even if status update fails
		} else {
			slog.Info("Notification forwarded to SQS", "id", notification.ID)
		}
	}
}


func (j *JobScheduler) processNotification(ctx context.Context, notification models.Notification) error {
	// Create notification message for SQS
	message := &models.SendNotificationInput{
		NotificationType:   notification.NotificationType,
		RecipientEmail:     notification.RecipientEmail,
		RecipientPushToken: notification.RecipientPushToken,
		Subject:            notification.Subject,
		Body:               notification.Body,
		Metadata:           notification.Metadata,
	}

	// Send to SQS
	if err := j.notifService.SendNotification(ctx, message); err != nil {
		return fmt.Errorf("failed to send notification to SQS: %w", err)
	}

	return nil
}
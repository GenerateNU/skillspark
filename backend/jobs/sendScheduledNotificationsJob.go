package jobs

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	"log/slog"
	"skillspark/internal/models"
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

	// Collect unique guardian IDs from notifications that have one
	seen := make(map[uuid.UUID]bool)
	guardianIDs := make([]uuid.UUID, 0)
	for _, n := range notifications {
		if n.GuardianID != nil && !seen[*n.GuardianID] {
			guardianIDs = append(guardianIDs, *n.GuardianID)
			seen[*n.GuardianID] = true
		}
	}

	// Bulk-fetch notification preferences for all relevant guardians
	var guardianPrefs map[uuid.UUID]models.GuardianNotificationPreferences
	if len(guardianIDs) > 0 {
		guardianPrefs, err = j.repo.Guardian.GetGuardianNotificationPreferences(ctx, guardianIDs)
		if err != nil {
			slog.Error("Failed to get guardian notification preferences", "error", err)
			return
		}
	}

	// Process each notification
	for _, notification := range notifications {
		// Check guardian's notification preferences before sending
		if notification.GuardianID != nil {
			if prefs, ok := guardianPrefs[*notification.GuardianID]; ok {
				if notification.NotificationType == models.NotificationTypeEmail && !prefs.EmailNotifications {
					slog.Info("Skipping notification: guardian has email notifications disabled", "id", notification.ID, "guardian_id", *notification.GuardianID)
					_, updateErr := j.repo.Notification.UpdateNotificationStatus(ctx, notification.ID, models.NotificationStatusSent)
					if updateErr != nil {
						slog.Error("Failed to update skipped notification status", "id", notification.ID, "error", updateErr)
					}
					continue
				}
				if notification.NotificationType == models.NotificationTypePush && !prefs.PushNotifications {
					slog.Info("Skipping notification: guardian has push notifications disabled", "id", notification.ID, "guardian_id", *notification.GuardianID)
					_, updateErr := j.repo.Notification.UpdateNotificationStatus(ctx, notification.ID, models.NotificationStatusSent)
					if updateErr != nil {
						slog.Error("Failed to update skipped notification status", "id", notification.ID, "error", updateErr)
					}
					continue
				}
			}
		}

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

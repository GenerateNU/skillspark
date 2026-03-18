package notification

import (
	"context"
	"skillspark/internal/models"
)

type NotificationServiceInterface interface {
	SendNotification(ctx context.Context, input *models.SendNotificationInput) error
	ScheduleNotification(ctx context.Context, input *models.CreateScheduledNotificationInput) (*models.Notification, error)
}

package notification

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *NotificationRepository) CreateScheduledNotification(ctx context.Context, input *models.CreateScheduledNotificationInput) (*models.Notification, error) {
	query, err := schema.ReadSQLBaseScript("create.sql", SqlNotificationFiles)
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	row := r.db.QueryRow(ctx, query,
		input.NotificationType,
		input.RecipientEmail,
		input.RecipientPushToken,
		input.Subject,
		input.Body,
		input.Metadata,
		input.ScheduledFor,
		models.NotificationStatusPending,
	)

	var notification models.Notification

	err = row.Scan(
		&notification.ID,
		&notification.NotificationType,
		&notification.RecipientEmail,
		&notification.RecipientPushToken,
		&notification.Subject,
		&notification.Body,
		&notification.Metadata,
		&notification.ScheduledFor,
		&notification.SentAt,
		&notification.Status,
		&notification.CreatedAt,
		&notification.UpdatedAt,
	)

	if err != nil {
		errr := errs.InternalServerError("Failed to create scheduled notification: ", err.Error())
		return nil, &errr
	}

	return &notification, nil
}


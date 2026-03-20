package notification

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
)

func (r *NotificationRepository) UpdateNotificationStatus(ctx context.Context, id uuid.UUID, status models.NotificationStatus) (*models.Notification, error) {
	query, err := schema.ReadSQLBaseScript("update_status.sql", SqlNotificationFiles)
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	row := r.db.QueryRow(ctx, query, id, status)

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
		errr := errs.InternalServerError("Failed to update notification status: ", err.Error())
		return nil, &errr
	}

	return &notification, nil
}


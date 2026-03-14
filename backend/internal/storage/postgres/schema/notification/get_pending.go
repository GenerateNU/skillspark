package notification

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *NotificationRepository) GetPendingNotifications(ctx context.Context) ([]models.Notification, error) {
	query, err := schema.ReadSQLBaseScript("get_pending.sql", SqlNotificationFiles)
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		errr := errs.InternalServerError("Failed to query pending notifications: ", err.Error())
		return nil, &errr
	}
	defer rows.Close()

	var notifications []models.Notification

	for rows.Next() {
		var notification models.Notification
		err := rows.Scan(
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
			errr := errs.InternalServerError("Failed to scan notification: ", err.Error())
			return nil, &errr
		}
		notifications = append(notifications, notification)
	}

	if err := rows.Err(); err != nil {
		errr := errs.InternalServerError("Error iterating notifications: ", err.Error())
		return nil, &errr
	}

	return notifications, nil
}


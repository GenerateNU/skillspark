package notification

import (
	"context"
	_ "embed"

	"skillspark/internal/models"

	"github.com/jackc/pgx/v5"
)

//go:embed sql/create_notification.sql
var createNotificationSQL string

func (r *NotificationRepository) CreateNotification(ctx context.Context, tx pgx.Tx, notification *models.Notification) error {
	row := tx.QueryRow(ctx, createNotificationSQL, notification.RegistrationID, notification.Type, notification.Payload)
	return row.Scan(&notification.ID, &notification.RegistrationID, &notification.Type, &notification.Payload, &notification.CreatedAt)
}

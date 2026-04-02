package notification

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/jackc/pgx/v5"
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

	notifications, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Notification])
	if err != nil {
		errr := errs.InternalServerError("Failed to scan notification: ", err.Error())
		return nil, &errr
	}

	return notifications, nil
}

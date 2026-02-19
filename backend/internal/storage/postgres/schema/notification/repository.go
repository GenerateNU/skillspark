package notification

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type NotificationRepository struct {
	Db *pgxpool.Pool
}

func NewNotificationRepository(db *pgxpool.Pool) *NotificationRepository {
	return &NotificationRepository{Db: db}
}

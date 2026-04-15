package guardian

import (
	"context"

	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
)

func (r *GuardianRepository) GetGuardianNotificationPreferences(ctx context.Context, ids []uuid.UUID) (map[uuid.UUID]models.GuardianNotificationPreferences, error) {
	query, err := schema.ReadSQLBaseScript("get_notification_preferences.sql", SqlGuardianFiles)
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	uuidStrings := make([]string, len(ids))
	for i, id := range ids {
		uuidStrings[i] = id.String()
	}

	rows, err := r.db.Query(ctx, query, uuidStrings)
	if err != nil {
		errr := errs.InternalServerError("Failed to query guardian notification preferences: ", err.Error())
		return nil, &errr
	}
	defer rows.Close()

	result := make(map[uuid.UUID]models.GuardianNotificationPreferences, len(ids))
	for rows.Next() {
		var id uuid.UUID
		var prefs models.GuardianNotificationPreferences
		if err := rows.Scan(&id, &prefs.PushNotifications, &prefs.EmailNotifications); err != nil {
			errr := errs.InternalServerError("Failed to scan guardian notification preferences: ", err.Error())
			return nil, &errr
		}
		result[id] = prefs
	}

	if err := rows.Err(); err != nil {
		errr := errs.InternalServerError("Failed to iterate guardian notification preferences: ", err.Error())
		return nil, &errr
	}

	return result, nil
}

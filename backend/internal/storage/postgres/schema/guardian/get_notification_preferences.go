package guardian

import (
	"context"

	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *GuardianRepository) GetGuardianNotificationPreferences(
	ctx context.Context,
	ids []uuid.UUID,
) (map[uuid.UUID]models.GuardianNotificationPreferences, error) {

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

	type row struct {
		ID uuid.UUID `db:"id"`
		models.GuardianNotificationPreferences
	}

	collected, err := pgx.CollectRows(rows, pgx.RowToStructByName[row])
	if err != nil {
		errr := errs.InternalServerError("Failed to collect guardian notification preferences: ", err.Error())
		return nil, &errr
	}

	result := make(map[uuid.UUID]models.GuardianNotificationPreferences, len(collected))
	for _, r := range collected {
		result[r.ID] = r.GuardianNotificationPreferences
	}

	return result, nil
}

package saved

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
	"skillspark/internal/utils"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *SavedRepository) GetByGuardianID(ctx context.Context, user_id uuid.UUID, pagination utils.Pagination) ([]models.Saved, error) {

	query, err := schema.ReadSQLBaseScript("get_all_for_user.sql", SqlSavedFiles)
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, err
	}

	rows, err := r.db.Query(
		ctx,
		query,
		user_id,
		pagination.Limit,
		pagination.GetOffset(),
	)

	if err != nil {
		err := errs.InternalServerError("Failed to fetch all saved: ", err.Error())
		return nil, &err
	}
	defer rows.Close()

	saved, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Saved])
	if err != nil {
		err := errs.InternalServerError("Failed to scan all event occurrences: ", err.Error())
		return nil, &err
	}
	return saved, nil
}

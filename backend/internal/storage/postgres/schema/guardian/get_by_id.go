package guardian

import (
	"context"

	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
)

func (r *GuardianRepository) GetGuardianByID(ctx context.Context, id uuid.UUID) (*models.Guardian, *errs.HTTPError) {
	query, err := schema.ReadSQLBaseScript("guardian/sql/get_by_id.sql")
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	row := r.db.QueryRow(ctx, query, id)

	var guardian models.Guardian

	err = row.Scan(&guardian.ID, &guardian.UserID, &guardian.CreatedAt, &guardian.UpdatedAt)
	if err != nil {
		err := errs.InternalServerError("Failed to get guardian by id: ", err.Error())
		return nil, &err
	}

	return &guardian, nil
}

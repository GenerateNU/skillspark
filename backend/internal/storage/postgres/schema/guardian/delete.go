package guardian

import (
	"context"

	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
)

func (r *GuardianRepository) DeleteGuardian(ctx context.Context, id uuid.UUID) (*models.Guardian, *errs.HTTPError) {
	query, err := schema.ReadSQLBaseScript("guardian/sql/delete.sql")
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	row := r.db.QueryRow(ctx, query, id)

	var deletedGuardian models.Guardian

	err = row.Scan(&deletedGuardian.ID, &deletedGuardian.UserID, &deletedGuardian.CreatedAt, &deletedGuardian.UpdatedAt)
	if err != nil {
		err := errs.InternalServerError("Failed to delete guardian: ", err.Error())
		return nil, &err
	}

	return &deletedGuardian, nil
}

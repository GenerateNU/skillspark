package guardian

import (
	"context"

	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *GuardianRepository) UpdateGuardian(ctx context.Context, guardian *models.UpdateGuardianInput) (*models.Guardian, *errs.HTTPError) {
	query, err := schema.ReadSQLBaseScript("guardian/sql/update.sql")
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	row := r.db.QueryRow(ctx, query, guardian.ID, guardian.Body.UserID)

	var updatedGuardian models.Guardian

	err = row.Scan(&updatedGuardian.ID, &updatedGuardian.UserID, &updatedGuardian.CreatedAt, &updatedGuardian.UpdatedAt)
	if err != nil {
		err := errs.InternalServerError("Failed to update guardian: ", err.Error())
		return nil, &err
	}

	return &updatedGuardian, nil
}

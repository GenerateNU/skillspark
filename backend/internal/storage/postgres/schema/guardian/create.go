package guardian

import (
	"context"

	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *GuardianRepository) CreateGuardian(ctx context.Context, guardian *models.CreateGuardianInput) (*models.Guardian, *errs.HTTPError) {
	query, err := schema.ReadSQLBaseScript("guardian/sql/create.sql")
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	row := r.db.QueryRow(ctx, query, guardian.Body.UserID)

	var createdGuardian models.Guardian

	err = row.Scan(&createdGuardian.ID, &createdGuardian.UserID, &createdGuardian.CreatedAt, &createdGuardian.UpdatedAt)
	if err != nil {
		err := errs.InternalServerError("Failed to create guardian: ", err.Error())
		return nil, &err
	}

	return &createdGuardian, nil
}

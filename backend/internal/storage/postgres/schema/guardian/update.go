package guardian

import (
	"context"

	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *GuardianRepository) UpdateGuardian(ctx context.Context, guardian *models.UpdateGuardianInput) (*models.Guardian, error) {
	query, err := schema.ReadSQLBaseScript("guardian/sql/update.sql")
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	row := r.db.QueryRow(ctx, query, guardian.ID, guardian.Body.Name, guardian.Body.Email, guardian.Body.Username, guardian.Body.ProfilePictureS3Key, guardian.Body.LanguagePreference)

	var updatedGuardian models.Guardian

	err = row.Scan(&updatedGuardian.ID, &updatedGuardian.UserID, &updatedGuardian.Name, &updatedGuardian.Email, &updatedGuardian.Username, &updatedGuardian.ProfilePictureS3Key, &updatedGuardian.LanguagePreference, &updatedGuardian.CreatedAt, &updatedGuardian.UpdatedAt)
	if err != nil {
		err := errs.InternalServerError("Failed to update guardian: ", err.Error())
		return nil, &err
	}

	return &updatedGuardian, nil
}

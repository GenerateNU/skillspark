package guardian

import (
	"context"

	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *GuardianRepository) UpdateGuardian(ctx context.Context, guardian *models.UpdateGuardianInput) (*models.Guardian, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	var updatedGuardian models.Guardian
	updatedGuardian.ID = guardian.ID

	guardianQuery, err := schema.ReadSQLBaseScript("guardian/sql/update_guardian.sql")
	if err != nil {
		_ = tx.Rollback(ctx)
		err := errs.InternalServerError("Failed to read guardian update query: ", err.Error())
		return nil, &err
	}

	err = tx.QueryRow(ctx, guardianQuery, guardian.ID).Scan(
		&updatedGuardian.UserID,
		&updatedGuardian.StripeCustomerID,
		&updatedGuardian.CreatedAt,
		&updatedGuardian.UpdatedAt,
	)

	if err != nil {
		_ = tx.Rollback(ctx)
		err := errs.InternalServerError("Failed to update guardian table: ", err.Error())
		return nil, &err
	}

	userQuery, err := schema.ReadSQLBaseScript("guardian/sql/update_user.sql")
	if err != nil {
		_ = tx.Rollback(ctx)
		err := errs.InternalServerError("Failed to read user update query: ", err.Error())
		return nil, &err
	}

	err = tx.QueryRow(ctx, userQuery,
		updatedGuardian.UserID,
		guardian.Body.Name,
		guardian.Body.Email,
		guardian.Body.Username,
		guardian.Body.ProfilePictureS3Key,
		guardian.Body.LanguagePreference,
	).Scan(
		&updatedGuardian.Name,
		&updatedGuardian.Email,
		&updatedGuardian.Username,
		&updatedGuardian.ProfilePictureS3Key,
		&updatedGuardian.LanguagePreference,
	)

	if err != nil {
		_ = tx.Rollback(ctx)
		err := errs.InternalServerError("Failed to update user table: ", err.Error())
		return nil, &err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &updatedGuardian, nil
}
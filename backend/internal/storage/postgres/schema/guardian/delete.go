package guardian

import (
	"context"
	"errors"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *GuardianRepository) DeleteGuardian(ctx context.Context, id uuid.UUID) (*models.Guardian, error) {
	query, err := schema.ReadSQLBaseScript("guardian/sql/delete.sql")
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	row := r.db.QueryRow(ctx, query, id)

	var deletedGuardian models.Guardian

	err = row.Scan(&deletedGuardian.ID, &deletedGuardian.UserID, &deletedGuardian.Name, &deletedGuardian.Email, &deletedGuardian.Username, &deletedGuardian.ProfilePictureS3Key, &deletedGuardian.LanguagePreference, &deletedGuardian.AuthID, &deletedGuardian.CreatedAt, &deletedGuardian.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err := errs.NotFound("Guardian", "id", id)
			return nil, &err
		}
		err := errs.InternalServerError("Failed to delete guardian: ", err.Error())
		return nil, &err
	}

	return &deletedGuardian, nil
}

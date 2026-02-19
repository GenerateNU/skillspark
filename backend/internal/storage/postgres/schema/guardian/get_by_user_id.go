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

func (r *GuardianRepository) GetGuardianByUserID(ctx context.Context, id uuid.UUID) (*models.Guardian, error) {
	query, err := schema.ReadSQLBaseScript("get_by_user_id.sql", SqlGuardianFiles)
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	row := r.db.QueryRow(ctx, query, id)

	var guardian models.Guardian

	err = row.Scan(&guardian.ID, &guardian.UserID, &guardian.Name, &guardian.Email, &guardian.Username, &guardian.ProfilePictureS3Key, &guardian.LanguagePreference, &guardian.CreatedAt, &guardian.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err := errs.NotFound("Guardian", "user_id", id)
			return nil, &err
		}
		err := errs.InternalServerError("Failed to get guardian by user id: ", err.Error())
		return nil, &err
	}

	return &guardian, nil
}

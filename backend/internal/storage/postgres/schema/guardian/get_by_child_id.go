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

func (r *GuardianRepository) GetGuardianByChildID(ctx context.Context, childID uuid.UUID) (*models.Guardian, error) {
	query, err := schema.ReadSQLBaseScript("get_by_child_id.sql", SqlGuardianFiles)
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	row := r.db.QueryRow(ctx, query, childID)

	var guardian models.Guardian

	err = row.Scan(&guardian.ID, &guardian.UserID, &guardian.Name, &guardian.Email, &guardian.Username, &guardian.ProfilePictureS3Key, &guardian.LanguagePreference, &guardian.AuthID, &guardian.CreatedAt, &guardian.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err := errs.BadRequest("Child with id: " + childID.String() + " not found")
			return nil, &err
		}
		err := errs.InternalServerError("Failed to get guardian by child id: ", err.Error())
		return nil, &err
	}

	return &guardian, nil
}

// TODO: do repo testing and then move onto actual endpoints. Verify types in

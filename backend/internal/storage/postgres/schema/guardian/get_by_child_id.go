package guardian

import (
	"context"

	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
)

func (r *GuardianRepository) GetGuardianByChildID(ctx context.Context, childID uuid.UUID) (*models.Guardian, *errs.HTTPError) {
	query, err := schema.ReadSQLBaseScript("guardian/sql/get_by_child_id.sql")
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	row := r.db.QueryRow(ctx, query, childID)

	var guardian models.Guardian

	err = row.Scan(&guardian.ID, &guardian.UserID, &guardian.CreatedAt, &guardian.UpdatedAt)
	if err != nil {
		err := errs.InternalServerError("Failed to get guardian by child id: ", err.Error())
		return nil, &err
	}

	return &guardian, nil
}

// TODO: do repo testing and then move onto actual endpoints. Verify types in

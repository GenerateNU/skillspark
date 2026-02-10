package guardian

import (
	"context"
	"errors"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/jackc/pgx/v5"
)

func (r *GuardianRepository) GetGuardianByAuthID(ctx context.Context, authID string) (*models.Guardian, error) {
	query, err := schema.ReadSQLBaseScript("guardian/sql/get_by_auth_id.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	row := r.db.QueryRow(ctx, query, authID)
	var guardian models.Guardian
	err = row.Scan(
		&guardian.ID,
		&guardian.UserID,
		&guardian.StripeCustomerID,
		&guardian.CreatedAt,
		&guardian.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err := errs.NotFound("Guardian", "auth_id", authID)
			return nil, &err
		}
		err := errs.InternalServerError("Failed to fetch guardian by auth_id: ", err.Error())
		return nil, &err
	}

	return &guardian, nil
}
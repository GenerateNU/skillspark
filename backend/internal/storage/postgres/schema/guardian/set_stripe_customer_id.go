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

func (r *GuardianRepository) SetStripeCustomerID(
	ctx context.Context,
	guardianID uuid.UUID,
	stripeCustomerID string,
) (*models.Guardian, error) {
	query, err := schema.ReadSQLBaseScript("set_stripe_customer_id.sql", SqlGuardianFiles)
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	row := r.db.QueryRow(ctx, query, guardianID, stripeCustomerID)

	var updatedGuardian models.Guardian

	err = row.Scan(
		&updatedGuardian.ID,
		&updatedGuardian.UserID,
		&updatedGuardian.StripeCustomerID,
		&updatedGuardian.CreatedAt,
		&updatedGuardian.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			errr := errs.NotFound("Guardian", "id", guardianID)
			return nil, &errr
		}
		errr := errs.InternalServerError("Failed to set stripe customer ID: ", err.Error())
		return nil, &errr
	}

	return &updatedGuardian, nil
}
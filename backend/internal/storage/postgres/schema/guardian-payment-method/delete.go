package guardianpaymentmethod

import (
	"context"
	"errors"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *GuardianPaymentMethodRepository) DeleteGuardianPaymentMethod(
	ctx context.Context,
	id uuid.UUID,
) (*models.GuardianPaymentMethod, error) {
	query, err := schema.ReadSQLBaseScript("guardian-payment-method/sql/delete.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	row := r.db.QueryRow(ctx, query, id)

	var deletedPaymentMethod models.GuardianPaymentMethod

	err = row.Scan(
		&deletedPaymentMethod.ID,
		&deletedPaymentMethod.GuardianID,
		&deletedPaymentMethod.StripePaymentMethodID,
		&deletedPaymentMethod.CardBrand,
		&deletedPaymentMethod.CardLast4,
		&deletedPaymentMethod.CardExpMonth,
		&deletedPaymentMethod.CardExpYear,
		&deletedPaymentMethod.IsDefault,
		&deletedPaymentMethod.CreatedAt,
		&deletedPaymentMethod.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			errr := errs.NotFound("Guardian Payment Method", "id", id)
			return nil, &errr
		}
		errr := errs.InternalServerError("Failed to delete payment method: ", err.Error())
		return nil, &errr
	}

	return &deletedPaymentMethod, nil
}
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

func (r *GuardianPaymentMethodRepository) UpdateGuardianPaymentMethod(
	ctx context.Context,
	id uuid.UUID,
	isDefault bool,
) (*models.GuardianPaymentMethod, error) {
	query, err := schema.ReadSQLBaseScript("guardian-payment-method/sql/update.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	row := r.db.QueryRow(ctx, query, id, isDefault)

	var updatedPaymentMethod models.GuardianPaymentMethod

	err = row.Scan(
		&updatedPaymentMethod.ID,
		&updatedPaymentMethod.GuardianID,
		&updatedPaymentMethod.StripePaymentMethodID,
		&updatedPaymentMethod.CardBrand,
		&updatedPaymentMethod.CardLast4,
		&updatedPaymentMethod.CardExpMonth,
		&updatedPaymentMethod.CardExpYear,
		&updatedPaymentMethod.IsDefault,
		&updatedPaymentMethod.CreatedAt,
		&updatedPaymentMethod.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			errr := errs.NotFound("Guardian Payment Method", "id", id)
			return nil, &errr
		}
		errr := errs.InternalServerError("Failed to update payment method: ", err.Error())
		return nil, &errr
	}

	return &updatedPaymentMethod, nil
}
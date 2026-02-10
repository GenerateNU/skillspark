package guardianpaymentmethod

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *GuardianPaymentMethodRepository) CreateGuardianPaymentMethod(
	ctx context.Context,
	input *models.CreateGuardianPaymentMethodInput,
) (*models.GuardianPaymentMethod, error) {
	query, err := schema.ReadSQLBaseScript("guardian-payment-method/sql/create.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	row := r.db.QueryRow(ctx, query,
		input.Body.GuardianID,
		input.Body.StripePaymentMethodID,
		input.Body.CardBrand,
		input.Body.CardLast4,
		input.Body.CardExpMonth,
		input.Body.CardExpYear,
		input.Body.IsDefault,
	)

	var createdPaymentMethod models.GuardianPaymentMethod

	err = row.Scan(
		&createdPaymentMethod.ID,
		&createdPaymentMethod.GuardianID,
		&createdPaymentMethod.StripePaymentMethodID,
		&createdPaymentMethod.CardBrand,
		&createdPaymentMethod.CardLast4,
		&createdPaymentMethod.CardExpMonth,
		&createdPaymentMethod.CardExpYear,
		&createdPaymentMethod.IsDefault,
		&createdPaymentMethod.CreatedAt,
		&createdPaymentMethod.UpdatedAt,
	)

	if err != nil {
		errr := errs.InternalServerError("Failed to create guardian payment method: ", err.Error())
		return nil, &errr
	}

	return &createdPaymentMethod, nil
}
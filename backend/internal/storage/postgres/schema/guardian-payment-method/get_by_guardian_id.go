package guardianpaymentmethod

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *GuardianPaymentMethodRepository) GetPaymentMethodsByGuardianID(
	ctx context.Context,
	guardianID uuid.UUID,
) ([]models.GuardianPaymentMethod, error) {
	query, err := schema.ReadSQLBaseScript("guardian-payment-method/sql/get_by_guardian_id.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	rows, err := r.db.Query(ctx, query, guardianID)
	if err != nil {
		errr := errs.InternalServerError("Failed to fetch payment methods: ", err.Error())
		return nil, &errr
	}
	defer rows.Close()

	paymentMethods, err := pgx.CollectRows(rows, scanGuardianPaymentMethod)
	if err != nil {
		errr := errs.InternalServerError("Failed to scan payment methods: ", err.Error())
		return nil, &errr
	}

	return paymentMethods, nil
}

func scanGuardianPaymentMethod(row pgx.CollectableRow) (models.GuardianPaymentMethod, error) {
	var pm models.GuardianPaymentMethod
	
	err := row.Scan(
		&pm.ID,
		&pm.GuardianID,
		&pm.StripePaymentMethodID,
		&pm.CardBrand,
		&pm.CardLast4,
		&pm.CardExpMonth,
		&pm.CardExpYear,
		&pm.IsDefault,
		&pm.CreatedAt,
		&pm.UpdatedAt,
	)
	
	return pm, err
}
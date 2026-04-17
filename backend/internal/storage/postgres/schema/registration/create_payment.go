package registration

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *RegistrationRepository) CreatePayment(ctx context.Context, input *models.CreatePaymentData) error {
	query, err := schema.ReadSQLBaseScript("create_payment.sql", SqlRegistrationFiles)
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return &errr
	}

	_, err = r.db.Exec(ctx, query,
		input.RegistrationID,
		input.StripePaymentIntentID,
		input.StripeCustomerID,
		input.OrgStripeAccountID,
		input.StripePaymentMethodID,
		input.TotalAmount,
		input.ProviderAmount,
		input.PlatformFeeAmount,
		input.Currency,
		input.PaymentIntentStatus,
	)
	if err != nil {
		errr := errs.InternalServerError("Failed to create payment record: ", err.Error())
		return &errr
	}

	return nil
}

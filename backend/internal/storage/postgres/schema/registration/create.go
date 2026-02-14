package registration

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *RegistrationRepository) CreateRegistration(ctx context.Context, input *models.CreateRegistrationWithPaymentData) (*models.CreateRegistrationOutput, error) {

	query, err := schema.ReadSQLBaseScript("registration/sql/create.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	row := r.db.QueryRow(ctx, query,
		input.ChildID,
		input.GuardianID,
		input.EventOccurrenceID,
		input.Status,
		input.StripePaymentIntentID,
		input.StripeCustomerID,
		input.OrgStripeAccountID,
		input.StripePaymentMethodID,
		input.TotalAmount,
		input.ProviderAmount, 
		input.PlatformFeeAmount, 
		input.Currency,
		input.PaymentIntentStatus)

	var createdRegistration models.CreateRegistrationOutput

	err = row.Scan(
		&createdRegistration.Body.ID,
		&createdRegistration.Body.ChildID,
		&createdRegistration.Body.GuardianID,
		&createdRegistration.Body.EventOccurrenceID,
		&createdRegistration.Body.Status,
		&createdRegistration.Body.CreatedAt,
		&createdRegistration.Body.UpdatedAt,
		&createdRegistration.Body.StripeCustomerID,
		&createdRegistration.Body.OrgStripeAccountID,
		&createdRegistration.Body.Currency,
		&createdRegistration.Body.PaymentIntentStatus,
		&createdRegistration.Body.CancelledAt,
		&createdRegistration.Body.StripePaymentIntentID,
		&createdRegistration.Body.TotalAmount,
		&createdRegistration.Body.ProviderAmount,
		&createdRegistration.Body.PlatformFeeAmount,
		&createdRegistration.Body.PaidAt,
		&createdRegistration.Body.StripePaymentMethodID,
		&createdRegistration.Body.EventName,
		&createdRegistration.Body.OccurrenceStartTime, 
	)

	if err != nil {
		errr := errs.InternalServerError("Failed to create registration: ", err.Error())
		return nil, &errr
	}

	return &createdRegistration, nil
}

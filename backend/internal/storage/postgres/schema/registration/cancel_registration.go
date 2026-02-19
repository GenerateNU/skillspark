package registration

import (
	"context"
	"errors"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/jackc/pgx/v5"
)

func (r *RegistrationRepository) CancelRegistration(ctx context.Context, input *models.CancelRegistrationInput) (*models.CancelRegistrationOutput, error) {
	query, err := schema.ReadSQLBaseScript("registration/sql/cancel_registration.sql", SqlRegistrationFiles)
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	row := r.db.QueryRow(ctx, query, input.ID)

	var output models.CancelRegistrationOutput

	err = row.Scan(
		&output.Body.Registration.ID,
		&output.Body.Registration.ChildID,
		&output.Body.Registration.GuardianID,
		&output.Body.Registration.EventOccurrenceID,
		&output.Body.Registration.Status,
		&output.Body.Registration.CreatedAt,
		&output.Body.Registration.UpdatedAt,
		&output.Body.Registration.StripeCustomerID,
		&output.Body.Registration.OrgStripeAccountID,
		&output.Body.Registration.Currency,
		&output.Body.Registration.PaymentIntentStatus,
		&output.Body.Registration.CancelledAt,
		&output.Body.Registration.StripePaymentIntentID,
		&output.Body.Registration.TotalAmount,
		&output.Body.Registration.ProviderAmount,
		&output.Body.Registration.PlatformFeeAmount,
		&output.Body.Registration.PaidAt,
		&output.Body.Registration.StripePaymentMethodID,
		&output.Body.Registration.EventName,
		&output.Body.Registration.OccurrenceStartTime,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			errr := errs.NotFound("Registration", "id", input.ID)
			return nil, &errr
		}
		errr := errs.InternalServerError("Failed to cancel registration: ", err.Error())
		return nil, &errr
	}

	return &output, nil
}
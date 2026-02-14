package registration

import (
	"context"
	"errors"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/jackc/pgx/v5"
)

func (r *RegistrationRepository) UpdateRegistration(ctx context.Context, input *models.UpdateRegistrationInput) (*models.UpdateRegistrationOutput, error) {
	query, err := schema.ReadSQLBaseScript("registration/sql/update.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	row := r.db.QueryRow(ctx, query,
		input.Body.ChildID,
		input.ID,
	)

	var updated models.UpdateRegistrationOutput

	err = row.Scan(
		&updated.Body.ID,
		&updated.Body.ChildID,
		&updated.Body.GuardianID,
		&updated.Body.EventOccurrenceID,
		&updated.Body.Status,
		&updated.Body.CreatedAt,
		&updated.Body.UpdatedAt,
		&updated.Body.StripeCustomerID,
		&updated.Body.OrgStripeAccountID,
		&updated.Body.Currency,
		&updated.Body.PaymentIntentStatus,
		&updated.Body.CancelledAt,
		&updated.Body.StripePaymentIntentID,
		&updated.Body.TotalAmount,
		&updated.Body.ProviderAmount,
		&updated.Body.PlatformFeeAmount,
		&updated.Body.PaidAt,
		&updated.Body.StripePaymentMethodID,
		&updated.Body.EventName,
		&updated.Body.OccurrenceStartTime,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			errr := errs.NotFound("Registration", "id", input.ID)
			return nil, &errr
		}
		errr := errs.InternalServerError("Failed to update registration: ", err.Error())
		return nil, &errr
	}

	return &updated, nil
}
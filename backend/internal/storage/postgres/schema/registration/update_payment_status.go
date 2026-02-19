package registration

import (
	"context"
	"errors"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/jackc/pgx/v5"
)

func (r *RegistrationRepository) UpdateRegistrationPaymentStatus(ctx context.Context, input *models.UpdateRegistrationPaymentStatusInput) (*models.UpdateRegistrationPaymentStatusOutput, error) {
	query, err := schema.ReadSQLBaseScript("registration/sql/update_payment_status.sql", SqlRegistrationFiles)
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	row := r.db.QueryRow(ctx, query, input.ID, input.Body.PaymentIntentStatus)

	var output models.UpdateRegistrationPaymentStatusOutput

	err = row.Scan(
		&output.Body.ID,
		&output.Body.ChildID,
		&output.Body.GuardianID,
		&output.Body.EventOccurrenceID,
		&output.Body.Status,
		&output.Body.CreatedAt,
		&output.Body.UpdatedAt,
		&output.Body.StripeCustomerID,
		&output.Body.OrgStripeAccountID,
		&output.Body.Currency,
		&output.Body.PaymentIntentStatus,
		&output.Body.CancelledAt,
		&output.Body.StripePaymentIntentID,
		&output.Body.TotalAmount,
		&output.Body.ProviderAmount,
		&output.Body.PlatformFeeAmount,
		&output.Body.PaidAt,
		&output.Body.StripePaymentMethodID,
		&output.Body.EventName,
		&output.Body.OccurrenceStartTime,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			errr := errs.NotFound("Registration", "id", input.ID)
			return nil, &errr
		}
		errr := errs.InternalServerError("Failed to update payment status: ", err.Error())
		return nil, &errr
	}

	return &output, nil
}
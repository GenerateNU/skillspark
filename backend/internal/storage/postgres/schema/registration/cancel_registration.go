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
	cancelQuery, err := schema.ReadSQLBaseScript("cancel_registration.sql", SqlRegistrationFiles)
	if err != nil {
		errr := errs.InternalServerError("Failed to read cancel query: ", err.Error())
		return nil, &errr
	}

	decrementQuery, err := schema.ReadSQLBaseScript("decrement_event_occurrence.sql", SqlRegistrationFiles)
	if err != nil {
		errr := errs.InternalServerError("Failed to read decrement query: ", err.Error())
		return nil, &errr
	}

	tx, err := r.db.Begin(ctx)
	if err != nil {
		errr := errs.InternalServerError("Failed to begin transaction: ", err.Error())
		return nil, &errr
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	row := tx.QueryRow(ctx, cancelQuery,
		input.ID,
		input.Status,
		input.PaymentIntentStatus,
	)

	var output models.CancelRegistrationOutput

	err = row.Scan(
		&output.Body.Registration.ID,
		&output.Body.Registration.ChildID,
		&output.Body.Registration.GuardianID,
		&output.Body.Registration.EventOccurrenceID,
		&output.Body.Registration.Status,
		&output.Body.Registration.StripePaymentIntentID,
		&output.Body.Registration.StripeCustomerID,
		&output.Body.Registration.OrgStripeAccountID,
		&output.Body.Registration.StripePaymentMethodID,
		&output.Body.Registration.TotalAmount,
		&output.Body.Registration.ProviderAmount,
		&output.Body.Registration.PlatformFeeAmount,
		&output.Body.Registration.Currency,
		&output.Body.Registration.PaymentIntentStatus,
		&output.Body.Registration.PaidAt,
		&output.Body.Registration.CancelledAt,
		&output.Body.Registration.CreatedAt,
		&output.Body.Registration.UpdatedAt,
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

	_, err = tx.Exec(ctx, decrementQuery, output.Body.Registration.EventOccurrenceID)
	if err != nil {
		errr := errs.InternalServerError("Failed to decrement enrolled: ", err.Error())
		return nil, &errr
	}

	if err = tx.Commit(ctx); err != nil {
		errr := errs.InternalServerError("Failed to commit transaction: ", err.Error())
		return nil, &errr
	}

	return &output, nil
}
package registration

import (
	"context"
	"errors"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/jackc/pgx/v5"
)

func (r *RegistrationRepository) GetRegistrationByPaymentIntentID(ctx context.Context, paymentIntentID string, AcceptLanguage string) (*models.Registration, error) {
	var titleEN string
	var titleTH *string

	query, err := schema.ReadSQLBaseScript("get_by_payment_intent_id.sql", SqlRegistrationFiles)
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	row := r.db.QueryRow(ctx, query, paymentIntentID)

	var registration models.Registration
	err = row.Scan(
		&registration.ID,
		&registration.ChildID,
		&registration.GuardianID,
		&registration.EventOccurrenceID,
		&registration.Status,
		&registration.StripePaymentIntentID,
		&registration.StripeCustomerID,
		&registration.OrgStripeAccountID,
		&registration.StripePaymentMethodID,
		&registration.TotalAmount,
		&registration.ProviderAmount,
		&registration.PlatformFeeAmount,
		&registration.Currency,
		&registration.PaymentIntentStatus,
		&registration.PaidAt,
		&registration.CancelledAt,
		&registration.CreatedAt,
		&registration.UpdatedAt,
		&titleEN,
		&titleTH,
		&registration.OccurrenceStartTime,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			errr := errs.NotFound("Registration", "stripe_payment_intent_id", paymentIntentID)
			return nil, &errr
		}
		errr := errs.InternalServerError("Failed to fetch registration by payment intent ID: ", err.Error())
		return nil, &errr
	}

	switch AcceptLanguage {
	case "th-TH":
		registration.EventName = *titleTH
	case "en-US":
		registration.EventName = titleEN
	}

	return &registration, nil
}

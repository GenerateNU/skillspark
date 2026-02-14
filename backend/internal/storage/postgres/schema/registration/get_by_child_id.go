package registration

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/jackc/pgx/v5"
)

func (r *RegistrationRepository) GetRegistrationsByChildID(ctx context.Context, input *models.GetRegistrationsByChildIDInput) (*models.GetRegistrationsByChildIDOutput, error) {
	query, err := schema.ReadSQLBaseScript("registration/sql/get_by_child_id.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	rows, err := r.db.Query(ctx, query, input.ChildID)
	if err != nil {
		errr := errs.InternalServerError("Failed to get registrations by child id: ", err.Error())
		return nil, &errr
	}
	defer rows.Close()

	registrations, err := pgx.CollectRows(rows, scanRegistration)
	if err != nil {
		errr := errs.InternalServerError("Failed to collect registrations: ", err.Error())
		return nil, &errr
	}

	var output models.GetRegistrationsByChildIDOutput

	output.Body.Registrations = registrations

	return &output, nil
}

func scanRegistration(row pgx.CollectableRow) (models.Registration, error) {
	var registration models.Registration
	err := row.Scan(
		&registration.ID,
		&registration.ChildID,
		&registration.GuardianID,
		&registration.EventOccurrenceID,
		&registration.Status,
		&registration.CreatedAt,
		&registration.UpdatedAt,
		&registration.StripeCustomerID,
		&registration.OrgStripeAccountID,
		&registration.Currency,
		&registration.PaymentIntentStatus,
		&registration.CancelledAt,
		&registration.StripePaymentIntentID,
		&registration.TotalAmount,
		&registration.ProviderAmount,
		&registration.PlatformFeeAmount,
		&registration.PaidAt,
		&registration.StripePaymentMethodID,
		&registration.EventName,
		&registration.OccurrenceStartTime,
	)
	return registration, err
}
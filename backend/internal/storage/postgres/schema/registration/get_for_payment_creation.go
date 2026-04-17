package registration

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/jackc/pgx/v5"
)

func (r *RegistrationRepository) GetRegistrationsForPaymentCreation(ctx context.Context) ([]models.RegistrationForPayment, error) {
	query, err := schema.ReadSQLBaseScript("get_for_payment_creation.sql", SqlRegistrationFiles)
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		errr := errs.InternalServerError("Failed to get registrations for payment creation: ", err.Error())
		return nil, &errr
	}
	defer rows.Close()

	registrations, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (models.RegistrationForPayment, error) {
		var reg models.RegistrationForPayment
		err := row.Scan(&reg.ID, &reg.GuardianID, &reg.EventOccurrenceID)
		return reg, err
	})
	if err != nil {
		errr := errs.InternalServerError("Failed to collect registrations for payment creation: ", err.Error())
		return nil, &errr
	}

	return registrations, nil
}

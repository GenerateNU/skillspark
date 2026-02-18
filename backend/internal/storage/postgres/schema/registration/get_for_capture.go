package registration

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
	"time"

	"github.com/jackc/pgx/v5"
)

func (r *RegistrationRepository) GetRegistrationsForCapture(ctx context.Context, startWindow time.Time, endWindow time.Time) ([]models.Registration, error) {
	query, err := schema.ReadSQLBaseScript("registration/sql/get_for_capture.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	rows, err := r.db.Query(ctx, query, startWindow, endWindow)
	if err != nil {
		errr := errs.InternalServerError("Failed to get registrations for capture: ", err.Error())
		return nil, &errr
	}
	defer rows.Close()

	registrations, err := pgx.CollectRows(rows, scanRegistration)
	if err != nil {
		errr := errs.InternalServerError("Failed to collect registrations: ", err.Error())
		return nil, &errr
	}

	return registrations, nil
}
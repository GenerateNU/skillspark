package registration

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/jackc/pgx/v5"
)

func (r *RegistrationRepository) GetRegistrationsByEventOccurrenceID(ctx context.Context, input *models.GetRegistrationsByEventOccurrenceIDInput) (*models.GetRegistrationsByEventOccurrenceIDOutput, error) {
	query, err := schema.ReadSQLBaseScript("registration/sql/get_by_event_occurrence_id.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	rows, err := r.db.Query(ctx, query, input.EventOccurrenceID)
	if err != nil {
		errr := errs.InternalServerError("Failed to get registrations by event occurrence id: ", err.Error())
		return nil, &errr
	}
	defer rows.Close()

	registrations, err := pgx.CollectRows(rows, scanRegistration)
	if err != nil {
		errr := errs.InternalServerError("Failed to collect registrations: ", err.Error())
		return nil, &errr
	}

	var output models.GetRegistrationsByEventOccurrenceIDOutput

	output.Body.Registrations = registrations

	return &output, nil
}

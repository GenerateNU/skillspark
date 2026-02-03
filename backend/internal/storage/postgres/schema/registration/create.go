package registration

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *RegistrationRepository) CreateRegistration(ctx context.Context, input *models.CreateRegistrationInput) (*models.CreateRegistrationOutput, error) {

	query, err := schema.ReadSQLBaseScript("registration/sql/create.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	row := r.db.QueryRow(ctx, query,
		input.Body.ChildID,
		input.Body.GuardianID,
		input.Body.EventOccurrenceID,
		input.Body.Status)

	var createdRegistration models.CreateRegistrationOutput

	err = row.Scan(
		&createdRegistration.Body.ID,
		&createdRegistration.Body.ChildID,
		&createdRegistration.Body.GuardianID,
		&createdRegistration.Body.EventOccurrenceID,
		&createdRegistration.Body.Status,
		&createdRegistration.Body.CreatedAt,
		&createdRegistration.Body.UpdatedAt,
		&createdRegistration.Body.EventName,
		&createdRegistration.Body.OccurrenceStartTime,
	)

	if err != nil {
		errr := errs.InternalServerError("Failed to create registration: ", err.Error())
		return nil, &errr
	}

	return &createdRegistration, nil
}

package registration

import (
	"context"
	"errors"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/jackc/pgx/v5"
)

func (r *RegistrationRepository) GetRegistrationByID(ctx context.Context, input *models.GetRegistrationByIDInput, tx *pgx.Tx) (*models.GetRegistrationByIDOutput, error) {
	query, err := schema.ReadSQLBaseScript("get_by_id.sql", SqlRegistrationFiles)
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	var row pgx.Row
	if tx != nil {
		row = (*tx).QueryRow(ctx, query, input.ID)
	} else {
		row = r.db.QueryRow(ctx, query, input.ID)
	}

	var registration models.GetRegistrationByIDOutput

	err = row.Scan(
		&registration.Body.ID,
		&registration.Body.ChildID,
		&registration.Body.GuardianID,
		&registration.Body.EventOccurrenceID,
		&registration.Body.Status,
		&registration.Body.CreatedAt,
		&registration.Body.UpdatedAt,
		&registration.Body.EventName,
		&registration.Body.OccurrenceStartTime,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err := errs.NotFound("Registration", "id", input.ID)
			return nil, &err
		}
		err := errs.InternalServerError("Failed to fetch registration by id: ", err.Error())
		return nil, &err
	}

	return &registration, nil
}

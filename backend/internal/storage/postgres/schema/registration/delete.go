package registration

import (
	"context"
	"errors"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/jackc/pgx/v5"
)

func (r *RegistrationRepository) DeleteRegistration(ctx context.Context, input *models.DeleteRegistrationInput) (*models.DeleteRegistrationOutput, error) {
	query, err := schema.ReadSQLBaseScript("registration/sql/delete.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	row := r.db.QueryRow(ctx, query, input.ID)

	var deleted models.DeleteRegistrationOutput

	err = row.Scan(
		&deleted.Body.ID,
		&deleted.Body.ChildID,
		&deleted.Body.GuardianID,
		&deleted.Body.EventOccurrenceID,
		&deleted.Body.Status,
		&deleted.Body.CreatedAt,
		&deleted.Body.UpdatedAt,
		&deleted.Body.EventName,
		&deleted.Body.OccurrenceStartTime,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err := errs.NotFound("Registration", "id", input.ID)
			return nil, &err
		}
		err := errs.InternalServerError("Failed to delete registration: ", err.Error())
		return nil, &err
	}

	return &deleted, nil
}
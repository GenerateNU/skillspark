package registration

import (
	"context"
	"errors"
	"log/slog"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *RegistrationRepository) UpdateRegistration(ctx context.Context, input *models.UpdateRegistrationInput) (*models.UpdateRegistrationOutput, error) {
	query, err := schema.ReadSQLBaseScript("registration/sql/update.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, errs.InternalServerError("Failed to begin transaction: ", err.Error())
	}

	getInput := &models.GetRegistrationByIDInput{
		ID: input.ID,
	}
	existingOutput, httpErr := r.GetRegistrationByID(ctx, getInput, &tx)
	if httpErr != nil {
		if err := tx.Rollback(ctx); err != nil {
			slog.Error("Failed to rollback transaction: " + err.Error())
		}
		return nil, httpErr
	}

	existing := existingOutput.Body

	if input.Body.ChildID != nil {
		existing.ChildID = *input.Body.ChildID
	}
	if input.Body.GuardianID != nil {
		existing.GuardianID = *input.Body.GuardianID
	}
	if input.Body.EventOccurrenceID != nil {
		existing.EventOccurrenceID = *input.Body.EventOccurrenceID
	}
	if input.Body.Status != nil {
		if *input.Body.Status == models.RegistrationStatusCancelled && existing.Status != models.RegistrationStatusCancelled {
			err = decreaseEventOccurrenceAttendeeCount(ctx, existing.EventOccurrenceID, tx)
			if err != nil {
				if err := tx.Rollback(ctx); err != nil {
					slog.Error("Failed to rollback transaction: " + err.Error())
				}
				return nil, errs.InternalServerError("Failed to decrease event occurrence attendee count: ", err.Error())
			}
		}
		existing.Status = *input.Body.Status
	}

	row := tx.QueryRow(ctx, query,
		existing.ChildID,
		existing.GuardianID,
		existing.EventOccurrenceID,
		existing.Status,
		input.ID,
	)

	var updated models.UpdateRegistrationOutput

	err = row.Scan(
		&updated.Body.ID,
		&updated.Body.ChildID,
		&updated.Body.GuardianID,
		&updated.Body.EventOccurrenceID,
		&updated.Body.Status,
		&updated.Body.CreatedAt,
		&updated.Body.UpdatedAt,
		&updated.Body.EventName,
		&updated.Body.OccurrenceStartTime,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			errr := errs.NotFound("Registration", "id", input.ID)
			if err := tx.Rollback(ctx); err != nil {
				slog.Error("Failed to rollback transaction: " + err.Error())
			}
			return nil, &errr
		}
		errr := errs.InternalServerError("Failed to update registration: ", err.Error())
		if err := tx.Rollback(ctx); err != nil {
			slog.Error("Failed to rollback transaction: " + err.Error())
		}
		return nil, &errr
	}

	if err := tx.Commit(ctx); err != nil {
		slog.Error("Failed to commit transaction: " + err.Error())
		if err := tx.Rollback(ctx); err != nil {
			slog.Error("Failed to rollback transaction: " + err.Error())
		}
		return nil, errs.InternalServerError("Failed to commit transaction: ", err.Error())
	}
	return &updated, nil
}

func decreaseEventOccurrenceAttendeeCount(ctx context.Context, eventOccurrenceID uuid.UUID, tx pgx.Tx) error {
	decrementEventOccurrenceQuery, err := schema.ReadSQLBaseScript("registration/sql/change_event_occurrence_by.sql")
	if err != nil {
		return errs.InternalServerError("Failed to read base query: ", err.Error())
	}
	_, err = tx.Exec(ctx, decrementEventOccurrenceQuery, eventOccurrenceID, -1)
	if err != nil {
		return errs.InternalServerError("Failed to decrement event occurrence attendee count: ", err.Error())
	}

	return nil
}

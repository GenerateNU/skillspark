package eventoccurrence

import (
	"context"
	"errors"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *EventOccurrenceRepository) GetEventOccurrenceByID(ctx context.Context, id uuid.UUID) (*models.EventOccurrence, *errs.HTTPError) {
	query, err := schema.ReadSQLBaseScript("event-occurrence/sql/get_by_id.sql")
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	row := r.db.QueryRow(ctx, query, id)
	var eventOccurrence models.EventOccurrence
	err = row.Scan(&eventOccurrence.ID, 
		&eventOccurrence.ManagerId,
		&eventOccurrence.Event.ID,
		&eventOccurrence.Location.ID,
		&eventOccurrence.StartTime,
		&eventOccurrence.EndTime,
		&eventOccurrence.MaxAttendees,
		&eventOccurrence.Language,
		&eventOccurrence.CurrEnrolled,
		&eventOccurrence.CreatedAt,
		&eventOccurrence.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err := errs.NotFound("Event Occurrence", "id", id)
			return nil, &err
		}
		err := errs.InternalServerError("Failed to fetch event occurrence by id: ", err.Error())
		return nil, &err
	}

	return &eventOccurrence, nil
}
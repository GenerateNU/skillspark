package eventoccurrence

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *EventOccurrenceRepository) GetEventOccurrencesByEventID(ctx context.Context, event_id uuid.UUID) ([]models.EventOccurrence, *errs.HTTPError) {
	query, err := schema.ReadSQLBaseScript("event-occurrence/sql/get_by_event_id.sql")
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	rows, err := r.db.Query(ctx, query, event_id)
	if err != nil {
		err := errs.InternalServerError("Failed to fetch event occurrences by event id: ", err.Error())
		return nil, &err
	}
	defer rows.Close()

	eventOccurrences, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.EventOccurrence])
	if err != nil {
		err := errs.InternalServerError("Failed to scan event occurrence: ", err.Error())
		return nil, &err
	}
	return eventOccurrences, nil
}
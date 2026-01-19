package eventoccurrence

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *EventOccurrenceRepository) CreateEventOccurrence(ctx context.Context, eventoccurrence *models.CreateEventOccurrenceInput) (*models.EventOccurrence, *errs.HTTPError) {
	query, err := schema.ReadSQLBaseScript("event-occurrence/sql/create.sql")
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	row := r.db.QueryRow(ctx, 
		query, 
		eventoccurrence.Body.ManagerId,
		eventoccurrence.Body.EventId,
		eventoccurrence.Body.LocationId,
		eventoccurrence.Body.StartTime,
		eventoccurrence.Body.EndTime,
		eventoccurrence.Body.MaxAttendees,
		eventoccurrence.Body.Language,
	)
	var createdEventOccurrence models.EventOccurrence
	err = row.Scan(&createdEventOccurrence.ID, 
		&createdEventOccurrence.ManagerId,
		&createdEventOccurrence.Event.ID,
		&createdEventOccurrence.Location.ID,
		&createdEventOccurrence.StartTime,
		&createdEventOccurrence.EndTime,
		&createdEventOccurrence.MaxAttendees,
		&createdEventOccurrence.Language,
		&createdEventOccurrence.CurrEnrolled,
		&createdEventOccurrence.CreatedAt,
		&createdEventOccurrence.UpdatedAt,
	)
	if err != nil {
		err := errs.InternalServerError("Failed to create event occurrence: ", err.Error())
		return nil, &err
	}

	return &createdEventOccurrence, nil
}
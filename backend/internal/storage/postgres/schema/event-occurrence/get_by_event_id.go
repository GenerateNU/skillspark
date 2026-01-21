package eventoccurrence

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
)

func (r *EventOccurrenceRepository) GetEventOccurrencesByEventID(ctx context.Context, event_id uuid.UUID) ([]models.EventOccurrence, error) {
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

	eventOccurrences := make([]models.EventOccurrence, 0)
	// populate data from each row individually
	for rows.Next() {
		var createdEventOccurrence models.EventOccurrence
		err := rows.Scan(
			// event occurrence fields
			&createdEventOccurrence.ID,
			&createdEventOccurrence.ManagerId,
			&createdEventOccurrence.StartTime,
			&createdEventOccurrence.EndTime,
			&createdEventOccurrence.MaxAttendees,
			&createdEventOccurrence.Language,
			&createdEventOccurrence.CurrEnrolled,
			&createdEventOccurrence.CreatedAt,
			&createdEventOccurrence.UpdatedAt,

			// event fields
			&createdEventOccurrence.Event.ID,
			&createdEventOccurrence.Event.Title,
			&createdEventOccurrence.Event.Description,
			&createdEventOccurrence.Event.OrganizationID,
			&createdEventOccurrence.Event.AgeRangeMin,
			&createdEventOccurrence.Event.AgeRangeMax,
			&createdEventOccurrence.Event.Category,
			&createdEventOccurrence.Event.HeaderImageS3Key,
			&createdEventOccurrence.Event.CreatedAt,
			&createdEventOccurrence.Event.UpdatedAt,

			// location fields
			&createdEventOccurrence.Location.ID,
			&createdEventOccurrence.Location.Latitude,
			&createdEventOccurrence.Location.Longitude,
			&createdEventOccurrence.Location.AddressLine1,
			&createdEventOccurrence.Location.AddressLine2,
			&createdEventOccurrence.Location.Subdistrict,
			&createdEventOccurrence.Location.District,
			&createdEventOccurrence.Location.Province,
			&createdEventOccurrence.Location.PostalCode,
			&createdEventOccurrence.Location.Country,
			&createdEventOccurrence.Location.CreatedAt,
			&createdEventOccurrence.Location.UpdatedAt,
		)
		if err != nil {
			err := errs.InternalServerError("Failed to scan event occurrence: ", err.Error())
			return nil, &err
		}
		eventOccurrences = append(eventOccurrences, createdEventOccurrence)
	}

	if err := rows.Err(); err != nil {
		err := errs.InternalServerError("Row iteration error: ", err.Error())
		return nil, &err
	}
	return eventOccurrences, nil
}
package eventoccurrence

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *EventOccurrenceRepository) CreateEventOccurrence(ctx context.Context, input *models.CreateEventOccurrenceInput) (*models.EventOccurrence, error) {
	query, err := schema.ReadSQLBaseScript("event-occurrence/sql/create.sql")
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	row := r.db.QueryRow(ctx,
		query,
		input.Body.ManagerId,
		input.Body.EventId,
		input.Body.LocationId,
		input.Body.StartTime,
		input.Body.EndTime,
		input.Body.MaxAttendees,
		input.Body.Language,
		input.Body.Price,
	)
	var createdEventOccurrence models.EventOccurrence

	// populate data in struct, embedding event and location data
	err = row.Scan(
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
		&createdEventOccurrence.Status,
		&createdEventOccurrence.Price,

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
		err := errs.InternalServerError("Failed to create event occurrence: ", err.Error())
		return nil, &err
	}

	return &createdEventOccurrence, nil
}
package eventoccurrence

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/jackc/pgx/v5"
)

func (r *EventOccurrenceRepository) UpdateEventOccurrence(ctx context.Context, input *models.UpdateEventOccurrenceInput, tx *pgx.Tx) (*models.EventOccurrence, error) {
	query, err := schema.ReadSQLBaseScript("update.sql", SqlEventOccurrenceFiles)
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	var row pgx.Row
	if tx != nil {
		row = (*tx).QueryRow(ctx,
			query,
			input.ID,
			input.Body.ManagerId,
			input.Body.EventId,
			input.Body.LocationId,
			input.Body.StartTime,
			input.Body.EndTime,
			input.Body.MaxAttendees,
			input.Body.Language,
			input.Body.CurrEnrolled,
		)
	} else {
		row = r.db.QueryRow(ctx,
			query,
			input.ID,
			input.Body.ManagerId,
			input.Body.EventId,
			input.Body.LocationId,
			input.Body.StartTime,
			input.Body.EndTime,
			input.Body.MaxAttendees,
			input.Body.Language,
			input.Body.CurrEnrolled,
		)
	}

	// null fields are handled in SQL query with coalesce
	var updatedEventOccurrence models.EventOccurrence
	var titleEN, descriptionEN string
	var titleTH, descriptionTH *string

	// populate data in struct, embedding event and location data
	err = row.Scan(
		// event occurrence fields
		&updatedEventOccurrence.ID,
		&updatedEventOccurrence.ManagerId,
		&updatedEventOccurrence.StartTime,
		&updatedEventOccurrence.EndTime,
		&updatedEventOccurrence.MaxAttendees,
		&updatedEventOccurrence.Language,
		&updatedEventOccurrence.CurrEnrolled,
		&updatedEventOccurrence.CreatedAt,
		&updatedEventOccurrence.UpdatedAt,
		&updatedEventOccurrence.Status,

		// event fields
		&updatedEventOccurrence.Event.ID,
		&titleEN,
		&titleTH,
		&descriptionEN,
		&descriptionTH,
		&updatedEventOccurrence.Event.OrganizationID,
		&updatedEventOccurrence.Event.AgeRangeMin,
		&updatedEventOccurrence.Event.AgeRangeMax,
		&updatedEventOccurrence.Event.Category,
		&updatedEventOccurrence.Event.HeaderImageS3Key,
		&updatedEventOccurrence.Event.CreatedAt,
		&updatedEventOccurrence.Event.UpdatedAt,

		// location fields
		&updatedEventOccurrence.Location.ID,
		&updatedEventOccurrence.Location.Latitude,
		&updatedEventOccurrence.Location.Longitude,
		&updatedEventOccurrence.Location.AddressLine1,
		&updatedEventOccurrence.Location.AddressLine2,
		&updatedEventOccurrence.Location.Subdistrict,
		&updatedEventOccurrence.Location.District,
		&updatedEventOccurrence.Location.Province,
		&updatedEventOccurrence.Location.PostalCode,
		&updatedEventOccurrence.Location.Country,
		&updatedEventOccurrence.Location.CreatedAt,
		&updatedEventOccurrence.Location.UpdatedAt,
	)

	// Default to English
	updatedEventOccurrence.Event.Title = titleEN
	updatedEventOccurrence.Event.Description = descriptionEN

	if err != nil {
		err := errs.InternalServerError("Failed to update event occurrence: ", err.Error())
		return nil, &err
	}

	return &updatedEventOccurrence, nil
}

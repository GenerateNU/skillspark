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

func (r *EventOccurrenceRepository) GetEventOccurrenceByID(ctx context.Context, id uuid.UUID) (*models.EventOccurrence, error) {
	query, err := schema.ReadSQLBaseScript("get_by_id.sql", SqlEventOccurrenceFiles)
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	row := r.db.QueryRow(ctx, query, id)
	var eventOccurrence models.EventOccurrence
	var titleEN, descriptionEN string
	var titleTH, descriptionTH *string
	// populate data in struct, embedding event and location data
	err = row.Scan(
		// event occurrence fields
		&eventOccurrence.ID,
		&eventOccurrence.ManagerId,
		&eventOccurrence.StartTime,
		&eventOccurrence.EndTime,
		&eventOccurrence.MaxAttendees,
		&eventOccurrence.Language,
		&eventOccurrence.CurrEnrolled,
		&eventOccurrence.CreatedAt,
		&eventOccurrence.UpdatedAt,
		&eventOccurrence.Status,

		// event fields
		&eventOccurrence.Event.ID,
		&titleEN,
		&titleTH,
		&descriptionEN,
		&descriptionTH,
		&eventOccurrence.Event.OrganizationID,
		&eventOccurrence.Event.AgeRangeMin,
		&eventOccurrence.Event.AgeRangeMax,
		&eventOccurrence.Event.Category,
		&eventOccurrence.Event.HeaderImageS3Key,
		&eventOccurrence.Event.CreatedAt,
		&eventOccurrence.Event.UpdatedAt,

		// location fields
		&eventOccurrence.Location.ID,
		&eventOccurrence.Location.Latitude,
		&eventOccurrence.Location.Longitude,
		&eventOccurrence.Location.AddressLine1,
		&eventOccurrence.Location.AddressLine2,
		&eventOccurrence.Location.Subdistrict,
		&eventOccurrence.Location.District,
		&eventOccurrence.Location.Province,
		&eventOccurrence.Location.PostalCode,
		&eventOccurrence.Location.Country,
		&eventOccurrence.Location.CreatedAt,
		&eventOccurrence.Location.UpdatedAt,
	)

	// Default to English
	eventOccurrence.Event.Title = titleEN
	eventOccurrence.Event.Description = descriptionEN

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

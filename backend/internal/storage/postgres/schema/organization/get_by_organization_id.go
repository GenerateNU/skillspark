package organization

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *OrganizationRepository) GetEventOccurrencesByOrganizationID(ctx context.Context, organization_id uuid.UUID) ([]models.EventOccurrence, error) {
	query, err := schema.ReadSQLBaseScript("get_by_organization_id.sql", SqlOrganizationFiles)
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	rows, err := r.db.Query(ctx, query, organization_id)
	if err != nil {
		err := errs.InternalServerError("Failed to fetch event occurrences by organization id: ", err.Error())
		return nil, &err
	}
	defer rows.Close()

	eventOccurrences, err := pgx.CollectRows(rows, scanEventOccurrence)
	if err != nil {
		err := errs.InternalServerError("Failed to scan all event occurrences: ", err.Error())
		return nil, &err
	}
	return eventOccurrences, nil
}

func scanEventOccurrence(row pgx.CollectableRow) (models.EventOccurrence, error) {
	var eventOccurrence models.EventOccurrence
	var titleEN, descriptionEN string
	var titleTH, descriptionTH *string
	// populate data from each row
	err := row.Scan(
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

	return eventOccurrence, err
}

package event

import (
	"context"
	"encoding/json"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *EventRepository) GetEventOccurrencesByEventID(ctx context.Context, event_id uuid.UUID, AcceptLanguage string) ([]models.EventOccurrence, error) {
	query, err := schema.ReadSQLBaseScript("get_by_event_id.sql", SqlEventFiles)
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

	eventOccurrences, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (models.EventOccurrence, error) {
		return scanEventOccurrenceByEventID(row, AcceptLanguage)
	})
	if err != nil {
		err := errs.InternalServerError("Failed to scan all event occurrences: ", err.Error())
		return nil, &err
	}
	return eventOccurrences, nil
}

func scanEventOccurrenceByEventID(row pgx.CollectableRow, language string) (models.EventOccurrence, error) {
	var createdEventOccurrence models.EventOccurrence
	var titleEN, descriptionEN string
	var titleTH, descriptionTH *string
	var orgLinks []byte
	err := row.Scan(
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
		&createdEventOccurrence.Currency,

		// event fields
		&createdEventOccurrence.Event.ID,
		&titleEN,
		&titleTH,
		&descriptionEN,
		&descriptionTH,
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

		&orgLinks,
	)
	if err != nil {
		return createdEventOccurrence, err
	}

	switch language {
	case "th-TH":
		if titleTH != nil {
			createdEventOccurrence.Event.Title = *titleTH
		} else {
			createdEventOccurrence.Event.Title = titleEN
		}
		if descriptionTH != nil {
			createdEventOccurrence.Event.Description = *descriptionTH
		} else {
			createdEventOccurrence.Event.Description = descriptionEN
		}
	default:
		createdEventOccurrence.Event.Title = titleEN
		createdEventOccurrence.Event.Description = descriptionEN
	}

	if orgLinks != nil {
		if jsonErr := json.Unmarshal(orgLinks, &createdEventOccurrence.OrgLinks); jsonErr != nil {
			createdEventOccurrence.OrgLinks = []models.OrgLink{}
		}
	} else {
		createdEventOccurrence.OrgLinks = []models.OrgLink{}
	}

	return createdEventOccurrence, nil
}

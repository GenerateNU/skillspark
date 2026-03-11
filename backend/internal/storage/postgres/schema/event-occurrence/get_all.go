package eventoccurrence

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
	"skillspark/internal/utils"

	"github.com/jackc/pgx/v5"
)

<<<<<<< Alen-Ganopolsky/Backend-Translation
var language string

func (r *EventOccurrenceRepository) GetAllEventOccurrences(ctx context.Context, pagination utils.Pagination, AcceptLanguage string, filters models.GetAllEventOccurrencesFilter) ([]models.EventOccurrence, error) {
	language = AcceptLanguage
=======
type PriceRange struct {
	MinPrice int
	MaxPrice *int
}

func (r *EventOccurrenceRepository) GetAllEventOccurrences(ctx context.Context, pagination utils.Pagination, filters models.GetAllEventOccurrencesFilter) ([]models.EventOccurrence, error) {

>>>>>>> main
	query, err := schema.ReadSQLBaseScript("get_all.sql", SqlEventOccurrenceFiles)
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	rows, err := r.db.Query(
		ctx,
		query,
		pagination.Limit,
		pagination.GetOffset(),
		filters.Search,
		filters.MinDurationMinutes,
		filters.MaxDurationMinutes,
		filters.Latitude,
		filters.Longitude,
		filters.RadiusKm,
		filters.MinAge,
		filters.MaxAge,
		filters.Category,
		filters.SoldOut,
		filters.MinDate,
		filters.MaxDate,
		filters.MinPrice,
		filters.MaxPrice,
	)

	if err != nil {
		err := errs.InternalServerError("Failed to fetch all event occurrences: ", err.Error())
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
	var createdEventOccurrence models.EventOccurrence
	var titleEN, descriptionEN string
	var titleTH, descriptionTH *string
	// populate data from each row
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
	)

	switch language {
	case "th-TH":
		createdEventOccurrence.Event.Title = *titleTH
		createdEventOccurrence.Event.Description = *descriptionTH
	case "en-US":
		createdEventOccurrence.Event.Title = titleEN
		createdEventOccurrence.Event.Description = descriptionEN
	}

	return createdEventOccurrence, err
}

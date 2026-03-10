package eventoccurrence

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
	"skillspark/internal/utils"

	"github.com/jackc/pgx/v5"
)

type PriceRange struct {
	MinPrice int
	MaxPrice *int
}

func (r *EventOccurrenceRepository) GetAllEventOccurrences(ctx context.Context, pagination utils.Pagination, filters models.GetAllEventOccurrencesFilter) ([]models.EventOccurrence, error) {

	query, err := schema.ReadSQLBaseScript("get_all.sql", SqlEventOccurrenceFiles)
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	max3000 := 3000
	max6000 := 6000

	// this part determines what the price range is, for $, $$ and $$$
	priceRange := map[string]PriceRange{
		"$":   {MinPrice: 0, MaxPrice: &max3000},
		"$$":  {MinPrice: 3001, MaxPrice: &max6000},
		"$$$": {MinPrice: 6001, MaxPrice: nil}, // no upper bound
	}

	var min interface{} = nil
	var max interface{} = nil

	if filters.PriceTier != nil {
		rng := priceRange[*filters.PriceTier]

		min = rng.MinPrice

		if rng.MaxPrice != nil {
			max = *rng.MaxPrice
		}
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
		min, // 15
		max, // 16
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
	return createdEventOccurrence, err
}

package recommendation

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
	"skillspark/internal/utils"
	"time"

	"github.com/jackc/pgx/v5"
)

func (r *RecommendationRepository) GetRecommendationsByChildID(ctx context.Context, childInterests []string, childBirthYear int, acceptLanguage string, pagination utils.Pagination, filters models.RecommendationFilters) ([]models.Event, error) {
	query, err := schema.ReadSQLBaseScript("get_by_child_id.sql", SqlRecommendationFiles)
	if err != nil {
		e := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &e
	}

	var lat, lng *float64
	if filters.Latitude.Set {
		lat = &filters.Latitude.Value
	}
	if filters.Longitude.Set {
		lng = &filters.Longitude.Value
	}

	var minDate, maxDate *time.Time
	if !filters.MinDate.IsZero() {
		minDate = &filters.MinDate
	}
	if !filters.MaxDate.IsZero() {
		maxDate = &filters.MaxDate
	}

	rows, err := r.db.Query(ctx, query, childInterests, childBirthYear, pagination.Limit, pagination.GetOffset(), minDate, maxDate, lat, lng, filters.RadiusKm)
	if err != nil {
		e := errs.InternalServerError("Failed to fetch recommendations: ", err.Error())
		return nil, &e
	}
	defer rows.Close()

	events, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (models.Event, error) {
		return scanEvent(row, acceptLanguage)
	})
	if err != nil {
		e := errs.InternalServerError("Failed to scan recommendations: ", err.Error())
		return nil, &e
	}

	return events, nil
}

func scanEvent(row pgx.CollectableRow, language string) (models.Event, error) {
	var e models.Event
	var titleEN, descriptionEN string
	var titleTH, descriptionTH *string
	var score int

	err := row.Scan(
		&e.ID,
		&titleEN,
		&titleTH,
		&descriptionEN,
		&descriptionTH,
		&e.OrganizationID,
		&e.AgeRangeMin,
		&e.AgeRangeMax,
		&e.Category,
		&e.HeaderImageS3Key,
		&e.CreatedAt,
		&e.UpdatedAt,
		&score,
	)

	e.Title = titleEN
	e.Description = descriptionEN

	if language == "th-TH" {
		if titleTH != nil {
			e.Title = *titleTH
		}
		if descriptionTH != nil {
			e.Description = *descriptionTH
		}
	}

	return e, err
}

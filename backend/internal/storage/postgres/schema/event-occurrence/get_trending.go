package eventoccurrence

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/jackc/pgx/v5"
)

func (r *EventOccurrenceRepository) GetTrendingEventOccurrences(ctx context.Context, input *models.GetTrendingEventOccurrencesInput) ([]models.EventOccurrence, error) {
	lang := input.AcceptLanguage
	query, err := schema.ReadSQLBaseScript("get_trending_eventoccurrences.sql", SqlEventOccurrenceFiles)
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	rows, err := r.db.Query(
		ctx,
		query,
		input.Latitude,
		input.Longitude,
		input.Radius,
		input.MaxReturns,
	)

	if err != nil {
		err := errs.InternalServerError("Failed to fetch all event occurrences: ", err.Error())
		return nil, &err
	}
	defer rows.Close()

	eventOccurrences, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (models.EventOccurrence, error) {
		return scanEventOccurrence(row, lang)
	})
	if err != nil {
		err := errs.InternalServerError("Failed to scan all event occurrences: ", err.Error())
		return nil, &err
	}
	return eventOccurrences, nil
}

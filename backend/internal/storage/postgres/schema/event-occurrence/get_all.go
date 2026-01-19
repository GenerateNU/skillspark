package eventoccurrence

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
	"skillspark/internal/utils"

	"github.com/jackc/pgx/v5"
)

func (r *EventOccurrenceRepository) GetAllEventOccurrences(ctx context.Context, pagination utils.Pagination) ([]models.EventOccurrence, *errs.HTTPError) {
	query, err := schema.ReadSQLBaseScript("event-occurrence/sql/get_all.sql")
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	rows, err := r.db.Query(ctx, query, pagination.Limit, pagination.GetOffset())
	if err != nil {
		err := errs.InternalServerError("Failed to fetch all event occurrences: ", err.Error())
		return nil, &err
	}
	defer rows.Close()

	eventOccurrences, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.EventOccurrence])
	if err != nil {
		err := errs.InternalServerError("Failed to scan event occurrence: ", err.Error())
		return nil, &err
	}
	return eventOccurrences, nil
}
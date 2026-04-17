package event

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
	"skillspark/internal/utils"

	"github.com/jackc/pgx/v5"
)

func (r *EventRepository) GetAllEvents(ctx context.Context, pagination utils.Pagination, AcceptLanguage string, filters models.GetAllEventsFilter) ([]models.Event, error) {
	query, err := schema.ReadSQLBaseScript("get_all.sql", SqlEventFiles)
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	rows, err := r.db.Query(ctx, query,
		pagination.Limit,
		pagination.GetOffset(),
		filters.Search,
		filters.Category,
		filters.MinAge,
		filters.MaxAge,
	)
	if err != nil {
		err := errs.InternalServerError("Failed to fetch all events: ", err.Error())
		return nil, &err
	}
	defer rows.Close()

	events, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (models.Event, error) {
		return scanEvent(row, AcceptLanguage)
	})
	if err != nil {
		err := errs.InternalServerError("Failed to scan all events: ", err.Error())
		return nil, &err
	}
	return events, nil
}

func scanEvent(row pgx.CollectableRow, language string) (models.Event, error) {
	var event models.Event
	var titleEN, descriptionEN string
	var titleTH, descriptionTH *string

	err := row.Scan(
		&event.ID,
		&titleEN,
		&titleTH,
		&descriptionEN,
		&descriptionTH,
		&event.OrganizationID,
		&event.AgeRangeMin,
		&event.AgeRangeMax,
		&event.Category,
		&event.HeaderImageS3Key,
		&event.CreatedAt,
		&event.UpdatedAt,
	)

	switch language {
	case "th-TH":
		if titleTH != nil {
			event.Title = *titleTH
		} else {
			event.Title = titleEN
		}
		if descriptionTH != nil {
			event.Description = *descriptionTH
		} else {
			event.Description = descriptionEN
		}
	default:
		event.Title = titleEN
		event.Description = descriptionEN
	}

	return event, err
}

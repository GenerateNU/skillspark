package event

import (
	"context"
	"errors"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *EventRepository) GetEventByID(ctx context.Context, id uuid.UUID) (*models.Event, error) {
	query, err := schema.ReadSQLBaseScript("event/sql/get_by_id.sql")
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	row := r.db.QueryRow(ctx, query, id)
	var event models.Event
	err = row.Scan(
		&event.ID,
		&event.Title,
		&event.Description,
		&event.OrganizationID,
		&event.AgeRangeMin,
		&event.AgeRangeMax,
		&event.Category,
		&event.HeaderImageS3Key,
		&event.CreatedAt,
		&event.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err := errs.NotFound("Event", "id", id)
			return nil, &err
		}
		err := errs.InternalServerError("Failed to fetch event by id: ", err.Error())
		return nil, &err
	}

	return &event, nil
}

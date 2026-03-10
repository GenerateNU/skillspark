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

func (r *EventRepository) GetEventByID(ctx context.Context, id uuid.UUID, AcceptLanguage string) (*models.Event, error) {

	var titleEN, descriptionEN string
	var titleTH, descriptionTH *string

	query, err := schema.ReadSQLBaseScript("get_by_id.sql", SqlEventFiles)
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	row := r.db.QueryRow(ctx, query, id)
	var event models.Event
	err = row.Scan(
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

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err := errs.NotFound("Event", "id", id)
			return nil, &err
		}
		err := errs.InternalServerError("Failed to fetch event by id: ", err.Error())
		return nil, &err
	}

	switch AcceptLanguage {
	case "th-TH":
		event.Title = *titleTH
		event.Description = *descriptionTH
	case "en-US":
		event.Title = titleEN
		event.Description = descriptionEN
	}

	return &event, nil
}

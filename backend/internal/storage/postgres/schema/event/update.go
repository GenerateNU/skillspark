package event

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *EventRepository) UpdateEvent(ctx context.Context, input *models.UpdateEventInput) (*models.Event, *errs.HTTPError) {
	query, err := schema.ReadSQLBaseScript("event/sql/update.sql")
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	row := r.db.QueryRow(ctx, query, input.ID, input.Body.Title, input.Body.Description, input.Body.OrganizationID, input.Body.AgeRangeMin, input.Body.AgeRangeMax, input.Body.Category, input.Body.HeaderImageS3Key)

	var event models.Event

	err = row.Scan(&event.ID, &event.Title, &event.Description, &event.OrganizationID, &event.AgeRangeMin, &event.AgeRangeMax, &event.Category, &event.HeaderImageS3Key, &event.CreatedAt, &event.UpdatedAt)
	if err != nil {
		err := errs.InternalServerError("Failed to update event: ", err.Error())
		return nil, &err
	}

	return &event, nil
}

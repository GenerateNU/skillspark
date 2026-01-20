package event

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *EventRepository) CreateEvent(ctx context.Context, event *models.CreateEventInput) (*models.Event, error) {
	query, err := schema.ReadSQLBaseScript("event/sql/create.sql")
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	row := r.db.QueryRow(ctx, query, event.Body.Title, event.Body.Description, event.Body.OrganizationID, event.Body.AgeRangeMin, event.Body.AgeRangeMax, event.Body.Category, event.Body.HeaderImageS3Key)

	var createdEvent models.Event

	err = row.Scan(&createdEvent.ID, &createdEvent.Title, &createdEvent.Description, &createdEvent.OrganizationID, &createdEvent.AgeRangeMin, &createdEvent.AgeRangeMax, &createdEvent.Category, &createdEvent.HeaderImageS3Key, &createdEvent.CreatedAt, &createdEvent.UpdatedAt)
	if err != nil {
		err := errs.InternalServerError("Failed to create event: ", err.Error())
		return nil, &err
	}

	return &createdEvent, nil
}

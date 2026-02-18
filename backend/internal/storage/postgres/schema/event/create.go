package event

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *EventRepository) CreateEvent(ctx context.Context, event *models.CreateEventDBInput, HeaderImageS3Key *string) (*models.Event, error) {
	query, err := schema.ReadSQLBaseScript("create.sql", SqlEventFiles)
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	row := r.db.QueryRow(ctx, query, event.Body.Title_EN, event.Body.Title_TH, event.Body.Description_EN, event.Body.Description_TH, event.Body.OrganizationID, event.Body.AgeRangeMin, event.Body.AgeRangeMax, event.Body.Category, HeaderImageS3Key)

	var createdEvent models.Event

	err = row.Scan(&createdEvent.ID, &createdEvent.Title, &createdEvent.Description, &createdEvent.OrganizationID, &createdEvent.AgeRangeMin, &createdEvent.AgeRangeMax, &createdEvent.Category, &createdEvent.HeaderImageS3Key, &createdEvent.CreatedAt, &createdEvent.UpdatedAt)
	if err != nil {
		err := errs.InternalServerError("Failed to create event: ", err.Error())
		return nil, &err
	}

	return &createdEvent, nil
}

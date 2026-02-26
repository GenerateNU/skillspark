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
	var titleEN, titleTH, descEN, descTH string

	err = row.Scan(&createdEvent.ID, &titleEN, &titleTH, &descEN, &descTH, &createdEvent.OrganizationID, &createdEvent.AgeRangeMin, &createdEvent.AgeRangeMax, &createdEvent.Category, &createdEvent.HeaderImageS3Key, &createdEvent.CreatedAt, &createdEvent.UpdatedAt)
	if err != nil {
		err := errs.InternalServerError("Failed to create event: ", err.Error())
		return nil, &err
	}

	switch event.AcceptLanguage {
	case "th-TH":
		createdEvent.Title = titleTH
		createdEvent.Description = descTH
	case "en-US":
		createdEvent.Title = titleEN
		createdEvent.Description = descEN
	}

	return &createdEvent, nil
}

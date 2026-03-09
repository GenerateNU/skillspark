package event

import (
	"context"
	"skillspark/internal/models"
)

func (h *Handler) CreateTranslateStruct(ctx context.Context, event *models.CreateEventInput, translatedTitle *string, translatedDescription *string) *models.CreateEventDBInput {

	eventBody := event.Body

	dbInitInput := &models.CreateEventDBInput{
		AcceptLanguage: event.AcceptLanguage,
		Body: models.CreateDBBody{
			OrganizationID: eventBody.OrganizationID,
			AgeRangeMin:    eventBody.AgeRangeMin,
			AgeRangeMax:    eventBody.AgeRangeMax,
			Category:       eventBody.Category,
		},
	}

	switch event.AcceptLanguage {
	case "th-TH":
		dbInitInput.Body.Title_EN = *translatedTitle
		dbInitInput.Body.Title_TH = &eventBody.Title
		dbInitInput.Body.Description_EN = *translatedDescription
		dbInitInput.Body.Description_TH = &eventBody.Description
	case "en-US":
		dbInitInput.Body.Title_EN = eventBody.Title
		dbInitInput.Body.Title_TH = translatedTitle
		dbInitInput.Body.Description_TH = translatedDescription
		dbInitInput.Body.Description_EN = eventBody.Description
	}

	return dbInitInput
}

func (h *Handler) UpdateTranslateStruct(ctx context.Context, event *models.UpdateEventInput, translatedTitle *string, translatedDescription *string) *models.UpdateEventDBInput {

	eventBody := event.Body

	dbInitInput := &models.UpdateEventDBInput{
		AcceptLanguage: event.AcceptLanguage,
		ID:             event.ID,
		Body: models.UpdateDBBody{
			OrganizationID: eventBody.OrganizationID,
			AgeRangeMin:    eventBody.AgeRangeMin,
			AgeRangeMax:    eventBody.AgeRangeMax,
			Category:       eventBody.Category,
		},
	}

	switch event.AcceptLanguage {
	case "th-TH":
		dbInitInput.Body.Title_EN = translatedTitle
		dbInitInput.Body.Title_TH = eventBody.Title
		dbInitInput.Body.Description_EN = translatedDescription
		dbInitInput.Body.Description_TH = eventBody.Description
	case "en-US":
		dbInitInput.Body.Title_EN = eventBody.Title
		dbInitInput.Body.Title_TH = translatedTitle
		dbInitInput.Body.Description_TH = translatedDescription
		dbInitInput.Body.Description_EN = eventBody.Description
	}

	return dbInitInput
}

package review

import (
	"context"
	"skillspark/internal/models"
)

func (h *Handler) CreateTranslateStruct(ctx context.Context, event *models.CreateReviewInput, translation *string) *models.CreateReviewDBInput {

	eventBody := event.Body

	dbInitInput := &models.CreateReviewDBInput{
		AcceptLanguage: event.AcceptLanguage,
		Body: models.CreateReviewDBBody{
			RegistrationID: eventBody.RegistrationID,
			GuardianID:     eventBody.GuardianID,
			Categories:     eventBody.Categories,
		},
	}

	switch event.AcceptLanguage {
	case "th-TH":
		dbInitInput.Body.Description_EN = *translation
		dbInitInput.Body.Description_TH = &eventBody.Description
	case "en-US":
		dbInitInput.Body.Description_EN = eventBody.Description
		dbInitInput.Body.Description_TH = translation
	}

	return dbInitInput
}

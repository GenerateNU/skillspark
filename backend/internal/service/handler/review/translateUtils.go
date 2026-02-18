package review

import (
	"context"
	"fmt"
	"skillspark/internal/models"
)

func (h *Handler) CallTranslateAPI(ctx context.Context, description_en *string) (*string, error) {

	deref := func(s *string) string {
		if s == nil {
			return ""
		}
		return *s
	}

	description := deref(description_en)

	if description == "" {
		err := fmt.Errorf("no title or description provided")
		return nil, err
	}

	response, err := h.TranslateClient.GetTranslation(ctx, description)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (h *Handler) CreateTranslateStruct(ctx context.Context, event *models.CreateReviewInput, description_th *string) *models.CreateReviewDBInput {

	eventBody := event.Body

	dbInitInput := &models.CreateReviewDBInput{
		Body: models.CreateReviewDBBody{
			RegistrationID: eventBody.RegistrationID,
			GuardianID:     eventBody.GuardianID,
			Description_EN: eventBody.Description,
			Description_TH: description_th,
			Categories:     eventBody.Categories,
		},
	}

	return dbInitInput
}

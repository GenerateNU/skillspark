package review

import (
	"context"
	"fmt"
	"skillspark/internal/models"
)

func (h *Handler) CallTranslateAPI(ctx context.Context, description_en *string, AcceptLanguage string) (*string, error) {
	var sl string
	var dl string

	deref := func(s *string) string {
		if s == nil {
			return ""
		}
		return *s
	}

	description := deref(description_en)

	if AcceptLanguage == "th" {
		sl = "th"
		dl = "en"
	} else {
		sl = "en"
		dl = "th"
	}

	if description == "" {
		err := fmt.Errorf("no title or description provided")
		return nil, err
	}

	response, err := h.TranslateClient.GetTranslation(ctx, description, sl, dl)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (h *Handler) CreateTranslateStruct(ctx context.Context, event *models.CreateReviewInput, description_th *string) *models.CreateReviewDBInput {

	eventBody := event.Body

	dbInitInput := &models.CreateReviewDBInput{
		AcceptLanguage: event.AcceptLanguage,
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

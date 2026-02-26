package review

import (
	"context"
	"fmt"
	"skillspark/internal/models"
)

func (h *Handler) CallTranslateAPI(ctx context.Context, src_description *string, AcceptLanguage string) (*string, error) {
	var sl string
	var dl string

	deref := func(s *string) string {
		if s == nil {
			return ""
		}
		return *s
	}

	description := deref(src_description)

	switch AcceptLanguage {
	case "th-TH":
		sl = "th"
		dl = "en"
	case "en-US":
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

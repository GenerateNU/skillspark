package event

import (
	"context"
	"fmt"
	"skillspark/internal/models"
	"strings"
)

type EventTranslationResponse struct {
	TranslatedTitle       *string `json:"translated_title"`
	TranslatedDescription *string `json:"translated_description"`
}

func (h *Handler) CallTranslateAPI(ctx context.Context, src_title *string, src_description *string, AcceptLanguage string) (*EventTranslationResponse, error) {
	var sl string
	var dl string
	deref := func(s *string) string {
		if s == nil {
			return ""
		}
		return *s
	}

	title_deref := deref(src_title)
	description_deref := deref(src_description)

	if title_deref == "" && description_deref == "" {
		err := fmt.Errorf("no title or description provided")
		return nil, err
	}

	switch AcceptLanguage {
	case "th-TH":
		sl = "th"
		dl = "en"
	case "en-US":
		sl = "en"
		dl = "th"
	}

	translationString := title_deref + "|*|" + description_deref

	response, err := h.TranslateClient.GetTranslation(ctx, translationString, sl, dl)
	if err != nil {
		return nil, err
	}

	parsedResponse := strings.Split(*response, "|*|")
	if len(parsedResponse) != 2 {
		err := fmt.Errorf("unexpected response length")
		return nil, err
	}

	translatedTitle := parsedResponse[0]
	translatedDescription := parsedResponse[1]

	return &EventTranslationResponse{TranslatedTitle: &translatedTitle,
		TranslatedDescription: &translatedDescription}, nil

}

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

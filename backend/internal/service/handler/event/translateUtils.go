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

func (h *Handler) CallTranslateAPI(ctx context.Context, title *string, description *string, AcceptLanguage string) (*EventTranslationResponse, error) {
	var sl string
	var dl string
	deref := func(s *string) string {
		if s == nil {
			return ""
		}
		return *s
	}

	title_t := deref(title)
	description_t := deref(description)

	if title_t == "" && description_t == "" {
		err := fmt.Errorf("no title or description provided")
		return nil, err
	}

	// Always translate from English to Thai
	// Accept-Language is only used for output selection, not input detection
	sl = "en"
	dl = "th"

	translationString := title_t + "|*|" + description_t

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

	// User always sends English, translation is always Thai
	// Accept-Language only controls what gets returned, not storage
	dbInitInput := &models.CreateEventDBInput{
		AcceptLanguage: event.AcceptLanguage,
		Body: models.CreateDBBody{
			Title_EN:       eventBody.Title,
			Title_TH:       translatedTitle,
			Description_EN: eventBody.Description,
			Description_TH: translatedDescription,
			OrganizationID: eventBody.OrganizationID,
			AgeRangeMin:    eventBody.AgeRangeMin,
			AgeRangeMax:    eventBody.AgeRangeMax,
			Category:       eventBody.Category,
		},
	}

	return dbInitInput
}

func (h *Handler) UpdateTranslateStruct(ctx context.Context, event *models.UpdateEventInput, translatedTitle *string, translatedDescription *string) *models.UpdateEventDBInput {

	eventBody := event.Body

	// User always sends English, translation is always Thai
	// Accept-Language only controls what gets returned, not storage
	dbInitInput := &models.UpdateEventDBInput{
		AcceptLanguage: event.AcceptLanguage,
		ID:             event.ID,
		Body: models.UpdateDBBody{
			Title_EN:       eventBody.Title,
			Title_TH:       translatedTitle,
			Description_EN: eventBody.Description,
			Description_TH: translatedDescription,
			OrganizationID: eventBody.OrganizationID,
			AgeRangeMin:    eventBody.AgeRangeMin,
			AgeRangeMax:    eventBody.AgeRangeMax,
			Category:       eventBody.Category,
		},
	}

	return dbInitInput
}

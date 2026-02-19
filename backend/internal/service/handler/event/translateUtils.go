package event

import (
	"context"
	"fmt"
	"skillspark/internal/models"
	"strings"
)

type EventTranslationResponse struct {
	Title_TH       *string `json:"title_th"`
	Description_TH *string `json:"description_th"`
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

	if AcceptLanguage == "th" {
		sl = "th"
		dl = "en"
	} else {
		sl = "en"
		dl = "th"
	}

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

	title_th := parsedResponse[0]
	description_th := parsedResponse[1]

	return &EventTranslationResponse{Title_TH: &title_th,
		Description_TH: &description_th}, nil

}

func (h *Handler) CreateTranslateStruct(ctx context.Context, event *models.CreateEventInput, title_th *string, description_th *string) *models.CreateEventDBInput {

	eventBody := event.Body

	dbInitInput := &models.CreateEventDBInput{
		AcceptLanguage: event.AcceptLanguage,
		Body: models.CreateDBBody{
			Title_EN:       eventBody.Title,
			Title_TH:       title_th,
			Description_EN: eventBody.Description,
			Description_TH: description_th,
			OrganizationID: eventBody.OrganizationID,
			AgeRangeMin:    eventBody.AgeRangeMin,
			AgeRangeMax:    eventBody.AgeRangeMax,
			Category:       eventBody.Category,
		},
	}

	return dbInitInput
}

func (h *Handler) UpdateTranslateStruct(ctx context.Context, event *models.UpdateEventInput, title_th *string, description_th *string) *models.UpdateEventDBInput {

	eventBody := event.Body

	dbInitInput := &models.UpdateEventDBInput{
		AcceptLanguage: event.AcceptLanguage,
		ID:             event.ID,
		Body: models.UpdateDBBody{
			Title_EN:       eventBody.Title,
			Title_TH:       title_th,
			Description_EN: eventBody.Description,
			Description_TH: description_th,
			OrganizationID: eventBody.OrganizationID,
			AgeRangeMin:    eventBody.AgeRangeMin,
			AgeRangeMax:    eventBody.AgeRangeMax,
			Category:       eventBody.Category,
		},
	}

	return dbInitInput
}

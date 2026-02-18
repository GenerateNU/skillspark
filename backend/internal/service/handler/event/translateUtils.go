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

func (h *Handler) CallTranslateAPI(ctx context.Context, title_en *string, description_en *string) (*EventTranslationResponse, error) {

	deref := func(s *string) string {
		if s == nil {
			return ""
		}
		return *s
	}

	title := deref(title_en)
	description := deref(description_en)

	if title == "" && description == "" {
		err := fmt.Errorf("no title or description provided")
		return nil, err
	}

	translationString := title + "|*|" + description
	response, err := h.TranslateClient.GetTranslation(ctx, translationString)
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
		ID: event.ID,
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

package event

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/s3_client"
	"strings"
	"sync"
)

func (h *Handler) UpdateEvent(ctx context.Context, input *models.UpdateEventInput, image_data *[]byte, s3Client s3_client.S3Interface) (*models.Event, error) {

	var key *string
	var url *string
	var wg sync.WaitGroup

	wg.Add(1)
	dbInput, err := h.PatchTranslateAPI(ctx, input)
	if err != nil {
		return nil, err
	}

	if image_data != nil {
		var err error
		url, key, err = h.UpdateEventS3Helper(ctx, s3Client, input, image_data)
		if err != nil {
			return nil, err
		}
	}

	event, err := h.EventRepository.UpdateEvent(ctx, input, key)
	if err != nil {
		return nil, err
	}

	event.PresignedURL = url

	return event, nil
}

func (h *Handler) UpdateEventS3Helper(ctx context.Context, s3Client s3_client.S3Interface, input *models.UpdateEventInput, image_data *[]byte) (*string, *string, error) {
	key, genErr := h.generateS3Key(input.ID)
	if genErr != nil {
		return nil, nil, genErr
	}
	url, errr := s3Client.UploadImage(ctx, key, *image_data)
	if errr != nil {
		return nil, nil, errr
	}

	return url, key, nil
}

func (h *Handler) UpdateInitInput(ctx context.Context, event *models.UpdateEventInput) *models.UpdateEventDBInput {

	eventBody := event.Body

	dbInitInput := &models.UpdateEventDBInput{
		ID: event.ID,
		Body: models.UpdateDBBody{
			Title_EN:       eventBody.Title,
			Title_TH:       nil,
			Description_EN: eventBody.Description,
			Description_TH: nil,
			OrganizationID: eventBody.OrganizationID,
			AgeRangeMin:    eventBody.AgeRangeMin,
			AgeRangeMax:    eventBody.AgeRangeMax,
			Category:       eventBody.Category,
		},
	}

	return dbInitInput
}

func (h *Handler) PatchTranslateAPI(ctx context.Context, event *models.UpdateEventInput) (*models.CreateEventDBInput, error) {

	eventBody := event.Body

	deref := func(s *string) string {
		if s == nil {
			return ""
		}
		return *s
	}

	title := deref(eventBody.Title)
	description := deref(eventBody.Description)

	translationString := title + "***" + description
	response, err := h.TranslateClient.GetTranslation(ctx, translationString)
	if err != nil {
		return nil, err
	}

	parsedResponse := strings.Split(*response, "***")

	dbInput := &models.UpdateEventDBInput{
		ID: event.ID,
		Body: models.UpdateDBBody{
			Title_EN:       eventBody.Title,
			Title_TH:       &parsedResponse[0],
			Description_EN: eventBody.Description,
			Description_TH: &parsedResponse[1],
			OrganizationID: eventBody.OrganizationID,
			AgeRangeMin:    eventBody.AgeRangeMin,
			AgeRangeMax:    eventBody.AgeRangeMax,
			Category:       eventBody.Category,
		},
	}

	return dbInput, nil
}

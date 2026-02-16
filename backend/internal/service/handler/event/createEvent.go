package event

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/s3_client"
	"strings"
	"sync"
)

func (h *Handler) CreateEvent(ctx context.Context, input *models.CreateEventInput, updateBody *models.UpdateEventBody, image_data *[]byte, s3Client s3_client.S3Interface) (*models.Event, error) {
	var key *string
	var url *string
	var wg sync.WaitGroup

	// make initial input struct into db
	initInput := h.CreateInitInput(ctx, input)

	// create event without key of translations
	event, err := h.EventRepository.CreateEvent(ctx, initInput, key)
	if err != nil {
		return nil, err
	}

	// make translation struct
	dbInput, err := h.PostTranslateAPI(ctx, input)
	if err != nil {
		return nil, err
	}

	// asynchronously update event with translations
	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err = h.EventRepository.UpdateEvent(ctx, dbInput, key)
	}()

	// asynchronously make presigned url if image_data is not nil
	if image_data != nil {

		url, err = h.CreateEventS3Helper(ctx, s3Client, event, updateBody, image_data)
		if err != nil {
			return nil, err
		}
	}

	event.PresignedURL = url
	wg.Wait()

	return event, nil
}

// helper for uploading image to s3
func (h *Handler) CreateEventS3Helper(ctx context.Context, s3Client s3_client.S3Interface, event *models.Event,
	updateBody *models.UpdateEventBody, image_data *[]byte) (*string, error) {

	key, err := h.generateS3Key(event.ID)
	if err != nil {
		return nil, err
	}

	url, errr := s3Client.UploadImage(ctx, key, *image_data)

	updateInput := &models.UpdateEventInput{
		ID:   event.ID,
		Body: *updateBody,
	}
	updateKeyValue, err := h.EventRepository.UpdateEvent(ctx, updateInput, key)
	if err != nil {
		return nil, nil
	}
	event.HeaderImageS3Key = updateKeyValue.HeaderImageS3Key
	if errr != nil {
		return nil, errr
	}

	return url, nil

}

func (h *Handler) PostTranslateAPI(ctx context.Context, event *models.CreateEventInput) (*models.CreateEventDBInput, error) {

	eventBody := event.Body
	// create string that includes title and description
	// call translation api to translate string
	// create CreateEventDBInput Struct and return if successful
	translationString := eventBody.Title + "***" + eventBody.Description
	response, err := h.TranslateClient.GetTranslation(ctx, translationString)
	if err != nil {
		return nil, err
	}

	parsedResponse := strings.Split(*response, "***")

	dbInput := &models.CreateEventDBInput{
		Body: models.CreateDBBody{
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

func (h *Handler) CreateInitInput(ctx context.Context, event *models.CreateEventInput) *models.CreateEventDBInput {

	eventBody := event.Body

	dbInitInput := &models.CreateEventDBInput{
		Body: models.CreateDBBody{
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

package event

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/s3_client"
	"sync"
)

func (h *Handler) CreateEvent(ctx context.Context, input *models.CreateEventInput, updateBody *models.UpdateEventBody, imageData *[]byte, s3Client s3_client.S3Interface) (*models.Event, error) {
	var key *string
	var url *string
	var wg sync.WaitGroup

	initInput := h.CreateTranslateStruct(ctx, input, nil, nil)
	event, err := h.EventRepository.CreateEvent(ctx, initInput, key)
	if err != nil {
		return nil, err
	}

	updateInput := &models.UpdateEventInput{
		ID:   event.ID,
		Body: *updateBody,
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		// TODO handle errors in goroutine
		translationResp, _ := h.CallTranslateAPI(ctx, &event.Title, &event.Description)
		translationsReinsertion := h.UpdateTranslateStruct(ctx, updateInput, translationResp.Title_TH, translationResp.Description_TH)
		_, err = h.EventRepository.UpdateEvent(ctx, translationsReinsertion, key)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// TODO handle errors in goroutine
		url, err = h.CreateEventS3Helper(ctx, s3Client, event, updateInput, imageData)
	}()
	event.PresignedURL = url

	wg.Wait()
	return event, nil
}

// helper for uploading image to s3
func (h *Handler) CreateEventS3Helper(ctx context.Context, s3Client s3_client.S3Interface, event *models.Event,
	updateInput *models.UpdateEventInput, imageData *[]byte) (*string, error) {

	if imageData != nil {

		key, err := h.generateS3Key(event.ID)
		if err != nil {
			return nil, err
		}

		url, errr := s3Client.UploadImage(ctx, key, *imageData)

		updateWithKey := h.UpdateTranslateStruct(ctx, updateInput, nil, nil)

		updateKeyValue, err := h.EventRepository.UpdateEvent(ctx, updateWithKey, key)
		if err != nil {
			return nil, err
		}
		event.HeaderImageS3Key = updateKeyValue.HeaderImageS3Key
		if errr != nil {
			return nil, errr
		}

		return url, nil

	}

	return nil, nil

}

package event

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/s3_client"
)

func (h *Handler) CreateEvent(ctx context.Context, input *models.CreateEventInput, updateBody *models.UpdateEventBody, imageData *[]byte, s3Client s3_client.S3Interface) (*models.Event, error) {
	var key *string
	var url *string
	//var wg sync.WaitGroup

	initInput := h.CreateTranslateStruct(ctx, input, nil, nil)
	event, err := h.EventRepository.CreateEvent(ctx, initInput, key)
	if err != nil {
		return nil, err
	}

	updateInput := &models.UpdateEventInput{
		ID:   event.ID,
		Body: *updateBody,
	}

	translationResp, err := h.CallTranslateAPI(ctx, &event.Title, &event.Description)
	if err != nil {
		return nil, err
	}
	translationsReinsertion := h.UpdateTranslateStruct(ctx, updateInput, translationResp.Title_TH, translationResp.Description_TH)
	_, err = h.EventRepository.UpdateEvent(ctx, translationsReinsertion, key)
	if err != nil {
		return nil, err
	}

	url, key, err = h.CreateEventS3Helper(ctx, s3Client, event, updateInput, imageData)
	if err != nil {
		return nil, err
	}
	_, err = h.EventRepository.UpdateEvent(ctx, translationsReinsertion, key)

	if err != nil {
		return nil, err
	}
	event.PresignedURL = url
	event.HeaderImageS3Key = key
	return event, nil
}

// helper for uploading image to s3
func (h *Handler) CreateEventS3Helper(ctx context.Context, s3Client s3_client.S3Interface, event *models.Event,
	updateInput *models.UpdateEventInput, imageData *[]byte) (*string, *string, error) {

	if imageData != nil {

		key, err := h.generateS3Key(event.ID)
		if err != nil {
			return nil, nil, err
		}

		url, errr := s3Client.UploadImage(ctx, key, *imageData)
		if errr != nil {
			return nil, nil, errr
		}

		return url, key, nil

	}

	return nil, nil, nil

}

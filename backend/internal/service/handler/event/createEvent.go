package event

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/s3_client"
)

func (h *Handler) CreateEvent(ctx context.Context, input *models.CreateEventInput, updateBody *models.UpdateEventBody, image_data *[]byte, s3Client s3_client.S3Interface) (*models.Event, error) {
	var key *string
	var url *string

	event, err := h.EventRepository.CreateEvent(ctx, input, key)
	if err != nil {
		return nil, err
	}

	if image_data != nil {

		url, err = h.CreateEventS3Helper(ctx, s3Client, event, updateBody, image_data)
		if err != nil {
			return nil, err
		}
	}

	event.PresignedURL = url

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

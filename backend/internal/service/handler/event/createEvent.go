package event

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/s3_client"
)

// TODO -> make helper
func (h *Handler) CreateEvent(ctx context.Context, input *models.CreateEventInput, image_data *[]byte, s3Client *s3_client.Client) (*models.Event, *string, *errs.HTTPError) {
	var key *string
	var url *string

	event, err := h.EventRepository.CreateEvent(ctx, input, key)
	if err != nil {
		return nil, nil, err.(*errs.HTTPError)
	}

	if image_data != nil {

		key, error := h.generateS3Key(event.ID)
		if error != nil {
			return nil, nil, error.(*errs.HTTPError)
		}
		uploadedUrl, errr := s3Client.UploadImage(ctx, key, *image_data)
		if errr != nil {
			return nil, nil, errr.(*errs.HTTPError)
		}
		url = uploadedUrl
	}

	return event, url, nil
}

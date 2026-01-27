package event

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/s3_client"
)

func (h *Handler) CreateEvent(ctx context.Context, input *models.CreateEventInput, image_data []byte, s3Client *s3_client.Client) (*models.Event, *string, *errs.HTTPError) {
	key, error := h.generateS3Key(input.Body.OrganizationID)
	var url *string

	if error != nil {
		return nil, nil, error.(*errs.HTTPError)
	}

	if image_data != nil {
		uploadedUrl, errr := s3Client.UploadImage(ctx, key, image_data)
		if errr != nil {
			return nil, nil, errr.(*errs.HTTPError)
		}
		url = &uploadedUrl
	}

	event, err := h.EventRepository.CreateEvent(ctx, input, &key)
	if err != nil {
		return nil, nil, err.(*errs.HTTPError)
	}

	return event, url, nil
}

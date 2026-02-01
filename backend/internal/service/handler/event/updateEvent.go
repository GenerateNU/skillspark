package event

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/s3_client"
)

// TODO -> make helper
func (h *Handler) UpdateEvent(ctx context.Context, input *models.UpdateEventInput, image_data *[]byte, s3Client *s3_client.Client) (*models.Event, *string, *errs.HTTPError) {

	var key *string
	var url *string

	occurences, err := h.EventRepository.GetEventOccurrencesByEventID(ctx, input.ID)
	if err != nil || len(occurences) == 0 {
		return nil, nil, err.(*errs.HTTPError)
	}

	key = occurences[0].Event.HeaderImageS3Key

	if image_data != nil {

		if key == nil {
			key, err = h.generateS3Key(input.ID)
			if err != nil {
				return nil, nil, err.(*errs.HTTPError)
			}
		}

		uploadedUrl, errr := s3Client.UploadImage(ctx, key, *image_data)
		if errr != nil {
			return nil, nil, errr.(*errs.HTTPError)
		}
		url = uploadedUrl

	}

	event, err := h.EventRepository.UpdateEvent(ctx, input, key)
	if err != nil {
		return nil, nil, err.(*errs.HTTPError)
	}

	return event, url, nil
}

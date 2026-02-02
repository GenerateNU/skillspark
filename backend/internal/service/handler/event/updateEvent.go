package event

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/s3_client"
)

func (h *Handler) UpdateEvent(ctx context.Context, input *models.UpdateEventInput, image_data *[]byte, s3Client *s3_client.Client) (*models.Event, error) {

	var key *string
	var url *string

	occurrences, err := h.EventRepository.GetEventOccurrencesByEventID(ctx, input.ID)
	if err != nil || len(occurrences) == 0 {
		return nil, err
	}

	key = occurrences[0].Event.HeaderImageS3Key

	if image_data != nil {

		if key == nil {
			key, err = h.generateS3Key(input.ID)
			if err != nil {
				return nil, err
			}
		}

		uploadedUrl, errr := s3Client.UploadImage(ctx, key, *image_data)
		if errr != nil {
			return nil, errr
		}
		url = uploadedUrl

	}

	event, err := h.EventRepository.UpdateEvent(ctx, input, key)
	event.PresignedURL = url

	if err != nil {
		return nil, err
	}

	return event, nil
}

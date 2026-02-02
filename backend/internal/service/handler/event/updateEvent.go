package event

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/s3_client"
)

// TODO -> make helper
func (h *Handler) UpdateEvent(ctx context.Context, input *models.UpdateEventInput, image_data *[]byte, s3Client *s3_client.Client) (*models.Event, *string, error) {

	var key *string
	var url *string
	var return_string string = "no occurences found"

	occurences, err := h.EventRepository.GetEventOccurrencesByEventID(ctx, input.ID)
	if err != nil || len(occurences) == 0 {
		return nil, &return_string, err
	}

	key = occurences[0].Event.HeaderImageS3Key

	if image_data != nil {

		if key == nil {
			key, err = h.generateS3Key(input.ID)
			if err != nil {
				return nil, nil, err
			}
		}

		uploadedUrl, errr := s3Client.UploadImage(ctx, key, *image_data)
		if errr != nil {
			return nil, nil, errr
		}
		url = uploadedUrl

	}

	event, err := h.EventRepository.UpdateEvent(ctx, input, key)
	if err != nil {
		return nil, nil, err
	}

	return event, url, nil
}

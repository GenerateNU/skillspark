package event

import (
	"context"
	"fmt"
	"skillspark/internal/models"
	"skillspark/internal/s3_client"
)

// TODO -> make helper
func (h *Handler) CreateEvent(ctx context.Context, input *models.CreateEventInput, image_data *[]byte, s3Client *s3_client.Client) (*models.Event, *string, error) {
	var key *string
	var url *string

	event, err := h.EventRepository.CreateEvent(ctx, input, key)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	if image_data != nil {

		key, err := h.generateS3Key(event.ID)
		if err != nil {
			fmt.Println("problem", err)
			return nil, nil, err
		}
		uploadedUrl, errr := s3Client.UploadImage(ctx, key, *image_data)
		if errr != nil {
			fmt.Println("also a problem", errr)
			return nil, nil, errr
		}
		url = uploadedUrl
	}

	return event, url, nil
}

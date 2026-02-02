package event

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/s3_client"
	"time"

	"github.com/google/uuid"
)

func (h *Handler) GetEventOccurrencesByEventID(ctx context.Context, input *models.GetEventOccurrencesByEventIDInput, s3Client *s3_client.Client) ([]models.EventOccurrence, *string, error) {
	id, parse_err := uuid.Parse(input.ID.String())
	if parse_err != nil {
		return nil, nil, errs.BadRequest("Invalid ID format")
	}

	eventOccurrence, err := h.EventRepository.GetEventOccurrencesByEventID(ctx, id)
	if err != nil {
		return nil, nil, err
	}

	if len(eventOccurrence) == 0 {
		return eventOccurrence, nil, nil
	}

	var url *string
	key := eventOccurrence[0].Event.HeaderImageS3Key
	if key != nil {
		presignedURL, err := s3Client.GeneratePresignedURL(ctx, *key, time.Hour)
		if err != nil {
			return nil, nil, err
		}

		url = &presignedURL
	}

	return eventOccurrence, url, nil
}

package eventoccurrence

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"time"

	"github.com/google/uuid"
)

func (h *Handler) GetEventOccurrenceByID(ctx context.Context, input *models.GetEventOccurrenceByIDInput) (*models.EventOccurrence, error) {

	if input.AcceptLanguage != "en-US" && input.AcceptLanguage != "th-TH" {
		e := errs.BadRequest("Invalid AcceptLanguage parameter: language does not exist")
		return nil, &e
	}

	id, parse_err := uuid.Parse(input.ID.String())
	if parse_err != nil {
		return nil, errs.BadRequest("Invalid ID format")
	}

	eventOccurrence, err := h.EventOccurrenceRepository.GetEventOccurrenceByID(ctx, id, input.AcceptLanguage)
	if err != nil {
		return nil, err
	}

	var url string
	var errr error

	key := eventOccurrence.Event.HeaderImageS3Key
	if key != nil {
		url, errr = h.s3Client.GeneratePresignedURL(ctx, *key, time.Hour)
		if errr != nil {
			return nil, err
		}
	}
	eventOccurrence.Event.PresignedURL = &url

	return eventOccurrence, nil
}

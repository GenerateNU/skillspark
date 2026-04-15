package organization

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"time"

	"github.com/google/uuid"
)

func (h *Handler) GetEventOccurrencesByOrganizationID(ctx context.Context, input *models.GetEventOccurrencesByOrganizationIDInput) ([]models.EventOccurrence, error) {

	id, parse_err := uuid.Parse(input.ID.String())
	if parse_err != nil {
		return nil, errs.BadRequest("Invalid ID format")
	}

	eventOccurrences, err := h.OrganizationRepository.GetEventOccurrencesByOrganizationID(ctx, id, input.AcceptLanguage)
	if err != nil {
		return nil, err
	}

	for i := range eventOccurrences {
		key := eventOccurrences[i].Event.HeaderImageS3Key
		if key != nil {
			url, errr := h.s3client.GeneratePresignedURL(ctx, *key, time.Hour)
			if errr != nil {
				return nil, errr
			}
			eventOccurrences[i].Event.PresignedURL = &url
		}
	}

	return eventOccurrences, nil
}

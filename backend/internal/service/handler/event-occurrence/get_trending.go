package eventoccurrence

import (
	"context"
	"skillspark/internal/models"
)

func (h *Handler) GetTrendingEventOccurrences(ctx context.Context, input *models.GetTrendingEventOccurrencesInput) ([]models.EventOccurrence, error) {

	eventOccurrence, err := h.EventOccurrenceRepository.GetTrendingEventOccurrences(ctx, input)
	if err != nil {
		return nil, err
	}

	for idx := range eventOccurrence {
		err = h.AssignURLS(ctx, eventOccurrence, idx)
		if err != nil {
			return nil, err
		}
	}

	return eventOccurrence, nil
}

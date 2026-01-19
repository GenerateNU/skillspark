package eventoccurrence

import (
	"context"
	"skillspark/internal/models"
)

func (h *Handler) CreateEventOccurrence(ctx context.Context, input *models.CreateEventOccurrenceInput) (*models.EventOccurrence, error) {
	eventOccurrence, err := h.EventOccurrenceRepository.CreateEventOccurrence(ctx, input)
	if err != nil {
		return nil, err
	}
	return eventOccurrence, nil
}
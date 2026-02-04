package eventoccurrence

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/utils"
)

func (h *Handler) GetAllEventOccurrences(ctx context.Context, pagination utils.Pagination) ([]models.EventOccurrence, error) {
	eventOccurrence, err := h.EventOccurrenceRepository.GetAllEventOccurrences(ctx, pagination)
	if err != nil {
		return nil, err
	}
	return eventOccurrence, nil
}

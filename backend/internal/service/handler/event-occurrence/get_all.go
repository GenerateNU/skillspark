package eventoccurrence

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/utils"
)

func (h *Handler) GetAllEventOccurrences(ctx context.Context, pagination utils.Pagination, AcceptLanguage string) ([]models.EventOccurrence, error) {
	eventOccurrence, err := h.EventOccurrenceRepository.GetAllEventOccurrences(ctx, pagination, AcceptLanguage)
	if err != nil {
		return nil, err
	}
	return eventOccurrence, nil
}

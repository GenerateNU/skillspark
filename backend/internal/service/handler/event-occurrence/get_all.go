package eventoccurrence

import (
	"context"
	"fmt"
	"log/slog"
	"skillspark/internal/models"
	"skillspark/internal/utils"
)

func (h *Handler) GetAllEventOccurrences(ctx context.Context, pagination utils.Pagination, filters models.GetAllEventOccurrencesFilter) ([]models.EventOccurrence, error) {
	fmt.Printf("Getting all event occurrences with filters: %v\n", filters)
	eventOccurrence, err := h.EventOccurrenceRepository.GetAllEventOccurrences(ctx, pagination, filters)
	if err != nil {
		return nil, err
	}
	slog.Info("Retrieved event occurrences:", "eventOccurrence", eventOccurrence)
	return eventOccurrence, nil
}

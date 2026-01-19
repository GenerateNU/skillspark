package event

import (
	"context"
	"skillspark/internal/models"
)

func (h *Handler) UpdateEvent(ctx context.Context, input *models.UpdateEventInput) (*models.Event, error) {
	event, err := h.EventRepository.UpdateEvent(ctx, input)
	if err != nil {
		return nil, err
	}

	return event, nil
}

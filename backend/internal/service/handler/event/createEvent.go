package event

import (
	"context"
	"skillspark/internal/models"
)

func (h *Handler) CreateEvent(ctx context.Context, input *models.CreateEventInput) (*models.Event, error) {
	event, err := h.EventRepository.CreateEvent(ctx, input)
	if err != nil {
		return nil, err
	}

	return event, nil
}

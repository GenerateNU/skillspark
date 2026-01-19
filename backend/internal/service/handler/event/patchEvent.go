package event

import (
	"context"
	"skillspark/internal/models"
)

func (h *Handler) PatchEvent(ctx context.Context, input *models.PatchEventInput) (*models.Event, error) {
	event, err := h.EventRepository.PatchEvent(ctx, input)
	if err != nil {
		return nil, err
	}

	return event, nil
}

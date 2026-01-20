package event

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
)

func (h *Handler) UpdateEvent(ctx context.Context, input *models.UpdateEventInput) (*models.Event, *errs.HTTPError) {
	event, err := h.EventRepository.UpdateEvent(ctx, input)
	if err != nil {
		return nil, err.(*errs.HTTPError)
	}

	return event, nil
}

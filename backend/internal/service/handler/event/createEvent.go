package event

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
)

func (h *Handler) CreateEvent(ctx context.Context, input *models.CreateEventInput) (*models.Event, *errs.HTTPError) {
	event, err := h.EventRepository.CreateEvent(ctx, input)
	if err != nil {
		return nil, err.(*errs.HTTPError)
	}

	return event, nil
}

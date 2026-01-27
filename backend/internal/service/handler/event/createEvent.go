package event

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
)

func (h *Handler) CreateEvent(ctx context.Context, input *models.CreateEventInput) (*models.Event, *errs.HTTPError) {

	// handler generates s3 key
	// handler calls on s3client method
	// handles passes key into

	event, err := h.EventRepository.CreateEvent(ctx, input)
	if err != nil {
		return nil, err.(*errs.HTTPError)
	}

	return event, nil
}

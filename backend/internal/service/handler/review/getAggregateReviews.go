package review

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

func (h *Handler) GetAggregateReviews(ctx context.Context, id uuid.UUID) (*models.ReviewAggregate, error) {

	// the language here does not matter, we are just checking that the event exists.
	if _, err := h.EventRepository.GetEventByID(ctx, id, "en-US"); err != nil {
		return nil, errs.BadRequest("Invalid event_id: event does not exist" + err.Error())
	}

	aggregate, httpErr := h.ReviewRepository.GetAggregateReviews(ctx, id)
	if httpErr != nil {
		return nil, httpErr
	}

	return aggregate, nil
}

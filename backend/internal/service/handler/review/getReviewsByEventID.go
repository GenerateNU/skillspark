package review

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/utils"

	"github.com/google/uuid"
)

func (h *Handler) GetReviewsByEventID(ctx context.Context, id uuid.UUID, pagination utils.Pagination) ([]models.Review, error) {

	if _, err := h.EventRepository.GetEventByID(ctx, id); err != nil {
		return nil, errs.BadRequest("Invalid event_id: event does not exist")
	}

	reviews, httpErr := h.ReviewRepository.GetReviewsByEventID(ctx, id, pagination)
	if httpErr != nil {
		return nil, httpErr
	}

	return reviews, nil
}

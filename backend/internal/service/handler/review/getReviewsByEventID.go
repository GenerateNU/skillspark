package review

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/utils"

	"github.com/google/uuid"
)

func (h *Handler) GetReviewsByEventID(ctx context.Context, id uuid.UUID, AcceptLanguage string, pagination utils.Pagination) ([]models.Review, error) {

	if AcceptLanguage != "en-US" && AcceptLanguage != "th-TH" {
		e := errs.BadRequest("Invalid AcceptLanguage parameter: language does not exist")
		return nil, &e
	}

	if _, err := h.EventRepository.GetEventByID(ctx, id, AcceptLanguage); err != nil {
		return nil, errs.BadRequest("Invalid event_id: event does not exist" + err.Error())
	}

	reviews, httpErr := h.ReviewRepository.GetReviewsByEventID(ctx, id, AcceptLanguage, pagination)
	if httpErr != nil {
		return nil, httpErr
	}

	return reviews, nil
}

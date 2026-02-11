package review

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/utils"

	"github.com/google/uuid"
)

func (h *Handler) GetReviewsByGuardianID(ctx context.Context, id uuid.UUID, pagination utils.Pagination) ([]models.Review, error) {

	if _, err := h.GuardianRepository.GetGuardianByID(ctx, id); err != nil {
		return nil, errs.BadRequest("Invalid guardian_id: guardian does not exist")
	}

	reviews, httpErr := h.ReviewRepository.GetReviewsByGuardianID(ctx, id, pagination)
	if httpErr != nil {
		return nil, httpErr
	}

	return reviews, nil
}

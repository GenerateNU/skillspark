package review

import (
	"context"
	"skillspark/internal/errs"

	"github.com/google/uuid"
)

func (h *Handler) DeleteReview(ctx context.Context, id uuid.UUID) (string, *errs.HTTPError) {
	err := h.ReviewRepository.DeleteReview(ctx, id)
	if err != nil {
		return "", err.(*errs.HTTPError)
	}

	return "Review successfully deleted.", nil
}

package event

import (
	"context"
	"skillspark/internal/errs"

	"github.com/google/uuid"
)

func (h *Handler) DeleteEvent(ctx context.Context, id uuid.UUID) (string, *errs.HTTPError) {
	err := h.EventRepository.DeleteEvent(ctx, id)
	if err != nil {
		return "", err.(*errs.HTTPError)
	}

	return "Event successfully deleted.", nil
}

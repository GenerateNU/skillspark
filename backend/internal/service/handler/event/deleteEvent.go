package event

import (
	"context"
	"github.com/google/uuid"
	"skillspark/internal/errs"
)

func (h *Handler) DeleteEvent(ctx context.Context, id uuid.UUID) *errs.HTTPError {
	err := h.EventRepository.DeleteEvent(ctx, id)
	if err != nil {
		return err.(*errs.HTTPError)
	}

	return nil
}

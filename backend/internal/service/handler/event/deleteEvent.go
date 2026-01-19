package event

import (
	"context"
	"github.com/google/uuid"
)

func (h *Handler) DeleteEvent(ctx context.Context, id uuid.UUID) error {
	_, err := h.EventRepository.DeleteEvent(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

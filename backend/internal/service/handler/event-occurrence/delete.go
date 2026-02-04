package eventoccurrence

import (
	"context"
	"skillspark/internal/errs"

	"github.com/google/uuid"
)

func (h *Handler) CancelEventOccurrence(ctx context.Context, id uuid.UUID) (string, *errs.HTTPError) {
	err := h.EventOccurrenceRepository.CancelEventOccurrence(ctx, id)
	if err != nil {
		return "", err.(*errs.HTTPError)
	}

	return "Event successfully deleted.", nil
}

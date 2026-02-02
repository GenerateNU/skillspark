package eventoccurrence

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

func (h *Handler) GetEventOccurrenceByID(ctx context.Context, input *models.GetEventOccurrenceByIDInput) (*models.EventOccurrence, error) {
	id, parse_err := uuid.Parse(input.ID.String())
	if parse_err != nil {
		return nil, errs.BadRequest("Invalid ID format")
	}

	eventOccurrence, err := h.EventOccurrenceRepository.GetEventOccurrenceByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return eventOccurrence, nil
}

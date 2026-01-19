package eventoccurrence

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

func (h *Handler) GetEventOccurrenceByID(ctx context.Context, input *models.GetEventOccurrenceByIDInput) (*models.EventOccurrence, error) {
	id, err := uuid.Parse(input.ID.String())
	if err != nil {
		return nil, errs.BadRequest("Invalid ID format")
	}
	
	eventOccurrence, errr := h.EventOccurrenceRepository.GetEventOccurrenceByID(ctx, id)
	if errr != nil {
		return nil, errr
	}
	return eventOccurrence, nil
}
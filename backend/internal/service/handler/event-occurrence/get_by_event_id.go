package eventoccurrence

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

func (h *Handler) GetEventOccurrencesByEventID(ctx context.Context, input *models.GetEventOccurrencesByEventIDInput) ([]models.EventOccurrence, error) {
	id, err := uuid.Parse(input.ID.String())
	if err != nil {
		return nil, errs.BadRequest("Invalid ID format")
	}
	
	eventOccurrence, errr := h.EventOccurrenceRepository.GetEventOccurrencesByEventID(ctx, id)
	if errr != nil {
		return nil, errr
	}
	return eventOccurrence, nil
}
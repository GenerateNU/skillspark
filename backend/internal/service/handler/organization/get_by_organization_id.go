package organization

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

func (h *Handler) GetEventOccurrencesByOrganizationID(ctx context.Context, input *models.GetEventOccurrencesByOrganizationIDInput) ([]models.EventOccurrence, error) {
	id, parse_err := uuid.Parse(input.ID.String())
	if parse_err != nil {
		return nil, errs.BadRequest("Invalid ID format")
	}
	
	eventOccurrence, err := h.OrganizationRepository.GetEventOccurrencesByOrganizationID(ctx, id)
	if err != nil {
		return nil, err
	}
	return eventOccurrence, nil
}
package registration

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

func (h *Handler) GetRegistrationsByEventOccurrenceID(ctx context.Context, input *models.GetRegistrationsByEventOccurrenceIDInput) (*models.GetRegistrationsByEventOccurrenceIDOutput, error) {
	eventOccurrenceID, err := uuid.Parse(input.EventOccurrenceID.String())
	if err != nil {
		return nil, errs.BadRequest("Invalid child ID format")
	}

	registrations, httpErr := h.RegistrationRepository.GetRegistrationsByEventOccurrenceID(ctx, &models.GetRegistrationsByEventOccurrenceIDInput{EventOccurrenceID: eventOccurrenceID})
	if httpErr != nil {
		return nil, httpErr
	}

	return registrations, nil
}

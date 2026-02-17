package registration

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
)

func (h *Handler) CreateRegistration(ctx context.Context, input *models.CreateRegistrationInput) (*models.CreateRegistrationOutput, error) {
	_, err := h.EventOccurrenceRepository.GetEventOccurrenceByID(ctx, input.Body.EventOccurrenceID)
	if err != nil {
		return nil, errs.BadRequest("Invalid event_occurrence_id: event occurrence does not exist")
	}

	if _, err := h.ChildRepository.GetChildByID(ctx, input.Body.ChildID); err != nil {
		return nil, errs.BadRequest("Invalid child_id: child does not exist")
	}

	if _, err := h.GuardianRepository.GetGuardianByID(ctx, input.Body.GuardianID); err != nil {
		return nil, errs.BadRequest("Invalid guardian_id: guardian does not exist")
	}

	registration, err := h.RegistrationRepository.CreateRegistration(ctx, input)
	if err != nil {
		return nil, err
	}

	return registration, nil
}

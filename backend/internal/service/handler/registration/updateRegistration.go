package registration

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
)

func (h *Handler) UpdateRegistration(ctx context.Context, input *models.UpdateRegistrationInput) (*models.UpdateRegistrationOutput, error) {
	if input.Body.ChildID != nil {
		if _, err := h.ChildRepository.GetChildByID(ctx, *input.Body.ChildID); err != nil {
			return nil, errs.BadRequest("Invalid child_id: child does not exist")
		}
	}

	if input.Body.GuardianID != nil {
		if _, err := h.GuardianRepository.GetGuardianByID(ctx, *input.Body.GuardianID); err != nil {
			return nil, errs.BadRequest("Invalid guardian_id: guardian does not exist")
		}
	}

	if input.Body.EventOccurrenceID != nil {
		if _, err := h.EventOccurrenceRepository.GetEventOccurrenceByID(ctx, *input.Body.EventOccurrenceID); err != nil {
			return nil, errs.BadRequest("Invalid event_occurrence_id: event occurrence does not exist")
		}
	}

	updated, err := h.RegistrationRepository.UpdateRegistration(ctx, input)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
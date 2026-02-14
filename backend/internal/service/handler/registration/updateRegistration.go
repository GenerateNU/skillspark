package registration

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
)

func (h *Handler) UpdateRegistration(ctx context.Context, input *models.UpdateRegistrationInput) (*models.UpdateRegistrationOutput, error) {
	if _, err := h.ChildRepository.GetChildByID(ctx, input.Body.ChildID); err != nil {
		return nil, errs.BadRequest("Invalid child_id: child does not exist")
	}

	updated, err := h.RegistrationRepository.UpdateRegistration(ctx, input)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

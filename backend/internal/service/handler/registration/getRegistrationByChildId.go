package registration

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

func (h *Handler) GetRegistrationsByChildID(ctx context.Context, input *models.GetRegistrationsByChildIDInput) (*models.GetRegistrationsByChildIDOutput, error) {
	childID, err := uuid.Parse(input.ChildID.String())
	if err != nil {
		return nil, errs.BadRequest("Invalid child ID format")
	}

	registrations, httpErr := h.RegistrationRepository.GetRegistrationsByChildID(ctx, &models.GetRegistrationsByChildIDInput{ChildID: childID})
	if httpErr != nil {
		return nil, httpErr
	}

	return registrations, nil
}

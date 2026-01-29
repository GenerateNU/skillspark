package registration

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

func (h *Handler) GetRegistrationsByGuardianID(ctx context.Context, input *models.GetRegistrationsByGuardianIDInput) (*models.GetRegistrationsByGuardianIDOutput, error) {
	guardianID, err := uuid.Parse(input.GuardianID.String())
	if err != nil {
		return nil, errs.BadRequest("Invalid guardian ID format")
	}

	registrations, httpErr := h.RegistrationRepository.GetRegistrationsByGuardianID(ctx, &models.GetRegistrationsByGuardianIDInput{GuardianID: guardianID})
	if httpErr != nil {
		return nil, httpErr
	}

	return registrations, nil
}
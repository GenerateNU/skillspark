package guardian

import (
	"context"

	"skillspark/internal/models"
)

func (h *Handler) DeleteGuardian(ctx context.Context, input *models.DeleteGuardianInput) (*models.Guardian, error) {
	guardian, err := h.GuardianRepository.DeleteGuardian(ctx, input.ID)

	// TODO: Verify that no child has this guardian as a parent 

	if err != nil {
		return nil, err
	}

	return guardian, nil
}
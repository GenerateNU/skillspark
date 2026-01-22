package guardian

import (
	"context"

	"skillspark/internal/models"
)

func (h *Handler) UpdateGuardian(ctx context.Context, input *models.UpdateGuardianInput) (*models.Guardian, error) {
	guardian, err := h.GuardianRepository.UpdateGuardian(ctx, input)
	if err != nil {
		return nil, err
	}

	return guardian, nil
}

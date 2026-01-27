package guardian

import (
	"context"
	"skillspark/internal/models"
)

func (h *Handler) CreateGuardian(ctx context.Context, input *models.CreateGuardianInput) (*models.Guardian, error) {
	guardian, err := h.GuardianRepository.CreateGuardian(ctx, input)
	if err != nil {
		return nil, err
	}

	return guardian, nil
}

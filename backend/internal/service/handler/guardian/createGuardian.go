package guardian

import (
	"context"
	"skillspark/internal/models"
)

func (h *Handler) CreateGuardian(ctx context.Context, input *models.CreateGuardianInput) (*models.Guardian, error) {
	// Check if user is already a manager
	_, err := h.ManagerRepository.GetManagerByUserID(ctx, input.Body.UserID)
	if err == nil {
		return nil, err
	}

	guardian, err := h.GuardianRepository.CreateGuardian(ctx, input)
	if err != nil {
		return nil, err
	}

	return guardian, nil
}

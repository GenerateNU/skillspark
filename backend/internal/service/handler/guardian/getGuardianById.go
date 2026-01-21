package guardian

import (
	"context"

	"skillspark/internal/models"
)

func (h *Handler) GetGuardianById(ctx context.Context, input *models.GetGuardianByIDInput) (*models.Guardian, error) {
	guardian, err := h.GuardianRepository.GetGuardianByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	return guardian, nil
}

package guardian

import (
	"context"

	"skillspark/internal/models"
)

func (h *Handler) GetGuardianById(ctx context.Context, input *models.GetGuardianByIDInput) (*models.Guardian, error) {
	guardian, err := h.GuardianRepository.GetGuardianByID(ctx, input.ID)
	print("At the handler level")
	if err != nil {
		return nil, err
	}

	return guardian, nil
}

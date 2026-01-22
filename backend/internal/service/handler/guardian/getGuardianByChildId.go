package guardian 

import (
	"context"

	"skillspark/internal/models"
)

func (h *Handler) GetGuardianByChildId(ctx context.Context, input *models.GetGuardianByChildIDInput) (*models.Guardian, error) {
	guardian, err := h.GuardianRepository.GetGuardianByChildID(ctx, input.ChildID)
	if err != nil {
		return nil, err
	}

	return guardian, nil
}

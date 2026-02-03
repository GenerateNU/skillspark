package guardian

import (
	"context"

	"skillspark/internal/models"
	//"skillspark/internal/auth"
)

func (h *Handler) DeleteGuardian(ctx context.Context, input *models.DeleteGuardianInput) (*models.Guardian, error) {
	guardian, err := h.GuardianRepository.DeleteGuardian(ctx, input.ID)

	if err != nil {
		return nil, err
	}

	// delete Supabase auth user
	// deleteErr := auth.SupabaseDeleteUser(&h.config, guardian.ID)
	// if deleteErr != nil {
	// 	return nil, deleteErr
	// }

	return guardian, nil
}

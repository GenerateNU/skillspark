package manager

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/auth"

	"github.com/google/uuid"
)

// DeleteManager handles DELETE /manager
func (h *Handler) DeleteManager(ctx context.Context, input *models.DeleteManagerInput) (*models.Manager, error) {
	// Input is already parsed and validated by Huma!
	// Just pass it to the repository
	id, _ := uuid.Parse(input.ID.String())
	manager, err := h.ManagerRepository.DeleteManager(ctx, id)
	if err != nil {
		return nil, err
	}

	// delete Supabase auth user
	deleteErr := auth.SupabaseDeleteUser(&h.config, manager.ID)
	if deleteErr != nil {
		return nil, deleteErr
	}

	return manager, nil
}

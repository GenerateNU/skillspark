package manager

import (
	"context"
	"skillspark/internal/models"
)

// CreateManager handles POST /manager
func (h *Handler) PatchManager(ctx context.Context, input *models.PatchManagerInput) (*models.Manager, error) {
	// Input is already parsed and validated by Huma!
	// Just pass it to the repository
	manager, err := h.ManagerRepository.PatchManager(ctx, input)
	if err != nil {
		return nil, err
	}

	return manager, nil
}

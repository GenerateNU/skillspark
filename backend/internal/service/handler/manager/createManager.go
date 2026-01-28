package manager

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
)

// CreateManager handles POST /manager
func (h *Handler) CreateManager(ctx context.Context, input *models.CreateManagerInput) (*models.Manager, error) {
	// Check if user is already a guardian
	_, err := h.GuardianRepository.GetGuardianByUserID(ctx, input.Body.UserID)
	if err == nil {
		return nil, errs.Conflict("Guardian", "user_id", input.Body.UserID)
	}

	// Input is already parsed and validated by Huma!
	// Just pass it to the repository
	manager, err := h.ManagerRepository.CreateManager(ctx, input)
	if err != nil {
		return nil, err
	}

	return manager, nil
}

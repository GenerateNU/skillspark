package manager

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

// GetManagerById handles GET /managers/:id
func (h *Handler) GetManagerByID(ctx context.Context, input *models.GetManagerByIDInput) (*models.Manager, error) {
	id, err := uuid.Parse(input.ID.String())
	if err != nil {
		return nil, errs.BadRequest("Invalid ID format")
	}

	manager, httpErr := h.ManagerRepository.GetManagerByID(ctx, id)
	if httpErr != nil {
		return nil, httpErr
	}
	return manager, nil
}

package child

import (
	"context"
	"skillspark/internal/models"
)

func (h *Handler) UpdateChildByID(ctx context.Context, input *models.UpdateChildInput) (*models.Child, error) {
	child, httpErr := h.ChildRepository.UpdateChildByID(ctx, input)
	if httpErr != nil {
		return nil, httpErr
	}
	return child, nil
}

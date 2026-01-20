package child

import (
	"context"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

func (h *Handler) UpdateChildByID(ctx context.Context, childID uuid.UUID, input *models.UpdateChildInput) (*models.Child, error) {
	child, httpErr := h.ChildRepository.UpdateChildByID(ctx, childID, input)
	if httpErr != nil {
		return nil, httpErr
	}
	return child, nil
}

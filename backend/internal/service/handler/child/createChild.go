package child

import (
	"context"
	"skillspark/internal/models"
)

func (h *Handler) CreateChild(ctx context.Context, input *models.CreateChildInput) (*models.Child, error) {
	child, err := h.ChildRepository.CreateChild(ctx, input)
	if err != nil {
		return nil, err
	}

	return child, nil
}

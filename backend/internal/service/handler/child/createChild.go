package child

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
)

func (h *Handler) CreateChild(ctx context.Context, input *models.CreateChildInput) (*models.Child, *errs.HTTPError) {
	child, err := h.ChildRepository.CreateChild(ctx, input)
	if err != nil {
		return nil, err.(*errs.HTTPError)
	}

	return child, nil
}

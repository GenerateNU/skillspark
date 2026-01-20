package child

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

func (h *Handler) UpdateChildByID(ctx context.Context, childID uuid.UUID, input *models.UpdateChildInput) (*models.Child, *errs.HTTPError) {
	child, httpErr := h.ChildRepository.UpdateChildByID(ctx, childID, input)
	if httpErr != nil {
		return nil, httpErr.(*errs.HTTPError)
	}
	return child, nil
}

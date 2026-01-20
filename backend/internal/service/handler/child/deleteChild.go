package child

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

func (h *Handler) DeleteChildByID(ctx context.Context, input *models.ChildIDInput) (*models.Child, error) {
	id, err := uuid.Parse(input.ID.String())
	if err != nil {
		return nil, errs.BadRequest("Invalid ID format")
	}

	child, httpErr := h.ChildRepository.DeleteChildByID(ctx, id)
	if httpErr != nil {
		return nil, httpErr
	}
	return child, nil
}

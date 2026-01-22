package child

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

func (h *Handler) GetChildrenByParentID(ctx context.Context, input *models.GuardianIDInput) ([]models.Child, error) {
	id, err := uuid.Parse(input.ID.String())
	if err != nil {
		return nil, errs.BadRequest("Invalid ID format")
	}

	children, httpErr := h.ChildRepository.GetChildrenByParentID(ctx, id)
	if httpErr != nil {
		return nil, httpErr
	}
	return children, nil
}

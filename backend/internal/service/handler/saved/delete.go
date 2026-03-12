package saved

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

func (h *Handler) DeleteSaved(ctx context.Context, input *models.DeleteSavedInput) error {

	id, err := uuid.Parse(input.ID.String())
	if err != nil {
		return errs.BadRequest("Invalid ID format")
	}

	httpErr := h.SavedRepository.DeleteSaved(ctx, id)
	if httpErr != nil {
		return httpErr
	}
	return nil
}

package saved

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
)

func (h *Handler) DeleteSaved(ctx context.Context, input *models.DeleteSavedInput) (string, *errs.HTTPError) {

	httpErr := h.SavedRepository.DeleteSaved(ctx, input.ID)
	if httpErr != nil {
		return "", httpErr.(*errs.HTTPError)
	}
	return "Saved successfully deleted.", nil
}

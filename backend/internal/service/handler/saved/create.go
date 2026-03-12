package saved

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
)

func (h *Handler) CreateSaved(ctx context.Context, input *models.CreateSavedInput) (*models.Saved, *errs.HTTPError) {

	saved, err := h.SavedRepository.CreateSaved(ctx, input)
	if err != nil {
		return nil, err.(*errs.HTTPError)
	}

	return saved, nil
}

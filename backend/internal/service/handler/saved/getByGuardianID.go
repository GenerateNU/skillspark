package saved

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/utils"

	"github.com/google/uuid"
)

func (h *Handler) GetByGuardianID(ctx context.Context, id uuid.UUID, pagination utils.Pagination) ([]models.Saved, error) {

	if _, err := h.GuardianRepository.GetGuardianByID(ctx, id); err != nil {
		return nil, errs.BadRequest("Invalid guardian_id: guardian does not exist")
	}

	reviews, httpErr := h.SavedRepository.GetByGuardianID(ctx, id, pagination)
	if httpErr != nil {
		return nil, httpErr
	}

	return reviews, nil
}

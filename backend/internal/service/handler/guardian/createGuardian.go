package guardian

import (
	"context"
	"errors"

	"skillspark/internal/errs"
	"skillspark/internal/models"
)

func (h *Handler) CreateGuardian(ctx context.Context, input *models.CreateGuardianInput) (*models.Guardian, error) {

	// TODO verify no guardian exists for this user

	guardian, err := h.GuardianRepository.GetGuardianByUserID(ctx, input.Body.UserID)


	
	if err != nil && errors.Is(err, errs.NotFound("Guardian", "user_id", input.Body.UserID)) {
		return nil, err
	}

	if guardian != nil {
		return nil, errs.BadRequest("Guardian already exists for this user")
	}

	guardian, err = h.GuardianRepository.CreateGuardian(ctx, input)
	if err != nil {
		return nil, err
	}

	return guardian, nil
}
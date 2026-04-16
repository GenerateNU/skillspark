package user

import (
	"context"
	"skillspark/internal/models"
)

func (h *Handler) UsernameExists(ctx context.Context, input *models.UsernameExistsInput) (bool, error) {
	exists, err := h.UserRepository.UsernameExists(ctx, input.Username)
	if err != nil {
		return false, err
	}

	return exists, nil
}

package user

import (
	"context"
	"skillspark/internal/models"
)

func (h *Handler) GetUserByUsername(ctx context.Context, input *models.GetUserByUsernameInput) (bool, error) {
	exists, err := h.UserRepository.GetUserByUsername(ctx, input.Username)
	if err != nil {
		return false, err
	}

	return exists, nil
}

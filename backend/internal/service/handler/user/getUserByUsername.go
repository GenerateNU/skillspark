package user

import (
	"context"
	"skillspark/internal/models"
)

func (h *Handler) GetUserByUsername(ctx context.Context, input *models.GetUserByUsernameInput) (*models.User, error) {
	user, err := h.UserRepository.GetUserByUsername(ctx, input.Username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

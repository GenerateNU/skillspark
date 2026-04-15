package user

import (
	"skillspark/internal/storage"
)

type Handler struct {
	UserRepository storage.UserRepository
}

func NewHandler(userRepository storage.UserRepository) *Handler {
	return &Handler{
		UserRepository: userRepository,
	}
}

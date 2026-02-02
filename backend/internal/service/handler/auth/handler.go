package auth

import (
	"skillspark/internal/config"
	"skillspark/internal/storage"
)

type Handler struct {
	config             config.Supabase
	userRepository     storage.UserRepository
	guardianRepository storage.GuardianRepository
	managerRepository  storage.ManagerRepository
}

func NewHandler(config config.Supabase, userRepository storage.UserRepository, guardianRepository storage.GuardianRepository, managerRepository storage.ManagerRepository) *Handler {
	return &Handler{
		config,
		userRepository,
		guardianRepository,
		managerRepository,
	}
}

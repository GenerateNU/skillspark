package auth

import (
	"skillspark/internal/config"
	"skillspark/internal/storage"
)

type Handler struct {
	config              config.Supabase
	userRepository      storage.UserRepository
	guardianRepository  storage.GuardianRepository
	managerRepository   storage.ManagerRepository
}

type Credentials struct {
	Email      string  `json:"email"`
	Password   string  `json:"password"`
	FirstName  *string `json:"first_name"`
	LastName   *string `json:"last_name"`
	RememberMe bool
}

func NewHandler(config config.Supabase, userRepository storage.UserRepository, guardianRepository storage.GuardianRepository, managerRepository storage.ManagerRepository) *Handler {
	return &Handler{
		config,
		userRepository,
		guardianRepository,
		managerRepository,
	}
}
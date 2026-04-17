package auth

import (
	"skillspark/internal/config"
	"skillspark/internal/storage"
)

type Handler struct {
	config             config.Supabase
	appConfig          config.Application
	userRepository     storage.UserRepository
	guardianRepository storage.GuardianRepository
	managerRepository  storage.ManagerRepository
}

func NewHandler(supabaseCfg config.Supabase, appCfg config.Application, userRepository storage.UserRepository, guardianRepository storage.GuardianRepository, managerRepository storage.ManagerRepository) *Handler {
	return &Handler{
		config:             supabaseCfg,
		appConfig:          appCfg,
		userRepository:     userRepository,
		guardianRepository: guardianRepository,
		managerRepository:  managerRepository,
	}
}

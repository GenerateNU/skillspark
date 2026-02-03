package manager

import (
	"skillspark/internal/storage"
	"skillspark/internal/config"
)

type Handler struct {
	ManagerRepository  storage.ManagerRepository
	GuardianRepository storage.GuardianRepository
	config             config.Supabase
}

func NewHandler(managerRepository storage.ManagerRepository, guardianRepository storage.GuardianRepository, config config.Supabase) *Handler {
	return &Handler{
		ManagerRepository:  managerRepository,
		GuardianRepository: guardianRepository,
		config: config,
	}
}

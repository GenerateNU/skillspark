package guardian

import (
	"skillspark/internal/storage" 
	"skillspark/internal/config"
)

type Handler struct {
	GuardianRepository storage.GuardianRepository
	ManagerRepository  storage.ManagerRepository
	config             config.Supabase
}

func NewHandler(guardianRepository storage.GuardianRepository, managerRepository storage.ManagerRepository, config config.Supabase) *Handler {
	return &Handler{
		GuardianRepository: guardianRepository,
		ManagerRepository:  managerRepository,
		config: config,
	}
}

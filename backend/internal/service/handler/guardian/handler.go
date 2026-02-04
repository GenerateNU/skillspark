package guardian

import (
	"skillspark/internal/storage" 
	"skillspark/internal/config"
)

type Handler struct {
	GuardianRepository storage.GuardianRepository
	config             config.Supabase
}

func NewHandler(guardianRepository storage.GuardianRepository, config config.Supabase) *Handler {
	return &Handler{
		GuardianRepository: guardianRepository,
		config: config,
	}
}

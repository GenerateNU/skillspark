package manager

import (
	"skillspark/internal/storage"
)

type Handler struct {
	ManagerRepository  storage.ManagerRepository
	GuardianRepository storage.GuardianRepository
}

func NewHandler(managerRepository storage.ManagerRepository, guardianRepository storage.GuardianRepository) *Handler {
	return &Handler{
		ManagerRepository:  managerRepository,
		GuardianRepository: guardianRepository,
	}
}

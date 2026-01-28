package guardian

import "skillspark/internal/storage"

type Handler struct {
	GuardianRepository storage.GuardianRepository
	ManagerRepository  storage.ManagerRepository
}

func NewHandler(guardianRepository storage.GuardianRepository, managerRepository storage.ManagerRepository) *Handler {
	return &Handler{
		GuardianRepository: guardianRepository,
		ManagerRepository:  managerRepository,
	}
}

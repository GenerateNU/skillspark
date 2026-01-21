package guardian

import "skillspark/internal/storage"

type Handler struct {
	GuardianRepository storage.GuardianRepository
}

func NewHandler(guardianRepository storage.GuardianRepository) *Handler {
	return &Handler{
		GuardianRepository: guardianRepository,
	}
}

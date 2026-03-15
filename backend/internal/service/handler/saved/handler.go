package saved

import (
	"skillspark/internal/storage"
)

type Handler struct {
	SavedRepository    storage.SavedRepository
	GuardianRepository storage.GuardianRepository
}

func NewHandler(savedRepository storage.SavedRepository, guardianRepository storage.GuardianRepository) *Handler {
	return &Handler{
		SavedRepository:    savedRepository,
		GuardianRepository: guardianRepository,
	}
}

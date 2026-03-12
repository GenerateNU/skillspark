package saved

import (
	"skillspark/internal/storage"
)

type Handler struct {
	SavedRepository storage.SavedRepository
}

func NewHandler(savedRepository storage.SavedRepository) *Handler {
	return &Handler{
		SavedRepository: savedRepository,
	}
}

package school

import (
	"skillspark/internal/storage"
)

type Handler struct {
	SchoolRepository storage.SchoolRepository
}

func NewHandler(schoolRepository storage.SchoolRepository) *Handler {
	return &Handler{
		SchoolRepository: schoolRepository,
	}
}

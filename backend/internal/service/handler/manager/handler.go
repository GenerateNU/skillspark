package manager

import (
	"skillspark/internal/storage"
)

type Handler struct {
	ManagerRepository storage.ManagerRepository
}

func NewHandler(managerRepository storage.ManagerRepository) *Handler {
	return &Handler{
		ManagerRepository: managerRepository,
	}
}

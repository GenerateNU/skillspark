package child

import (
	"skillspark/internal/storage"
)

type Handler struct {
	ChildRepository storage.ChildRepository
}

func NewHandler(childRepository storage.ChildRepository) *Handler {
	return &Handler{
		ChildRepository: childRepository,
	}
}

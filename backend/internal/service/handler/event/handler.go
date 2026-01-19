package event

import (
	"skillspark/internal/storage"
)

type Handler struct {
	EventRepository storage.EventRepository
}

func NewHandler(eventRepository storage.EventRepository) *Handler {
	return &Handler{
		EventRepository: eventRepository,
	}
}

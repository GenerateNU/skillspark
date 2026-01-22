package eventoccurrence

import (
	"skillspark/internal/storage"
)

type Handler struct {
	EventOccurrenceRepository storage.EventOccurrenceRepository
}

func NewHandler(eventOccurrenceRepository storage.EventOccurrenceRepository) *Handler {
	return &Handler{
		EventOccurrenceRepository: eventOccurrenceRepository,
	}
}
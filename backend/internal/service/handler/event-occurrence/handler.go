package eventoccurrence

import (
	"skillspark/internal/storage"
)

type Handler struct {
	EventOccurrenceRepository storage.EventOccurrenceRepository
	ManagerRepository storage.ManagerRepository
	EventRepository storage.EventRepository
	LocationRepository storage.LocationRepository
}

func NewHandler(
	eventOccurrenceRepository storage.EventOccurrenceRepository, 
	managerRepository storage.ManagerRepository,
	eventRepository storage.EventRepository,
	locationRepository storage.LocationRepository) *Handler {
	return &Handler{
		EventOccurrenceRepository: eventOccurrenceRepository,
		ManagerRepository: managerRepository,
		EventRepository: eventRepository,
		LocationRepository: locationRepository,
	}
}
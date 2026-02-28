package eventoccurrence

import (
	"skillspark/internal/s3_client"
	"skillspark/internal/storage"
)

type Handler struct {
	EventOccurrenceRepository storage.EventOccurrenceRepository
	ManagerRepository         storage.ManagerRepository
	EventRepository           storage.EventRepository
	LocationRepository        storage.LocationRepository
	s3Client                  s3_client.S3Interface
}

func NewHandler(
	eventOccurrenceRepository storage.EventOccurrenceRepository,
	managerRepository storage.ManagerRepository,
	eventRepository storage.EventRepository,
	locationRepository storage.LocationRepository, s3client s3_client.S3Interface) *Handler {
	return &Handler{
		EventOccurrenceRepository: eventOccurrenceRepository,
		ManagerRepository:         managerRepository,
		EventRepository:           eventRepository,
		LocationRepository:        locationRepository,
		s3Client:                  s3client,
	}
}

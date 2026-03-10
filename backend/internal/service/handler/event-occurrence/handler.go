package eventoccurrence

import (
	"skillspark/internal/s3_client"
	"skillspark/internal/storage"
	"skillspark/internal/stripeClient"
)

type Handler struct {
	EventOccurrenceRepository storage.EventOccurrenceRepository
	ManagerRepository         storage.ManagerRepository
	EventRepository           storage.EventRepository
	LocationRepository        storage.LocationRepository
	s3Client                  s3_client.S3Interface
	RegistrationRepository	  storage.RegistrationRepository
	StripeClient              stripeClient.StripeClientInterface
}

func NewHandler(
	eventOccurrenceRepository storage.EventOccurrenceRepository,
	managerRepository storage.ManagerRepository,
	eventRepository storage.EventRepository,
	locationRepository storage.LocationRepository,
  s3client s3_client.S3Interface,
	registrationRepository storage.RegistrationRepository,
	stripeClient stripeClient.StripeClientInterface) *Handler {
	return &Handler{
		EventOccurrenceRepository: eventOccurrenceRepository,
		ManagerRepository:         managerRepository,
		EventRepository:           eventRepository,
		LocationRepository:        locationRepository,
		s3Client:                  s3client,
		RegistrationRepository:    registrationRepository,
		StripeClient: 			   stripeClient,	
	}
}

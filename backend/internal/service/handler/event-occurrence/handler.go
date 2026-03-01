package eventoccurrence

import (
	"skillspark/internal/storage"
	"skillspark/internal/stripeClient"
)

type Handler struct {
	EventOccurrenceRepository storage.EventOccurrenceRepository
	ManagerRepository         storage.ManagerRepository
	EventRepository           storage.EventRepository
	LocationRepository        storage.LocationRepository
	RegistrationRepository	  storage.RegistrationRepository
	StripeClient              stripeClient.StripeClientInterface
}

func NewHandler(
	eventOccurrenceRepository storage.EventOccurrenceRepository,
	managerRepository storage.ManagerRepository,
	eventRepository storage.EventRepository,
	locationRepository storage.LocationRepository,
	registrationRepository storage.RegistrationRepository,
	stripeClient stripeClient.StripeClientInterface) *Handler {
	return &Handler{
		EventOccurrenceRepository: eventOccurrenceRepository,
		ManagerRepository:         managerRepository,
		EventRepository:           eventRepository,
		LocationRepository:        locationRepository,
		RegistrationRepository:    registrationRepository,
		StripeClient: 			   stripeClient,	
	}
}

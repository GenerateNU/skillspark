package registration

import (
	"skillspark/internal/storage"
	"skillspark/internal/stripeClient"
	"skillspark/internal/notification"
)

type Handler struct {
	RegistrationRepository    storage.RegistrationRepository
	EventOccurrenceRepository storage.EventOccurrenceRepository
	GuardianRepository        storage.GuardianRepository
	ChildRepository           storage.ChildRepository
	OrganizationRepository    storage.OrganizationRepository
	StripeClient 			  stripeClient.StripeClientInterface
	NotificationService       *notification.Service
}

func NewHandler(registrationRepo storage.RegistrationRepository, childRepo storage.ChildRepository, 
	guardianRepo storage.GuardianRepository, eventOccurrenceRepo storage.EventOccurrenceRepository, 
	organizationRepo storage.OrganizationRepository, sc stripeClient.StripeClientInterface, notifService *notification.Service) *Handler {
	return &Handler{
		RegistrationRepository:    registrationRepo,
		ChildRepository:           childRepo,
		GuardianRepository:        guardianRepo,
		EventOccurrenceRepository: eventOccurrenceRepo,
		NotificationService:       notifService,
		OrganizationRepository:    organizationRepo,
		StripeClient:              sc,
	}
}

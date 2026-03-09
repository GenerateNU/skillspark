package registration

import (
	"skillspark/internal/storage"
	"skillspark/internal/stripeClient"
)

type Handler struct {
	RegistrationRepository    storage.RegistrationRepository
	EventOccurrenceRepository storage.EventOccurrenceRepository
	GuardianRepository        storage.GuardianRepository
	ChildRepository           storage.ChildRepository
	OrganizationRepository    storage.OrganizationRepository
	StripeClient              stripeClient.StripeClientInterface
}

func NewHandler(registrationRepo storage.RegistrationRepository, childRepo storage.ChildRepository, 
	guardianRepo storage.GuardianRepository, eventOccurrenceRepo storage.EventOccurrenceRepository, 
	organizationRepo storage.OrganizationRepository, sc stripeClient.StripeClientInterface) *Handler {
	return &Handler{
		RegistrationRepository:    registrationRepo,
		ChildRepository:           childRepo,
		GuardianRepository:        guardianRepo,
		EventOccurrenceRepository: eventOccurrenceRepo,
		OrganizationRepository:    organizationRepo,
		StripeClient:              sc,
	}
}

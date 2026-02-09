package payment

import (
	"skillspark/internal/stripeClient"
	"skillspark/internal/storage"
)

type Handler struct {
	OrganizationRepository storage.OrganizationRepository
	ManagerRepository      storage.ManagerRepository
	RegistrationRepository storage.RegistrationRepository
	LocationRepository     storage.LocationRepository
	StripeClient           stripeClient.StripeClient
}

func NewHandler(orgRepo storage.OrganizationRepository, managerRepo storage.ManagerRepository, registrationRepo storage.RegistrationRepository, locRepo storage.LocationRepository, sc stripeClient.StripeClient) *Handler {
	return &Handler{
		OrganizationRepository: orgRepo,
		ManagerRepository:     managerRepo,
		RegistrationRepository: registrationRepo,
		LocationRepository: locRepo,
		StripeClient: sc,
	}
}




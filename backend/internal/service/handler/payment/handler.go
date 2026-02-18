package payment

import (
	"skillspark/internal/storage"
	"skillspark/internal/stripeClient"
)

type Handler struct {
	OrganizationRepository storage.OrganizationRepository
	ManagerRepository      storage.ManagerRepository
	RegistrationRepository storage.RegistrationRepository
	LocationRepository     storage.LocationRepository
	GuardianRepository     storage.GuardianRepository
	StripeClient           stripeClient.StripeClientInterface
}

func NewHandler(
	orgRepo storage.OrganizationRepository,
	managerRepo storage.ManagerRepository,
	registrationRepo storage.RegistrationRepository,
	locRepo storage.LocationRepository,
	guardianRepo storage.GuardianRepository,
	sc stripeClient.StripeClientInterface,
) *Handler {
	return &Handler{
		OrganizationRepository: orgRepo,
		ManagerRepository:      managerRepo,
		RegistrationRepository: registrationRepo,
		LocationRepository:     locRepo,
		GuardianRepository:     guardianRepo,
		StripeClient:           sc,
	}
}

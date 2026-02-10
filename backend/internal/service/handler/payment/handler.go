package payment

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage"
	"skillspark/internal/stripeClient"
)

type Handler struct {
	OrganizationRepository storage.OrganizationRepository
	ManagerRepository      storage.ManagerRepository
	RegistrationRepository storage.RegistrationRepository
	LocationRepository     storage.LocationRepository
	GuardianRepository     storage.GuardianRepository
	StripeClient           *stripeClient.StripeClient
}

func (h *Handler) CreateAccountOnboardingLink(ctx context.Context, input *models.CreateStripeOnboardingLinkInput) (*models.CreateStripeOnboardingLinkOutput, error) {
	panic("unimplemented")
}

func NewHandler(
	orgRepo storage.OrganizationRepository,
	managerRepo storage.ManagerRepository,
	registrationRepo storage.RegistrationRepository,
	locRepo storage.LocationRepository,
	guardianRepo storage.GuardianRepository,
	sc *stripeClient.StripeClient,
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

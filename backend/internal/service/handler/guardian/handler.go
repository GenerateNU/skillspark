package guardian

import (
	"skillspark/internal/storage"
	"skillspark/internal/stripeClient"
)

type Handler struct {
	GuardianRepository storage.GuardianRepository
	StripeClient stripeClient.StripeClientInterface
	}

func NewHandler(guardianRepository storage.GuardianRepository, sc stripeClient.StripeClientInterface) *Handler {
	return &Handler{
		GuardianRepository: guardianRepository,
		StripeClient: sc,

	}
}

package guardian

import (
	"skillspark/internal/storage"
	"skillspark/internal/stripeClient"
)

type Handler struct {
	GuardianRepository storage.GuardianRepository
	StripeClient stripeClient.StripeClient
	}

func NewHandler(guardianRepository storage.GuardianRepository, sc *stripeClient.StripeClient) *Handler {
	return &Handler{
		GuardianRepository: guardianRepository,
		StripeClient: *sc,

	}
}

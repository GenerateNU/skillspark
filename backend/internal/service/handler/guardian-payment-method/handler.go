package guardianpaymentmethod

import (
	"skillspark/internal/stripeClient"
	"skillspark/internal/storage"
)

type Handler struct {
	GuardianRepository              storage.GuardianRepository
	GuardianPaymentMethodRepository storage.GuardianPaymentMethodRepository
	StripeClient                    *stripeClient.StripeClient
}

func NewHandler(
	guardianRepo storage.GuardianRepository,
	paymentMethodRepo storage.GuardianPaymentMethodRepository,
	sc *stripeClient.StripeClient,
) *Handler {
	return &Handler{
		GuardianRepository:              guardianRepo,
		GuardianPaymentMethodRepository: paymentMethodRepo,
		StripeClient:                    sc,
	}
}
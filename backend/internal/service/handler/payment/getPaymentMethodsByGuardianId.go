package payment

import (
	"context"
	"errors"
	"skillspark/internal/models"
)

func (h *Handler) GetPaymentMethodsByGuardianID(ctx context.Context, input *models.GetPaymentMethodsByGuardianIDInput) (*models.GetPaymentMethodsByGuardianIDOutput, error) {
	guardian, err := h.GuardianRepository.GetGuardianByID(ctx, input.GuardianID)
	if err != nil {
		return nil, err
	}

	if guardian.StripeCustomerID == nil {
		return nil, errors.New("guardian does not have a Stripe customer account")
	}

	paymentMethods, err := h.StripeClient.GetPaymentMethodsByCustomerID(ctx, *guardian.StripeCustomerID)
	if err != nil {
		return nil, err
	}

	return paymentMethods, nil
}

package guardianpaymentmethod

import (
	"context"
	"skillspark/internal/models"
)

func (h *Handler) GetPaymentMethodsByGuardianID(
	ctx context.Context,
	input *models.GetGuardianPaymentMethodsByGuardianIDInput,
) (*models.GetGuardianPaymentMethodsByGuardianIDOutput, error) {
	
	guardian, err := h.GuardianRepository.GetGuardianByID(ctx, input.GuardianID)
	if err != nil {
		return nil, err
	}

	if guardian.StripeCustomerID == nil || *guardian.StripeCustomerID == "" {
		return &models.GetGuardianPaymentMethodsByGuardianIDOutput{
			Body: []models.GuardianPaymentMethod{},
		}, nil
	}

	paymentMethods, err := h.GuardianPaymentMethodRepository.GetPaymentMethodsByGuardianID(ctx, input.GuardianID)
	if err != nil {
		return nil, err
	}

	return &models.GetGuardianPaymentMethodsByGuardianIDOutput{
		Body: paymentMethods,
	}, nil
}
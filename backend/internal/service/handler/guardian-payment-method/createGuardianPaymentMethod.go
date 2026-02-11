package guardianpaymentmethod

import (
	"context"
	"errors"
	"skillspark/internal/models"
)

func (h *Handler) CreateGuardianPaymentMethod(
	ctx context.Context,
	input *models.CreateGuardianPaymentMethodInput,
) (*models.CreateGuardianPaymentMethodOutput, error) {
	
	guardian, err := h.GuardianRepository.GetGuardianByID(ctx, input.Body.GuardianID)
	if err != nil {
		return nil, err
	}

	if guardian.StripeCustomerID == nil || *guardian.StripeCustomerID == "" {
		return nil, errors.New("guardian must have stripe customer ID")
	}

	

	paymentMethod, err := h.GuardianPaymentMethodRepository.CreateGuardianPaymentMethod(ctx, input)
	if err != nil {
		return nil, err
	}

	return &models.CreateGuardianPaymentMethodOutput{
		Body: *paymentMethod,
	}, nil
}
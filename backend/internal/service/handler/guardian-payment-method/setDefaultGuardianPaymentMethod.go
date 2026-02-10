package guardianpaymentmethod

import (
	"context"
	"errors"
	"skillspark/internal/models"
)

func (h *Handler) SetDefaultGuardianPaymentMethod(
	ctx context.Context,
	input *models.SetDefaultPaymentMethodInput,
) (*models.SetDefaultPaymentMethodOutput, error) {
	
	guardian, err := h.GuardianRepository.GetGuardianByID(ctx, input.GuardianID)
	if err != nil {
		return nil, err
	}

	if guardian.StripeCustomerID == nil || *guardian.StripeCustomerID == "" {
		return nil, errors.New("guardian must have stripe customer ID")
	}

	updatedPaymentMethod, err := h.GuardianPaymentMethodRepository.UpdateGuardianPaymentMethod(ctx, input.PaymentMethodID, true)
	if err != nil {
		return nil, err
	}

	return &models.SetDefaultPaymentMethodOutput{
		Body: *updatedPaymentMethod,
	}, nil
}
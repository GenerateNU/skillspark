package guardianpaymentmethod

import (
	"context"
	"errors"
	"skillspark/internal/models"
)

func (h *Handler) DeleteGuardianPaymentMethod(
	ctx context.Context,
	input *models.DeleteGuardianPaymentMethodInput,
) (*models.DeleteGuardianPaymentMethodOutput, error) {
	
	paymentMethod, err := h.GuardianPaymentMethodRepository.DeleteGuardianPaymentMethod(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	err = h.StripeClient.DetachPaymentMethod(ctx, paymentMethod.StripePaymentMethodID)
	if err != nil {
		return nil, errors.New("payment method deleted from database but failed to detach from Stripe")
	}

	output := &models.DeleteGuardianPaymentMethodOutput{}
	output.Body.Message = "Payment method deleted successfully"

	return output, nil
}
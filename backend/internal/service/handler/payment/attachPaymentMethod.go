package payment

import (
	"context"
	"errors"
	"skillspark/internal/models"
)

func (h *Handler) AttachPaymentMethod(ctx context.Context, input *models.AttachPaymentMethodInput) (*models.AttachPaymentMethodOutput, error) {
	guardian, err := h.GuardianRepository.GetGuardianByID(ctx, input.GuardianID)
	if err != nil {
		return nil, err
	}

	if guardian.StripeCustomerID == nil {
		return nil, errors.New("guardian does not have a Stripe customer account")
	}

	if err := h.StripeClient.AttachPaymentMethod(ctx, input.Body.PaymentMethodID, *guardian.StripeCustomerID); err != nil {
		return nil, err
	}

	output := &models.AttachPaymentMethodOutput{}
	output.Body.PaymentMethodID = input.Body.PaymentMethodID
	output.Body.CustomerID = *guardian.StripeCustomerID

	return output, nil
}
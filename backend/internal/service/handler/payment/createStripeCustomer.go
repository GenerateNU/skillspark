package payment

import (
	"context"
	"errors"
	"skillspark/internal/models"
)

func (h *Handler) CreateStripeCustomer(ctx context.Context, input *models.CreateStripeCustomerInput) (*models.CreateStripeCustomerOutput, error) {

	guardian, err := h.GuardianRepository.GetGuardianByID(ctx, input.GuardianID)

	if (err != nil) {
		return nil, err
	}

	if (guardian.StripeCustomerID == nil) {
		return nil, errors.New("Customer already has a Stripe account.")
	}

	customer, err := h.StripeClient.CreateCustomer(ctx, guardian.Email, guardian.Name)

	if (err != nil) {
		return nil, err
	}

	updatedGuardian, err := h.GuardianRepository.SetStripeCustomerID(ctx, guardian.ID, customer.ID)

	if (err != nil) {
		return nil, err
	}

	output := models.CreateStripeCustomerOutput{}
	output.Body = *updatedGuardian

	return &output, nil
}
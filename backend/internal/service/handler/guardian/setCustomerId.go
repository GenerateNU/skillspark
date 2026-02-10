package guardian

import (
	"context"
	"errors"
	"skillspark/internal/models"
)

func (h *Handler) CreateStripeCustomer(
	ctx context.Context,
	input *models.CreateStripeCustomerInput,
) (*models.CreateStripeCustomerOutput, error) {
	
	guardian, err := h.GuardianRepository.GetGuardianByID(ctx, input.GuardianID)
	if err != nil {
		return nil, err
	}

	if guardian.StripeCustomerID != nil && *guardian.StripeCustomerID != "" {
		return nil, errors.New("guardian already has stripe customer")
	}

	customer, err := h.StripeClient.CreateCustomer(ctx, guardian.Email, guardian.Name)
	if err != nil {
		return nil, err
	}

	updatedGuardian, err := h.GuardianRepository.SetStripeCustomerID(ctx, guardian.ID, customer.ID)
	if err != nil {
		return nil, err
	}

	return &models.CreateStripeCustomerOutput{
		Body: *updatedGuardian,
	}, nil
}
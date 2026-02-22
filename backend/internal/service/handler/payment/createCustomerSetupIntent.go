package payment

import (
	"context"
	"errors"
	"skillspark/internal/models"
)

func (h *Handler) CreateSetupIntent(
	ctx context.Context,
	input *models.CreateSetupIntentInput,
) (*models.CreateSetupIntentOutput, error) {
	
	guardian, err := h.GuardianRepository.GetGuardianByID(ctx, input.GuardianID)
	if err != nil {
		return nil, err
	}
	
	if guardian.StripeCustomerID == nil || *guardian.StripeCustomerID == "" {
		return nil, errors.New("guardian must have stripe customer ID")
	}
	
	clientSecret, err := h.StripeClient.CreateSetupIntent(ctx, *guardian.StripeCustomerID)
	if err != nil {
		return nil, err
	}
	
	output := &models.CreateSetupIntentOutput{}
	output.Body.ClientSecret = clientSecret
	
	return output, nil
}
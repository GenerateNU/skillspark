package payment

import (
	"context"
	"errors"
	"skillspark/internal/models"
)

func (h *Handler) CreateAccountOnboardingLink(ctx context.Context, input *models.CreateStripeOnboardingLinkInput) (*models.CreateStripeOnboardingLinkOutput, error) {

	org, err := h.OrganizationRepository.GetOrganizationByID(ctx, input.OrganizationID)

	if (err != nil) {
		return nil, err
	}

	stripeId := org.StripeAccountID

	if stripeId == nil {
		return nil, errors.New("this organization does not have a Stripe Account");
	}

	stripeClientInput := models.CreateStripeOnboardingLinkClientInput{}
	stripeClientInput.Body.RefreshURL = input.Body.RefreshURL
	stripeClientInput.Body.ReturnURL = input.Body.ReturnURL
	stripeClientInput.Body.StripeAccountID = *stripeId

	link, err := h.StripeClient.CreateAccountOnboardingLink(ctx, &stripeClientInput)

	if (err != nil) {
		return nil, err
	}

	return link, nil
}
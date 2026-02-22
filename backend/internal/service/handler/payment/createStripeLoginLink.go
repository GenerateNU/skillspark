package payment

import (
	"context"
	"errors"
	"skillspark/internal/models"
)

func (h *Handler) CreateOrgLoginLink(
	ctx context.Context,
	input *models.CreateOrgLoginLinkInput,
) (*models.CreateOrgLoginLinkOutput, error) {
	
	org, err := h.OrganizationRepository.GetOrganizationByID(ctx, input.OrganizationID)
	if err != nil {
		return nil, err
	}
	
	if org.StripeAccountID == nil || *org.StripeAccountID == "" {
		return nil, errors.New("organization does not have associated Stripe Account.")
	}
	
	loginURL, err := h.StripeClient.CreateLoginLink(ctx, *org.StripeAccountID)
	if err != nil {
		return nil, err
	}
	
	output := &models.CreateOrgLoginLinkOutput{}
	output.Body.LoginURL = loginURL
	
	return output, nil
}
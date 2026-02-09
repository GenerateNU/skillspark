package payment

import (
	"context"
	"skillspark/internal/models"
	

)

func (h *Handler) CreateOrgStripeAccount(
	ctx context.Context,
	input *models.CreateOrgStripeAccountInput,
) (*models.CreateOrgStripeAccountOutput, error) {
	
	org, orgErr := h.OrganizationRepository.GetOrganizationByID(ctx, input.Body.OrganizationID)

	if orgErr != nil {
		return nil, orgErr
	}

	manager, manErr := h.ManagerRepository.GetManagerByOrgID(ctx, input.Body.OrganizationID)

	if manErr != nil {
		return nil, manErr
	}

	location, locErr := h.LocationRepository.GetLocationByOrganizationID(ctx, input.Body.OrganizationID)

	if locErr != nil {
		return nil, locErr
	}

	stripeAccount, err := h.StripeClient.CreateOrganizationAccount(ctx, org.Name, manager.Email, location.Country)

	if (err != nil) {
		return nil, err
	}

	return nil, nil
	
}
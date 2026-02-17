package payment

import (
	"context"
	"errors"
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

	if org.StripeAccountID != nil {
		return nil, errors.New("Stripe account already exists for this organization.")
	}

	manager, manErr := h.ManagerRepository.GetManagerByOrgID(ctx, input.Body.OrganizationID)

	if manErr != nil {
		return nil, manErr
	}

	location, locErr := h.LocationRepository.GetLocationByOrganizationID(ctx, input.Body.OrganizationID)

	if locErr != nil {
		return nil, locErr
	}

	stripeAccount, err := h.StripeClient.CreateOrganizationAccount(ctx, org.Name, manager.Email, GetCountryCode(location.Country))

	if (err != nil) {
		return nil, err
	}

	updatedOrg, err := h.OrganizationRepository.SetStripeAccountID(ctx, input.Body.OrganizationID, stripeAccount.Body.Account.ID)

	if (err != nil) {
		return nil, err
	}

	output := models.CreateOrgStripeAccountOutput{}
	output.Body.Account = *updatedOrg

	return &output, nil
	
}

var countryNameToCode = map[string]string{
	"Thailand":      "TH",
	"United States": "US",
}

func GetCountryCode(countryName string) (string) {
	code, ok := countryNameToCode[countryName]
	if !ok {
		return ""
	}
	return code
}
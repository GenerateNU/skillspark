package organization

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
)

func (h *Handler) CreateOrganization(ctx context.Context, input *models.CreateOrganizationInput) (*models.CreateOrganizationOutput, error) {
	if input.Body.LocationID != nil {
		if _, err := h.LocationRepository.GetLocationByID(ctx, *input.Body.LocationID); err != nil {
			return nil, errs.BadRequest("Invalid location_id: location does not exist")
		}
	}

	created, err := h.OrganizationRepository.CreateOrganization(ctx, input)
	if err != nil {
		return nil, err
	}

	return &models.CreateOrganizationOutput{Body: *created}, nil
}

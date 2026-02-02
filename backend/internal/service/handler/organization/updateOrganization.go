package organization

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
)

func (h *Handler) UpdateOrganization(ctx context.Context, input *models.UpdateOrganizationInput) (*models.UpdateOrganizationOutput, error) {
	if input.Body.LocationID != nil {
		if _, err := h.LocationRepository.GetLocationByID(ctx, *input.Body.LocationID); err != nil {
			return nil, errs.BadRequest("Invalid location_id: location does not exist")
		}
	}

	updated, err := h.OrganizationRepository.UpdateOrganization(ctx, input)
	if err != nil {
		return nil, err
	}

	return &models.UpdateOrganizationOutput{Body: *updated}, nil
}

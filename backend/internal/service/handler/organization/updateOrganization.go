package organization

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

func (h *Handler) UpdateOrganization(ctx context.Context, input *models.UpdateOrganizationInput) (*models.UpdateOrganizationOutput, error) {
	_, err := uuid.Parse(input.ID.String())
	if err != nil {
		return nil, errs.BadRequest("Invalid ID format")
	}

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
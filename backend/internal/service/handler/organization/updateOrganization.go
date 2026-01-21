package organization

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"time"

	"github.com/google/uuid"
)

// UpdateOrganization handles PATCH /organizations/:id
func (h *Handler) UpdateOrganization(ctx context.Context, input *models.UpdateOrganizationInput) (*models.UpdateOrganizationOutput, error) {
	id, err := uuid.Parse(input.ID.String())
	if err != nil {
		return nil, errs.BadRequest("Invalid ID format")
	}

	// Fetch existing organization
	existing, httpErr := h.OrganizationRepository.GetOrganizationByID(ctx, id)
	if httpErr != nil {
		return nil, httpErr
	}

	// Update only provided fields
	if input.Body.Name != nil {
		existing.Name = *input.Body.Name
	}
	if input.Body.Active != nil {
		existing.Active = *input.Body.Active
	}
	if input.Body.PfpS3Key != nil {
		existing.PfpS3Key = input.Body.PfpS3Key
	}
	if input.Body.LocationID != nil {
		// Validate location exists if provided
		_, locErr := h.LocationRepository.GetLocationByID(ctx, *input.Body.LocationID)
		if locErr != nil {
			return nil, errs.BadRequest("Invalid location_id: location does not exist")
		}
		existing.LocationID = input.Body.LocationID
	}

	existing.UpdatedAt = time.Now()

	httpErr = h.OrganizationRepository.UpdateOrganization(ctx, existing)
	if httpErr != nil {
		return nil, httpErr
	}

	return &models.UpdateOrganizationOutput{Body: *existing}, nil
}
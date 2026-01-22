package organization

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"time"

	"github.com/google/uuid"
)

func (h *Handler) CreateOrganization(ctx context.Context, input *models.CreateOrganizationInput) (*models.CreateOrganizationOutput, error) {
	org := models.Organization{
		ID:         uuid.New(),
		Name:       input.Body.Name,
		Active:     true,
		PfpS3Key:   input.Body.PfpS3Key,
		LocationID: input.Body.LocationID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if input.Body.Active != nil {
		org.Active = *input.Body.Active
	}

	if org.LocationID != nil {
		_, httpErr := h.LocationRepository.GetLocationByID(ctx, *org.LocationID)
		if httpErr != nil {
			return nil, errs.BadRequest("Invalid location_id: location does not exist")
		}
	}

	httpErr := h.OrganizationRepository.CreateOrganization(ctx, &org)
	if httpErr != nil {
		return nil, httpErr
	}

	return &models.CreateOrganizationOutput{Body: org}, nil
}
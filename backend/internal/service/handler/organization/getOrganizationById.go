package organization

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

func (h *Handler) GetOrganizationById(ctx context.Context, input *models.GetOrganizationByIDInput) (*models.GetOrganizationByIDOutput, error) {
	id, err := uuid.Parse(input.ID.String())
	if err != nil {
		return nil, errs.BadRequest("Invalid ID format")
	}

	organization, httpErr := h.OrganizationRepository.GetOrganizationByID(ctx, id)
	if httpErr != nil {
		return nil, httpErr
	}

	return &models.GetOrganizationByIDOutput{Body: *organization}, nil
}

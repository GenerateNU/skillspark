package organization

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

func (h *Handler) DeleteOrganization(ctx context.Context, input *models.DeleteOrganizationInput) (*models.DeleteOrganizationOutput, error) {
	id, err := uuid.Parse(input.ID.String())
	if err != nil {
		return nil, errs.BadRequest("Invalid ID format")
	}

	deleted, httpErr := h.OrganizationRepository.DeleteOrganization(ctx, id)
	if httpErr != nil {
		return nil, httpErr
	}


	return &models.DeleteOrganizationOutput{Body: *deleted}, nil
}
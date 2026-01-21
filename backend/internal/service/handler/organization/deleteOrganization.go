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

	httpErr := h.OrganizationRepository.DeleteOrganization(ctx, id)
	if httpErr != nil {
		return nil, httpErr
	}

	resp := &models.DeleteOrganizationOutput{}
	resp.Body.Message = "Organization deleted successfully"
	resp.Body.ID = id.String()

	return resp, nil
}
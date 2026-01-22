package manager

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

// GetManagerByOrgId handles GET /managers/:organization_id
func (h *Handler) GetManagerByOrgID(ctx context.Context, input *models.GetManagerByOrgIDInput) (*models.Manager, error) {
	id, err := uuid.Parse(input.OrganizationID.String())
	if err != nil {
		return nil, errs.BadRequest("Invalid ID format")
	}

	manager, httpErr := h.ManagerRepository.GetManagerByOrgID(ctx, id)
	if httpErr != nil {
		return nil, httpErr
	}
	return manager, nil
}

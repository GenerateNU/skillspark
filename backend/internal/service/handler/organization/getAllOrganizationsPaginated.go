package organization

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/utils"
)

func (h *Handler) GetAllOrganizations(ctx context.Context, pagination utils.Pagination) ([]models.Organization, error) {
	organizations, err := h.OrganizationRepository.GetAllOrganizations(ctx, pagination)
	if err != nil {
		return nil, err
	}
	return organizations, nil
}

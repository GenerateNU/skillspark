package organization

import (
	"context"
	"log"
	"skillspark/internal/models"
)

func (h *Handler) GetAllOrganizations(ctx context.Context, input *models.GetAllOrganizationsInput) (*models.GetAllOrganizationsOutput, error) {
	page := input.Page
	pageSize := input.PageSize
	offset := (page - 1) * pageSize

	organizations, totalCount, httpErr := h.OrganizationRepository.GetAllOrganizations(ctx, offset, pageSize)
	if httpErr != nil {
		return nil, httpErr
	}

	totalPages := (totalCount + pageSize - 1) / pageSize

	resp := &models.GetAllOrganizationsOutput{}
	resp.Body.Organizations = organizations
	resp.Body.Page = page
	resp.Body.PageSize = pageSize
	resp.Body.TotalCount = totalCount
	resp.Body.TotalPages = totalPages

	return resp, nil
}
package organization

import (
	"context"
	"log"
	"skillspark/internal/models"
)

func (h *Handler) GetAllOrganizations(ctx context.Context, input *models.GetAllOrganizationsInput) (*models.GetAllOrganizationsOutput, error) {
	log.Printf("GetAllOrganizations called - Page: %d, PageSize: %d", input.Page, input.PageSize)

	page := input.Page
	pageSize := input.PageSize
	offset := (page - 1) * pageSize

	log.Printf("Calling repository with offset: %d, pageSize: %d", offset, pageSize)

	organizations, totalCount, httpErr := h.OrganizationRepository.GetAllOrganizations(ctx, offset, pageSize)
	if httpErr != nil {
		log.Printf("Repository error: %v", httpErr)
		return nil, httpErr
	}

	log.Printf("Got %d organizations, total count: %d", len(organizations), totalCount)

	totalPages := (totalCount + pageSize - 1) / pageSize

	resp := &models.GetAllOrganizationsOutput{}
	resp.Body.Organizations = organizations
	resp.Body.Page = page
	resp.Body.PageSize = pageSize
	resp.Body.TotalCount = totalCount
	resp.Body.TotalPages = totalPages

	log.Printf("Returning response: %+v", resp)
	return resp, nil
}
package organization

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/s3_client"
	"skillspark/internal/utils"
	"time"
)

func (h *Handler) GetAllOrganizations(ctx context.Context, pagination utils.Pagination, s3Client *s3_client.Client) ([]models.Organization, []string, error) {
	organizations, err := h.OrganizationRepository.GetAllOrganizations(ctx, pagination)
	if err != nil {
		return nil, nil, err
	}

	var urls []string
	for i := 0; i < len(organizations); i++ {
		key := organizations[i].PfpS3Key
		if key != nil {
			url, httpErr := s3Client.GeneratePresignedURL(ctx, *key, time.Hour)
			if httpErr != nil {
				return nil, nil, httpErr
			}
			urls = append(urls, url)

		} else {
			urls = append(urls, "")
		}
	}

	return organizations, urls, nil

}

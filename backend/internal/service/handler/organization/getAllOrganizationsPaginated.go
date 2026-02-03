package organization

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/s3_client"
	"skillspark/internal/utils"
	"time"
)

func (h *Handler) GetAllOrganizations(ctx context.Context, pagination utils.Pagination, s3Client s3_client.S3Interface) ([]models.Organization, error) {
	organizations, err := h.OrganizationRepository.GetAllOrganizations(ctx, pagination)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(organizations); i++ {
		key := organizations[i].PfpS3Key
		if key != nil {
			url, httpErr := s3Client.GeneratePresignedURL(ctx, *key, time.Hour)
			if httpErr != nil {
				return nil, httpErr
			}
			organizations[i].PresignedURL = &url

		} else {
			organizations[i].PresignedURL = nil
		}
	}

	return organizations, nil

}

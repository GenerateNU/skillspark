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

	for idx := range organizations {
		_, err := GetAllS3Helper(ctx, s3Client, idx, organizations)
		if err != nil {
			return nil, err
		}
	}

	return organizations, nil

}

// helper for iterating over all organizations and grabbing presigned url
func GetAllS3Helper(ctx context.Context, s3Client s3_client.S3Interface, idx int, organizations []models.Organization) ([]models.Organization, error) {

	key := organizations[idx].PfpS3Key

	if key != nil {
		url, httpErr := s3Client.GeneratePresignedURL(ctx, *key, time.Hour)
		if httpErr != nil {
			return nil, httpErr
		}
		organizations[idx].PresignedURL = &url

	} else {
		organizations[idx].PresignedURL = nil
	}

	return organizations, nil
}

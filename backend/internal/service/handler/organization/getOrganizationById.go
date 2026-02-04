package organization

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/s3_client"
	"time"

	"github.com/google/uuid"
)

func (h *Handler) GetOrganizationById(ctx context.Context, input *models.GetOrganizationByIDInput, s3Client s3_client.S3Interface) (*models.Organization, error) {
	id, err := uuid.Parse(input.ID.String())
	if err != nil {
		return nil, errs.BadRequest("Invalid ID format")
	}

	organization, httpErr := h.OrganizationRepository.GetOrganizationByID(ctx, id)
	if httpErr != nil {
		return nil, httpErr
	}

	url, err := GetOrgS3Helper(ctx, s3Client, organization)
	if err != nil {
		return nil, err
	}

	organization.PresignedURL = url

	return organization, nil
}

// helper for generating s3 presigned url
func GetOrgS3Helper(ctx context.Context, s3Client s3_client.S3Interface, organization *models.Organization) (*string, error) {
	var key *string
	var url *string
	key = organization.PfpS3Key
	if key != nil {
		presignedURL, err := s3Client.GeneratePresignedURL(ctx, *key, time.Hour)
		if err != nil {
			return nil, err
		}

		url = &presignedURL
	}

	organization.PresignedURL = url
	return url, nil
}

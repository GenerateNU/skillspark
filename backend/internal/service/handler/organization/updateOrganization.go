package organization

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/s3_client"
)

func (h *Handler) UpdateOrganization(ctx context.Context, input *models.UpdateOrganizationInput, image_data *[]byte, s3Client s3_client.S3Interface) (*models.Organization, error) {
	if input.Body.LocationID != nil {
		if _, err := h.LocationRepository.GetLocationByID(ctx, *input.Body.LocationID); err != nil {
			return nil, errs.BadRequest("Invalid location_id: location does not exist")
		}
	}

	var key *string
	var url *string

	if image_data != nil {
		var err error
		url, key, err = h.UpdateOrgS3Helper(ctx, s3Client, input, image_data)
		if err != nil {
			return nil, err
		}
	}

	organization, updateErr := h.OrganizationRepository.UpdateOrganization(ctx, input, key)
	if updateErr != nil {
		return nil, updateErr
	}

	organization.PresignedURL = url

	return organization, nil
}

func (h *Handler) UpdateOrgS3Helper(ctx context.Context, s3Client s3_client.S3Interface, input *models.UpdateOrganizationInput, image_data *[]byte) (*string, *string, error) {
	key, genErr := h.generateS3Key(input.ID)
	if genErr != nil {
		return nil, nil, genErr
	}
	url, errr := s3Client.UploadImage(ctx, key, *image_data)
	if errr != nil {
		return nil, nil, errr
	}

	return url, key, nil
}

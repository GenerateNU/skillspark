package organization

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/s3_client"
)

func (h *Handler) CreateOrganization(ctx context.Context, input *models.CreateOrganizationInput, updateBody *models.UpdateOrganizationBody, image_data *[]byte, s3Client s3_client.S3Interface) (*models.Organization, error) {
	if input.Body.LocationID != nil {
		if _, err := h.LocationRepository.GetLocationByID(ctx, *input.Body.LocationID); err != nil {
			return nil, errs.BadRequest("Invalid location_id: location does not exist")
		}
	}

	var key *string
	var url *string

	organization, err := h.OrganizationRepository.CreateOrganization(ctx, input, key)
	if err != nil {
		return nil, err
	}

	if image_data != nil {

		url, err = h.CreateOrgS3Helper(ctx, s3Client, organization, updateBody, image_data)
		if err != nil {
			return nil, err
		}
	}

	organization.PresignedURL = url

	return organization, nil
}

func (h *Handler) CreateOrgS3Helper(ctx context.Context, s3Client s3_client.S3Interface, organization *models.Organization,
	updateBody *models.UpdateOrganizationBody, image_data *[]byte) (*string, error) {

	key, err := h.generateS3Key(organization.ID)
	if err != nil {
		return nil, err
	}
	url, errr := s3Client.UploadImage(ctx, key, *image_data)

	updateInput := &models.UpdateOrganizationInput{
		ID:   organization.ID,
		Body: *updateBody,
	}
	updateKeyValue, err := h.OrganizationRepository.UpdateOrganization(ctx, updateInput, key)
	if err != nil {
		return nil, nil
	}
	organization.PfpS3Key = updateKeyValue.PfpS3Key
	if errr != nil {
		return nil, errr
	}

	return url, nil

}

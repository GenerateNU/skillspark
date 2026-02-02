package organization

import (
	"context"
	"fmt"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/s3_client"
)

func (h *Handler) CreateOrganization(ctx context.Context, input *models.CreateOrganizationInput, updateBody *models.UpdateOrganizationBody, image_data *[]byte, s3Client *s3_client.Client) (*models.Organization, *string, error) {
	if input.Body.LocationID != nil {
		if _, err := h.LocationRepository.GetLocationByID(ctx, *input.Body.LocationID); err != nil {
			return nil, nil, errs.BadRequest("Invalid location_id: location does not exist")
		}
	}

	var key *string
	var url *string

	organization, err := h.OrganizationRepository.CreateOrganization(ctx, input, key)
	if err != nil {
		return nil, nil, err
	}

	if image_data != nil {

		key, error := h.generateS3Key(organization.ID)
		if error != nil {
			return nil, nil, error.(*errs.HTTPError)
		}
		uploadedUrl, errr := s3Client.UploadImage(ctx, key, *image_data)

		updateInput := &models.UpdateOrganizationInput{
			ID:   organization.ID,
			Body: *updateBody,
		}
		updateKeyValue, err := h.OrganizationRepository.UpdateOrganization(ctx, updateInput, key)
		if err != nil {
			fmt.Println("problem", err)
			return nil, nil, err
		}
		organization.PfpS3Key = updateKeyValue.PfpS3Key
		if errr != nil {
			return nil, nil, errr.(*errs.HTTPError)
		}
		url = uploadedUrl
	}

	return organization, url, nil
}

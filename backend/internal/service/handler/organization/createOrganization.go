package organization

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/s3_client"
)

func (h *Handler) CreateOrganization(ctx context.Context, input *models.CreateOrganizationInput, image_data *[]byte, s3Client *s3_client.Client) (*models.Organization, *string, error) {
	key, error := h.generateS3Key(*input.Body.LocationID)
	var url *string

	if error != nil {
		return nil, nil, error.(*errs.HTTPError)
	}

	if input.Body.LocationID != nil {
		if _, err := h.LocationRepository.GetLocationByID(ctx, *input.Body.LocationID); err != nil {
			return nil, nil, errs.BadRequest("Invalid location_id: location does not exist")
		}
	}

	if image_data != nil {
		uploadedUrl, errr := s3Client.UploadImage(ctx, key, image_data)
		if errr != nil {
			return nil, nil, errr.(*errs.HTTPError)
		}
		url = &uploadedUrl
	}

	organization, err := h.OrganizationRepository.CreateOrganization(ctx, input, &key)
	if err != nil {
		return nil, nil, err
	}

	return organization, url, nil
}

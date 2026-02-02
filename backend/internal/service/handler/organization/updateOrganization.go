package organization

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/s3_client"
)

func (h *Handler) UpdateOrganization(ctx context.Context, input *models.UpdateOrganizationInput, image_data *[]byte, s3Client *s3_client.Client) (*models.Organization, error) {
	if input.Body.LocationID != nil {
		if _, err := h.LocationRepository.GetLocationByID(ctx, *input.Body.LocationID); err != nil {
			return nil, errs.BadRequest("Invalid location_id: location does not exist")
		}
	}

	var key *string
	var url *string

	occurrence, err := h.OrganizationRepository.GetOrganizationByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	key = occurrence.PfpS3Key

	if image_data != nil {

		if key == nil {
			var genErr error
			key, genErr = h.generateS3Key(input.ID)
			if genErr != nil {
				return nil, genErr
			}
		}

		uploadedUrl, errr := s3Client.UploadImage(ctx, key, *image_data)
		if errr != nil {
			return nil, errr.(*errs.HTTPError)
		}
		url = uploadedUrl

	}

	organization, updateErr := h.OrganizationRepository.UpdateOrganization(ctx, input, key)
	organization.PresignedURL = url
	if updateErr != nil {
		return nil, updateErr
	}

	return organization, nil
}

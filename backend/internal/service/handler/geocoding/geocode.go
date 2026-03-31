package geocoding

import (
	"context"
	"skillspark/internal/models"
)

func (h *Handler) GeocodeAddress(ctx context.Context, input *models.GeocodeAddressInput) (*models.GeocodeAddressOutput, error) {
	lat, lng, httpErr := h.service.Geocode(ctx, input.Body.Address)
	if httpErr != nil {
		return nil, httpErr
	}

	out := &models.GeocodeAddressOutput{}
	out.Body.Latitude = *lat
	out.Body.Longitude = *lng
	return out, nil
}

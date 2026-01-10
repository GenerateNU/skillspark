package location

import (
	"context"
	"skillspark/internal/models"
)

// CreateLocation handles POST /locations
func (h *Handler) CreateLocation(ctx context.Context, input *models.CreateLocationInput) (*models.Location, error) {
	// Input is already parsed and validated by Huma!
	// Just pass it to the repository
	location, err := h.LocationRepository.CreateLocation(ctx, input)
	if err != nil {
		return nil, err
	}

	return location, nil
}

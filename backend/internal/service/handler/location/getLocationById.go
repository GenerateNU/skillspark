package location

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

// GetLocationById handles GET /locations/:id
func (h *Handler) GetLocationById(ctx context.Context, input *models.GetLocationByIDInput) (*models.Location, error) {
	id, err := uuid.Parse(input.ID.String())
	if err != nil {
		return nil, errs.BadRequest("Invalid ID format")
	}

	location, httpErr := h.LocationRepository.GetLocationByID(ctx, id)
	if httpErr != nil {
		return nil, httpErr
	}
	return location, nil
}

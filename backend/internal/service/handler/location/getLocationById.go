package location

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

// GetLocationById handles GET /locations/:id
func (h *Handler) GetLocationById(ctx context.Context, input *models.GetLocationByIDInput) (*models.Location, *errs.HTTPError) {
	id, err := uuid.Parse(input.ID.String())
	if err != nil {
		return nil, &errs.HTTPError{Code: 400, Message: "Invalid ID format"}
	}

	location, err := h.LocationRepository.GetLocationByID(ctx, id)
	if err != nil {
		return nil, err.(*errs.HTTPError)
	}
	return location, nil
}

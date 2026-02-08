package location

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/utils"
)

func (h *Handler) GetAllLocations(ctx context.Context, pagination utils.Pagination) ([]models.Location, *errs.HTTPError) {
	locations, err := h.LocationRepository.GetAllLocations(ctx, pagination)
	if err != nil {
		return nil, err.(*errs.HTTPError)
	}
	return locations, nil
}

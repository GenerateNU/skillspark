package school

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/utils"
)

func (h *Handler) GetAllSchools(ctx context.Context, pagination utils.Pagination) ([]models.School, error) {
	schools, err := h.SchoolRepository.GetAllSchools(ctx, pagination)
	if err != nil {
		return nil, err.(*errs.HTTPError)
	}
	return schools, nil
}

package school

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/utils"
)

func (h *Handler) GetAllSchools(ctx context.Context, pagination utils.Pagination) ([]models.School, *errs.HTTPError) {
	schools, err := h.SchoolRepository.GetAllSchools(ctx, pagination)
	if err != nil {
		return nil, err
	}
	return schools, nil
}

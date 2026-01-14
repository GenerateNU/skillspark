package school

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
)

func (h *Handler) GetAllSchools(ctx context.Context) ([]models.School, *errs.HTTPError) {
	schools, err := h.SchoolRepository.GetAllSchools(ctx)
	if err != nil {
		return nil, err
	}
	return schools, nil
}

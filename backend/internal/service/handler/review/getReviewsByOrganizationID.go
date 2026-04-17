package review

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/utils"

	"github.com/google/uuid"
)

func (h *Handler) GetReviewsByOrganizationID(ctx context.Context, id uuid.UUID, AcceptLanguage string, pagination utils.Pagination) ([]models.Review, error) {

	if _, err := h.OrganizationRepository.GetOrganizationByID(ctx, id, "en-US"); err != nil {
		return nil, errs.BadRequest("Invalid organization_id: organization does not exist" + err.Error())
	}

	reviews, httpErr := h.ReviewRepository.GetReviewsByOrganizationID(ctx, id, AcceptLanguage, pagination)
	if httpErr != nil {
		return nil, httpErr
	}

	return reviews, nil
}

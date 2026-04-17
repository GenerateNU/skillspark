package review

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/utils"

	"github.com/google/uuid"
)

func (h *Handler) GetEventReviewsForOrganization(ctx context.Context, id uuid.UUID, pagination utils.Pagination, AcceptLanguage string, sortBy string) ([]models.SimpleReviewAggregate, error) {

	// the language here does not matter, we are just checking that the organization exists.
	if _, err := h.OrganizationRepository.GetOrganizationByID(ctx, id, "en-US"); err != nil {
		return nil, errs.BadRequest("Invalid organization_id: organization does not exist: " + err.Error())
	}

	aggregate, httpErr := h.ReviewRepository.GetEventReviewsForOrganization(ctx, id, pagination, AcceptLanguage, sortBy)
	if httpErr != nil {
		return nil, httpErr
	}

	return aggregate, nil
}

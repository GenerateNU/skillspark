package review

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

func (h *Handler) GetAggregateReviewsForOrganization(ctx context.Context, id uuid.UUID) (*models.ReviewAggregate, error) {

	// the language here does not matter, we are just checking that the organization exists.
	if _, err := h.OrganizationRepository.GetOrganizationByID(ctx, id); err != nil {
		return nil, errs.BadRequest("Invalid organization_id: organization does not exist" + err.Error())
	}

	aggregate, httpErr := h.ReviewRepository.GetAggregateReviewsForOrganization(ctx, id)
	if httpErr != nil {
		return nil, httpErr
	}

	return aggregate, nil
}

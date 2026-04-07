package recommendation

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/utils"
	"time"

	"github.com/google/uuid"
)


func (h *Handler) GetRecommendationsByChildID(ctx context.Context, childID uuid.UUID, acceptLanguage string, pagination utils.Pagination, minDate *time.Time, maxDate *time.Time) ([]models.Event, error) {
	child, err := h.ChildRepository.GetChildByID(ctx, childID)
	if err != nil {
		return nil, err
	}

	return h.RecommendationRepository.GetRecommendationsByChildID(ctx, child.Interests, child.BirthYear, acceptLanguage, pagination, minDate, maxDate)
}

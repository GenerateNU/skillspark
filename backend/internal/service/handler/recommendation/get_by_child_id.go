package recommendation

import (
	"context"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

func (h *Handler) GetRecommendationsByChildID(ctx context.Context, childID uuid.UUID, acceptLanguage string) ([]models.EventOccurrence, error) {
	recommendations, err := h.RecommendationRepository.GetRecommendationsByChildID(ctx, childID, acceptLanguage, 5)
	if err != nil {
		return nil, err
	}
	return recommendations, nil
}

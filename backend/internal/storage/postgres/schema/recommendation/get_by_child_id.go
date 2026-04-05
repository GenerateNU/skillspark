package recommendation

import (
	"context"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

func (r *RecommendationRepository) GetRecommendationsByChildID(ctx context.Context, childID uuid.UUID, AcceptLanguage string, k int) ([]models.EventOccurrence, error) {
	// TODO: implement recommendation logic
	return []models.EventOccurrence{}, nil

	// call all event occurrences
	// filter by radius, date
	// perform mapping of tags to preferences
}

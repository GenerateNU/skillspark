package recommendation

import "skillspark/internal/storage"

type Handler struct {
	RecommendationRepository storage.RecommendationRepository
}

func NewHandler(repo storage.RecommendationRepository) *Handler {
	return &Handler{RecommendationRepository: repo}
}

package recommendation

import "skillspark/internal/storage"

type Handler struct {
	ChildRepository          storage.ChildRepository
	RecommendationRepository storage.RecommendationRepository
}

func NewHandler(repo *storage.Repository) *Handler {
	return &Handler{
		ChildRepository:          repo.Child,
		RecommendationRepository: repo.Recommendation,
	}
}

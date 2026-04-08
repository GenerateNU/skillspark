package recommendation

import (
	"skillspark/internal/s3_client"
	"skillspark/internal/storage"
)

type Handler struct {
	ChildRepository          storage.ChildRepository
	RecommendationRepository storage.RecommendationRepository
	S3Client                 s3_client.S3Interface
}

func NewHandler(repo *storage.Repository, s3Client s3_client.S3Interface) *Handler {
	return &Handler{
		ChildRepository:          repo.Child,
		RecommendationRepository: repo.Recommendation,
		S3Client:                 s3Client,
	}
}

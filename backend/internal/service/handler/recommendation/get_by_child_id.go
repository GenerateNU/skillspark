package recommendation

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/s3_client"
	"skillspark/internal/utils"
	"time"

	"github.com/google/uuid"
)

func (h *Handler) GetRecommendationsByChildID(ctx context.Context, childID uuid.UUID, acceptLanguage string, pagination utils.Pagination, filters models.RecommendationFilters) ([]models.Event, error) {
	child, err := h.ChildRepository.GetChildByID(ctx, childID)
	if err != nil {
		return nil, err
	}

	output, err := h.RecommendationRepository.GetRecommendationsByChildID(ctx, child.Interests, child.BirthYear, acceptLanguage, pagination, filters)
	if err != nil {
		return nil, err
	}

	for idx := range output {
		if err := AssignURL(ctx, &output[idx], h.S3Client); err != nil {
			return nil, err
		}
	}

	return output, nil
}

func AssignURL(ctx context.Context, event *models.Event, s3Client s3_client.S3Interface) error {

	key := event.HeaderImageS3Key
	if key != nil {
		url, err := s3Client.GeneratePresignedURL(ctx, *key, time.Hour)
		if err != nil {
			return err
		}

		event.PresignedURL = &url
	}

	return nil

}

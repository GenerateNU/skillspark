package routes

import (
	"context"
	"net/http"
	"skillspark/internal/models"
	recommendation "skillspark/internal/service/handler/recommendation"
	"skillspark/internal/storage"

	"github.com/danielgtaylor/huma/v2"
)

func SetupRecommendationRoutes(api huma.API, repo *storage.Repository) {
	recommendationHandler := recommendation.NewHandler(repo.Recommendation)

	huma.Register(api, huma.Operation{
		OperationID: "get-recommendations-by-child-id",
		Method:      http.MethodGet,
		Path:        "/api/v1/recommendations/{child_id}",
		Summary:     "Get recommendations by child ID",
		Description: "Returns a list of recommended event occurrences for a child",
		Tags:        []string{"Recommendations"},
	}, func(ctx context.Context, input *models.GetRecommendationsByChildIDInput) (*models.GetRecommendationsByChildIDOutput, error) {
		recommendations, err := recommendationHandler.GetRecommendationsByChildID(ctx, input.ChildID, input.AcceptLanguage)
		if err != nil {
			return nil, err
		}
		return &models.GetRecommendationsByChildIDOutput{
			Body: recommendations,
		}, nil
	})
}

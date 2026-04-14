package routes

import (
	"context"
	"net/http"
	"skillspark/internal/models"
	"skillspark/internal/s3_client"
	recommendation "skillspark/internal/service/handler/recommendation"
	"skillspark/internal/storage"
	"skillspark/internal/utils"

	"github.com/danielgtaylor/huma/v2"
)

func SetupRecommendationRoutes(api huma.API, repo *storage.Repository, s3Client s3_client.S3Interface) {

	recommendationHandler := recommendation.NewHandler(repo, s3Client)

	huma.Register(api, huma.Operation{
		OperationID: "get-recommendations-by-child-id",
		Method:      http.MethodGet,
		Path:        "/api/v1/recommendations/{child_id}",
		Summary:     "Get recommendations by child ID",
		Description: "Returns a list of recommended events for a child",
		Tags:        []string{"Recommendations"},
	}, func(ctx context.Context, input *models.GetRecommendationsByChildIDInput) (*models.GetRecommendationsByChildIDOutput, error) {
		pagination := utils.Pagination{Page: input.Page, Limit: input.Limit}
		filters := models.RecommendationFilters{
			Latitude:  input.Latitude,
			Longitude: input.Longitude,
			RadiusKm:  input.RadiusKm,
			MinDate:   input.MinDate,
			MaxDate:   input.MaxDate,
		}
		if !filters.MinDate.IsZero() && !filters.MaxDate.IsZero() && filters.MinDate.After(filters.MaxDate) {
			return nil, huma.Error400BadRequest("min_date must be less than or equal to max_date")
		}

		recommendations, err := recommendationHandler.GetRecommendationsByChildID(ctx, input.ChildID, input.AcceptLanguage, pagination, filters)
		if err != nil {
			return nil, err
		}
		return &models.GetRecommendationsByChildIDOutput{
			Body: recommendations,
		}, nil
	})
}

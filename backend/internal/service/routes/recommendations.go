package routes

import (
	"context"
	"net/http"
	"skillspark/internal/models"
	recommendation "skillspark/internal/service/handler/recommendation"
	"skillspark/internal/storage"
	"skillspark/internal/utils"
	"time"

	"github.com/danielgtaylor/huma/v2"
)

func SetupRecommendationRoutes(api huma.API, repo *storage.Repository) {
	recommendationHandler := recommendation.NewHandler(repo)

	huma.Register(api, huma.Operation{
		OperationID: "get-recommendations-by-child-id",
		Method:      http.MethodGet,
		Path:        "/api/v1/recommendations/{child_id}",
		Summary:     "Get recommendations by child ID",
		Description: "Returns a list of recommended events for a child",
		Tags:        []string{"Recommendations"},
	}, func(ctx context.Context, input *models.GetRecommendationsByChildIDInput) (*models.GetRecommendationsByChildIDOutput, error) {
		pagination := utils.Pagination{Page: input.Page, Limit: input.Limit}

		var minDate, maxDate *time.Time
		if !input.MinDate.IsZero() {
			minDate = &input.MinDate
		}
		if !input.MaxDate.IsZero() {
			maxDate = &input.MaxDate
		}
		if minDate != nil && maxDate != nil && minDate.After(*maxDate) {
			return nil, huma.Error400("min_date must be less than or equal to max_date")
		}

		recommendations, err := recommendationHandler.GetRecommendationsByChildID(ctx, input.ChildID, input.AcceptLanguage, pagination, minDate, maxDate)
		if err != nil {
			return nil, err
		}
		return &models.GetRecommendationsByChildIDOutput{
			Body: recommendations,
		}, nil
	})
}

package routes

import (
	"context"
	"net/http"
	"skillspark/internal/models"
	"skillspark/internal/service/handler/review"
	"skillspark/internal/storage"
	"skillspark/internal/utils"

	"github.com/danielgtaylor/huma/v2"
)

func SetUpReviewRoutes(api huma.API, repo *storage.Repository) {

	reviewHandler := review.NewHandler(repo.Registration, repo.Review, repo.Guardian, repo.Event)

	huma.Register(api, huma.Operation{
		OperationID: "get-review-by-guardian-id",
		Method:      http.MethodGet,
		Path:        "/api/v1/review/guardian/{id}",
		Summary:     "Get reviews by guardian ID",
		Description: "Returns all reviews with the given guardian ID",
		Tags:        []string{"Review"},
	}, func(ctx context.Context, input *models.GetReviewsByGuardianIDInput) (*models.ReviewsOutput, error) {

		page := input.Page
		if page == 0 {
			page = 1
		}
		limit := input.PageSize
		if limit == 0 {
			limit = 10
		}

		pagination := utils.Pagination{
			Page:  page,
			Limit: limit,
		}

		reviews, err := reviewHandler.GetReviewsByGuardianID(ctx, input.ID, pagination)
		if err != nil {
			return nil, err
		}

		return &models.ReviewsOutput{
			Body: reviews,
		}, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "get-review-by-event-id",
		Method:      http.MethodGet,
		Path:        "/api/v1/review/event/{id}",
		Summary:     "Get reviews by event ID",
		Description: "Returns all reviews with the given event ID",
		Tags:        []string{"Review"},
	}, func(ctx context.Context, input *models.GetReviewsByEventIDInput) (*models.ReviewsOutput, error) {

		page := input.Page
		if page == 0 {
			page = 1
		}
		limit := input.PageSize
		if limit == 0 {
			limit = 10
		}

		pagination := utils.Pagination{
			Page:  page,
			Limit: limit,
		}

		reviews, err := reviewHandler.GetReviewsByEventID(ctx, input.ID, pagination)
		if err != nil {
			return nil, err
		}

		return &models.ReviewsOutput{
			Body: reviews,
		}, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "delete-review",
		Method:      http.MethodDelete,
		Path:        "/api/v1/review/{id}",
		Summary:     "Delete an existing review by id",
		Description: "Deletes an existing review by id",
		Tags:        []string{"Review"},
	}, func(ctx context.Context, input *models.DeleteReviewInput) (*models.DeleteReviewOutput, error) {
		msg, err := reviewHandler.DeleteReview(ctx, input.ID)
		if err != nil {
			return nil, err
		}

		return &models.DeleteReviewOutput{
			Body: struct {
				Message string `json:"message" doc:"Success message"`
			}{
				Message: msg,
			},
		}, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "create-review",
		Method:      http.MethodPost,
		Path:        "/api/v1/review",
		Summary:     "Creates a review",
		Description: "Creates a review",
		Tags:        []string{"Review"},
	}, func(ctx context.Context, input *models.CreateReviewInput) (*models.CreateReviewOutput, error) {

		reviewOutput, err := reviewHandler.CreateReview(ctx, input)
		if err != nil {
			return nil, err
		}

		return reviewOutput, nil
	})

}

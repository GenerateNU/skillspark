package routes

import (
	"context"
	"net/http"
	"skillspark/internal/models"
	"skillspark/internal/service/handler/review"
	"skillspark/internal/storage"
	translations "skillspark/internal/translation"
	"skillspark/internal/utils"

	"github.com/danielgtaylor/huma/v2"
)

func SetUpReviewRoutes(api huma.API, repo *storage.Repository, translateClient translations.TranslationInterface) {

	reviewHandler := review.NewHandler(repo.Registration, repo.Review, repo.Guardian, repo.Event, repo.Organization, translateClient)

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

		reviews, err := reviewHandler.GetReviewsByGuardianID(ctx, input.ID, input.AcceptLanguage, pagination)
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

		reviews, err := reviewHandler.GetReviewsByEventID(ctx, input.ID, input.AcceptLanguage, pagination)
		if err != nil {
			return nil, err
		}

		return &models.ReviewsOutput{
			Body: reviews,
		}, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "get-review-by-organization-id",
		Method:      http.MethodGet,
		Path:        "/api/v1/review/organization/{id}",
		Summary:     "Get reviews by organization ID",
		Description: "Returns all reviews with the given organization ID",
		Tags:        []string{"Review"},
	}, func(ctx context.Context, input *models.GetReviewsByOrganizationIDInput) (*models.ReviewsOutput, error) {

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

		reviews, err := reviewHandler.GetReviewsByOrganizationID(ctx, input.ID, input.AcceptLanguage, pagination)
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

		if input.Body.Rating != 0 && (input.Body.Rating < 1 || input.Body.Rating > 5) {
			return nil, huma.Error400BadRequest("rating must be between 1 and 5")
		}

		reviewOutput, err := reviewHandler.CreateReview(ctx, input)
		if err != nil {
			return nil, err
		}

		return reviewOutput, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "get-review-aggregate",
		Method:      http.MethodGet,
		Path:        "/api/v1/review/event_aggregate/{id}",
		Summary:     "Get review aggregate by event id",
		Description: "Get review aggregate by event id",
		Tags:        []string{"Review"},
	}, func(ctx context.Context, input *models.GetReviewsAggregateInput) (*models.ReviewsAggregateOutput, error) {

		aggregate, err := reviewHandler.GetAggregateReviews(ctx, input.ID)

		if err != nil {
			return nil, err
		}

		return &models.ReviewsAggregateOutput{
			Body: *aggregate,
		}, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "get-review-aggregate-organization",
		Method:      http.MethodGet,
		Path:        "/api/v1/review/organization_aggregate/{id}",
		Summary:     "Get review aggregate by organization id",
		Description: "Get review aggregate by organization id",
		Tags:        []string{"Review"},
	}, func(ctx context.Context, input *models.GetReviewsAggregateInput) (*models.ReviewsAggregateOutput, error) {

		aggregate, err := reviewHandler.GetAggregateReviewsForOrganization(ctx, input.ID)

		if err != nil {
			return nil, err
		}

		return &models.ReviewsAggregateOutput{
			Body: *aggregate,
		}, nil
	})

}

package routes

import (
	"context"
	"net/http"
	"skillspark/internal/models"
	"skillspark/internal/opensearch"
	"skillspark/internal/s3_client"
	searchHandler "skillspark/internal/service/handler/search"
	"skillspark/internal/utils"

	"github.com/danielgtaylor/huma/v2"
)

func SetupSearchRoutes(api huma.API, osClient *opensearch.Client, s3 s3_client.S3Interface) {
	handler := searchHandler.NewHandler(osClient, s3)

	huma.Register(api, huma.Operation{
		OperationID: "search-events",
		Method:      http.MethodGet,
		Path:        "/api/v1/search/events",
		Summary:     "Fuzzy search events",
		Description: "Returns events matching the search query using fuzzy full-text search",
		Tags:        []string{"Search"},
	}, func(ctx context.Context, input *models.SearchEventsInput) (*models.SearchEventsOutput, error) {
		pagination := utils.Pagination{Page: input.Page, Limit: input.Limit}

		events, err := handler.SearchEvents(ctx, input.Query, input.AcceptLanguage, pagination)
		if err != nil {
			return nil, err
		}

		return &models.SearchEventsOutput{Body: events}, nil
	})
}

package routes

import (
	"context"
	"net/http"
	"skillspark/internal/models"
	"skillspark/internal/service/handler/location"
	"skillspark/internal/storage"
	"skillspark/internal/utils"

	"github.com/danielgtaylor/huma/v2"
)

func SetupLocationsRoutes(api huma.API, repo *storage.Repository) {
	locationHandler := location.NewHandler(repo.Location)
	huma.Register(api, huma.Operation{
		OperationID: "get-location-by-id",
		Method:      http.MethodGet,
		Path:        "/api/v1/locations/{id}",
		Summary:     "Get a location by id",
		Description: "Returns a location by id",
		Tags:        []string{"Locations"},
	}, func(ctx context.Context, input *models.GetLocationByIDInput) (*models.GetLocationByIDOutput, error) {
		location, err := locationHandler.GetLocationById(ctx, input)
		if err != nil {
			return nil, err
		}

		return &models.GetLocationByIDOutput{
			Body: location,
		}, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "post-location",
		Method:      http.MethodPost,
		Path:        "/api/v1/locations",
		Summary:     "Create a new location",
		Description: "Creates a new location",
		Tags:        []string{"Locations"},
	}, func(ctx context.Context, input *models.CreateLocationInput) (*models.CreateLocationOutput, error) {
		location, err := locationHandler.CreateLocation(ctx, input)
		if err != nil {
			return nil, err
		}

		return &models.CreateLocationOutput{
			Body: location,
		}, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "get-all-locations",
		Method:      http.MethodGet,
		Path:        "/api/v1/locations",
		Summary:     "Get all locations",
		Description: "Returns all locations",
		Tags:        []string{"Locations"},
	}, func(ctx context.Context, input *models.GetAllLocationsInput) (*models.GetAllLocationsOutput, error) {
		page := input.Page
		if page == 0 {
			page = 1
		}

		limit := input.Limit
		if limit == 0 {
			limit = 10
		}

		pagination := utils.Pagination{
			Page:  page,
			Limit: limit,
		}

		locations, err := locationHandler.GetAllLocations(ctx, pagination)
		if err != nil {
			return nil, err
		}

		return &models.GetAllLocationsOutput{
			Body: locations,
		}, nil
	})
}

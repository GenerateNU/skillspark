package routes

import (
	"context"
	"net/http"
	"skillspark/internal/models"
	"skillspark/internal/service/handler/location"
	"skillspark/internal/storage"

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
		Tags:        []string{"Examples"},
	}, func(ctx context.Context, input *models.GetLocationByIDInput) (*models.Location, error) {
		location, err := locationHandler.GetLocationById(ctx, input)
		if err != nil {
			return nil, err
		}
		resp := location
		return resp, nil
	})
}

package routes

import (
	"context"
	"net/http"
	"skillspark/internal/geocoding"
	geocodingHandler "skillspark/internal/service/handler/geocoding"
	"skillspark/internal/models"
	"skillspark/internal/storage"

	"github.com/danielgtaylor/huma/v2"
)

func SetupGeocodingRoutes(api huma.API, repo *storage.Repository) error {
	client, err := geocoding.NewClient()
	if err != nil {
		return err
	}

	service := geocoding.NewService(client, repo.GeocodeCache)
	handler := geocodingHandler.NewHandler(service)

	huma.Register(api, huma.Operation{
		OperationID: "geocode-address",
		Method:      http.MethodPost,
		Path:        "/api/v1/geocode",
		Summary:     "Geocode a text address",
		Description: "Validates a text address via OpenCage and returns its latitude and longitude. Results are cached in the database.",
		Tags:        []string{"Geocoding"},
	}, func(ctx context.Context, input *models.GeocodeAddressInput) (*models.GeocodeAddressOutput, error) {
		return handler.GeocodeAddress(ctx, input)
	})

	return nil
}

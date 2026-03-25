package routes

import (
	"context"
	"net/http"
	"skillspark/internal/geocoding"
	geocodingHandler "skillspark/internal/service/handler/geocoding"
	"skillspark/internal/models"

	"github.com/danielgtaylor/huma/v2"
)

func SetupGeocodingRoutes(api huma.API, svc geocoding.GeocoderServiceInterface) {
	registerGeocodingRoutes(api, geocodingHandler.NewHandler(svc))
}

// SetupGeocodingRoutesWithGeocoder registers geocoding routes using the provided
// GeocoderServiceInterface implementation. Intended for use in tests.
func SetupGeocodingRoutesWithGeocoder(api huma.API, svc geocoding.GeocoderServiceInterface) {
	registerGeocodingRoutes(api, geocodingHandler.NewHandler(svc))
}

func registerGeocodingRoutes(api huma.API, handler *geocodingHandler.Handler) {
	huma.Register(api, huma.Operation{
		OperationID: "geocode-address",
		Method:      http.MethodPost,
		Path:        "/api/v1/geocode",
		Summary:     "Geocode a text address",
		Description: "Validates a text address via OpenCage and returns its latitude and longitude.",
		Tags:        []string{"Geocoding"},
	}, func(ctx context.Context, input *models.GeocodeAddressInput) (*models.GeocodeAddressOutput, error) {
		return handler.GeocodeAddress(ctx, input)
	})
}

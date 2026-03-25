package location

import (
	"context"
	"fmt"
	"math"
	"skillspark/internal/errs"
	"skillspark/internal/models"
)

const maxGeocodeDistanceKm = 50.0

// CreateLocation handles POST /locations
func (h *Handler) CreateLocation(ctx context.Context, input *models.CreateLocationInput) (*models.Location, *errs.HTTPError) {
	address := fmt.Sprintf("%s, %s, %s, %s",
		input.Body.AddressLine1,
		input.Body.District,
		input.Body.Province,
		input.Body.Country,
	)
	geocodedLat, geocodedLng, httpErr := h.GeocodingService.Geocode(ctx, address)
	if httpErr != nil {
		return nil, httpErr
	}
	if input.Body.Latitude != nil && input.Body.Longitude != nil {
		dist := haversineKm(*input.Body.Latitude, *input.Body.Longitude, *geocodedLat, *geocodedLng)
		if dist > maxGeocodeDistanceKm {
			e := errs.BadRequest(fmt.Sprintf(
				"provided coordinates are %.1f km from the geocoded address (max allowed: %.0f km)",
				dist, maxGeocodeDistanceKm,
			))
			return nil, &e
		}
	}
	input.Body.Latitude = geocodedLat
	input.Body.Longitude = geocodedLng

	location, err := h.LocationRepository.CreateLocation(ctx, input)
	if err != nil {
		return nil, err.(*errs.HTTPError)
	}

	return location, nil
}

// haversineKm returns the great-circle distance in kilometres between two points.
func haversineKm(lat1, lng1, lat2, lng2 float64) float64 {
	const R = 6371.0
	φ1 := lat1 * math.Pi / 180
	φ2 := lat2 * math.Pi / 180
	Δφ := (lat2 - lat1) * math.Pi / 180
	Δλ := (lng2 - lng1) * math.Pi / 180
	a := math.Sin(Δφ/2)*math.Sin(Δφ/2) + math.Cos(φ1)*math.Cos(φ2)*math.Sin(Δλ/2)*math.Sin(Δλ/2)
	return R * 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
}

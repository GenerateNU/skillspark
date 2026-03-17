package geocoding

import (
	"context"
	"errors"
	"skillspark/internal/errs"
	"skillspark/internal/storage"
	"strings"
)

type Service struct {
	client *Client
	cache  storage.GeocodeCacheRepository
}

func NewService(client *Client, cache storage.GeocodeCacheRepository) *Service {
	return &Service{client: client, cache: cache}
}

// Geocode returns lat/lng for the given address.
// It normalizes the input, checks the DB cache first, and on a miss calls
// OpenCage, saves the result to the cache, then returns the coordinates
func (s *Service) Geocode(ctx context.Context, address string) (lat, lng float64, httpErr *errs.HTTPError) {
	normalized := strings.ToLower(strings.TrimSpace(address))

	cached, err := s.cache.GetGeocodeCache(ctx, normalized)
	if err == nil {
		return cached.Latitude, cached.Longitude, nil
	}

	lat, lng, err = s.client.Geocode(ctx, normalized)
	if err != nil {
		if errors.Is(err, ErrInvalidAddress) {
			e := errs.BadRequest(ErrInvalidAddress.Error())
			return 0, 0, &e
		}
		e := errs.InternalServerError("geocoding failed: " + err.Error())
		return 0, 0, &e
	}

	// Save to cache — best-effort, do not fail the request on cache write error
	_, _ = s.cache.CreateGeocodeCache(ctx, normalized, address, lat, lng)

	return lat, lng, nil
}

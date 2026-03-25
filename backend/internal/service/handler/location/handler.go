package location

import (
	"skillspark/internal/geocoding"
	"skillspark/internal/storage"
)

type Handler struct {
	LocationRepository storage.LocationRepository
	GeocodingService   geocoding.GeocoderServiceInterface
}

func NewHandler(locationRepository storage.LocationRepository, geocodingService geocoding.GeocoderServiceInterface) *Handler {
	return &Handler{
		LocationRepository: locationRepository,
		GeocodingService:   geocodingService,
	}
}

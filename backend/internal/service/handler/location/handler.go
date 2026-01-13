package location

import (
	"skillspark/internal/storage"
)

type Handler struct {
	LocationRepository storage.LocationRepository
}

func NewHandler(locationRepository storage.LocationRepository) *Handler {
	return &Handler{
		LocationRepository: locationRepository,
	}
}

package geocoding

import (
	"skillspark/internal/geocoding"
)

type Handler struct {
	service geocoding.GeocoderServiceInterface
}

func NewHandler(service geocoding.GeocoderServiceInterface) *Handler {
	return &Handler{service: service}
}

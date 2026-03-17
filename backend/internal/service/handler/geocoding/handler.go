package geocoding

import (
	"skillspark/internal/geocoding"
)

type Handler struct {
	service *geocoding.Service
}

func NewHandler(service *geocoding.Service) *Handler {
	return &Handler{service: service}
}

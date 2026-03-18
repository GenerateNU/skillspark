package geocoding

import (
	"context"
	"skillspark/internal/errs"
)

type Geocoder interface {
	Geocode(ctx context.Context, address string) (lat, lng float64, httpErr *errs.HTTPError)
}

type Handler struct {
	service Geocoder
}

func NewHandler(service Geocoder) *Handler {
	return &Handler{service: service}
}

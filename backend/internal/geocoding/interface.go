package geocoding

import (
	"context"
	"skillspark/internal/errs"
)

type GeocoderServiceInterface interface {
	Geocode(ctx context.Context, address string) (lat, lng *float64, httpErr *errs.HTTPError)
}

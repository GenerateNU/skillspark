package geocoding

import (
	"context"
	"errors"
	"skillspark/internal/errs"
)

type Service struct {
	client *Client
}

func NewService(client *Client) *Service {
	return &Service{client: client}
}

// Geocode returns lat/lng for the given address via OpenCage.
func (s *Service) Geocode(ctx context.Context, address string) (lat, lng float64, httpErr *errs.HTTPError) {
	var err error
	lat, lng, err = s.client.Geocode(ctx, address)
	if err != nil {
		if errors.Is(err, ErrInvalidAddress) {
			e := errs.BadRequest(ErrInvalidAddress.Error())
			return 0, 0, &e
		}
		e := errs.InternalServerError("geocoding failed: " + err.Error())
		return 0, 0, &e
	}

	return lat, lng, nil
}

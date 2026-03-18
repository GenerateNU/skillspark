package geocoding

import (
	"context"
	"net/http"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockGeocoder struct {
	mock.Mock
}

func (m *mockGeocoder) Geocode(ctx context.Context, address string) (float64, float64, *errs.HTTPError) {
	args := m.Called(ctx, address)
	lat := args.Get(0).(float64)
	lng := args.Get(1).(float64)
	if args.Get(2) == nil {
		return lat, lng, nil
	}
	return lat, lng, args.Get(2).(*errs.HTTPError)
}

func TestHandler_GeocodeAddress(t *testing.T) {
	badRequestCode := http.StatusBadRequest
	internalErrCode := http.StatusInternalServerError

	tests := []struct {
		name       string
		address    string
		mockSetup  func(*mockGeocoder)
		wantLat    float64
		wantLng    float64
		wantErrCode *int
	}{
		{
			name:    "successful geocode - Bangkok",
			address: "1 Sukhumvit Rd, Bangkok",
			mockSetup: func(m *mockGeocoder) {
				m.On("Geocode", mock.Anything, "1 Sukhumvit Rd, Bangkok").
					Return(13.7563, 100.5018, nil)
			},
			wantLat: 13.7563,
			wantLng: 100.5018,
		},
		{
			name:    "successful geocode - Chiang Mai",
			address: "Nimman Rd, Chiang Mai",
			mockSetup: func(m *mockGeocoder) {
				m.On("Geocode", mock.Anything, "Nimman Rd, Chiang Mai").
					Return(18.7883, 98.9853, nil)
			},
			wantLat: 18.7883,
			wantLng: 98.9853,
		},
		{
			name:    "invalid address returns bad request",
			address: "zzzzz not a real place",
			mockSetup: func(m *mockGeocoder) {
				e := errs.BadRequest("address is invalid or could not be geocoded with sufficient confidence")
				m.On("Geocode", mock.Anything, "zzzzz not a real place").
					Return(0.0, 0.0, &e)
			},
			wantErrCode: &badRequestCode,
		},
		{
			name:    "geocoding service failure returns internal server error",
			address: "Some Valid Address",
			mockSetup: func(m *mockGeocoder) {
				e := errs.InternalServerError("geocoding failed: connection refused")
				m.On("Geocode", mock.Anything, "Some Valid Address").
					Return(0.0, 0.0, &e)
			},
			wantErrCode: &internalErrCode,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			svc := new(mockGeocoder)
			tt.mockSetup(svc)

			handler := &Handler{service: svc}
			input := &models.GeocodeAddressInput{}
			input.Body.Address = tt.address

			out, err := handler.GeocodeAddress(context.Background(), input)

			if tt.wantErrCode != nil {
				assert.Nil(t, out)
				assert.NotNil(t, err)
				assert.Equal(t, *tt.wantErrCode, err.(*errs.HTTPError).Code)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, out)
				assert.Equal(t, tt.wantLat, out.Body.Latitude)
				assert.Equal(t, tt.wantLng, out.Body.Longitude)
			}

			svc.AssertExpectations(t)
		})
	}
}

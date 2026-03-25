package location

import (
	"context"
	"net/http"
	"skillspark/internal/errs"
	"skillspark/internal/geocoding"
	"skillspark/internal/models"
	repomocks "skillspark/internal/storage/repo-mocks"
	"skillspark/internal/utils"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockGeocodingService struct {
	mock.Mock
}

var _ geocoding.GeocoderServiceInterface = (*mockGeocodingService)(nil)

func (m *mockGeocodingService) Geocode(ctx context.Context, address string) (*float64, *float64, *errs.HTTPError) {
	args := m.Called(ctx, address)
	lat := args.Get(0).(float64)
	lng := args.Get(1).(float64)
	if args.Get(2) == nil {
		return &lat, &lng, nil
	}
	return nil, nil, args.Get(2).(*errs.HTTPError)
}

func TestHandler_GetLocationById(t *testing.T) {
	statusCodeNotFound := http.StatusNotFound
	messageSubstringNotFound := "Not found"
	statusCodeInternalServerError := http.StatusInternalServerError
	messageSubstringInternalServerError := "Internal server error"
	tests := []struct {
		name             string
		id               string
		mockSetup        func(*repomocks.MockLocationRepository)
		wantErr          bool
		statusCode       *int
		messageSubstring *string
	}{
		{
			name: "successful get location by id - New York",
			id:   "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
			mockSetup: func(m *repomocks.MockLocationRepository) {
				m.On("GetLocationByID", mock.Anything, uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")).Return(&models.Location{
					ID:           uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"),
					Latitude:     40.7128,
					Longitude:    -74.0060,
					AddressLine1: "123 Broadway",
					AddressLine2: nil,
					Subdistrict:  "Manhattan",
					District:     "New York County",
					Province:     "NY",
					PostalCode:   "10001",
					Country:      "USA",
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "successful get location by id - Boston",
			id:   "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a19",
			mockSetup: func(m *repomocks.MockLocationRepository) {
				m.On("GetLocationByID", mock.Anything, uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a19")).
					Return(&models.Location{
						ID:           uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a19"),
						Latitude:     42.3601,
						Longitude:    -71.0589,
						AddressLine1: "600 Boylston Street",
						District:     "Boston",
						Province:     "MA",
						PostalCode:   "02116",
						Country:      "USA",
						CreatedAt:    time.Now(),
						UpdatedAt:    time.Now(),
					}, nil)
			},
			wantErr: false,
		},
		{
			name: "location not found",
			id:   "00000000-0000-0000-0000-000000000000",
			mockSetup: func(m *repomocks.MockLocationRepository) {
				m.On("GetLocationByID", mock.Anything, uuid.MustParse("00000000-0000-0000-0000-000000000000")).
					Return(nil, &errs.HTTPError{
						Code:    errs.NotFound("Location", "id", "00000000-0000-0000-0000-000000000000").Code,
						Message: "Not found",
					})
			},
			wantErr:          true,
			statusCode:       &statusCodeNotFound,
			messageSubstring: &messageSubstringNotFound,
		},
		{
			name: "internal server error",
			id:   "ffffffff-ffff-ffff-ffff-ffffffffffff",
			mockSetup: func(m *repomocks.MockLocationRepository) {
				m.On("GetLocationByID", mock.Anything, uuid.MustParse("ffffffff-ffff-ffff-ffff-ffffffffffff")).
					Return(nil, &errs.HTTPError{
						Code:    errs.InternalServerError("Internal server error").Code,
						Message: "Internal server error",
					})
			},
			wantErr:          true,
			statusCode:       &statusCodeInternalServerError,
			messageSubstring: &messageSubstringInternalServerError,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable for parallel
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockLocationRepository)
			tt.mockSetup(mockRepo)

			handler := NewHandler(mockRepo, nil)
			ctx := context.Background()

			input := &models.GetLocationByIDInput{ID: uuid.MustParse(tt.id)}
			location, err := handler.GetLocationById(ctx, input)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Nil(t, location)
				assert.Equal(t, *tt.statusCode, err.Code)
				assert.Contains(t, err.Message, *tt.messageSubstring)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, location)
				assert.Equal(t, tt.id, location.ID.String())
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func ptr(f float64) *float64 { return &f }

func TestHandler_CreateLocation(t *testing.T) {
	tests := []struct {
		name         string
		input        *models.CreateLocationInput
		geocodeSetup func(*mockGeocodingService)
		repoSetup    func(*repomocks.MockLocationRepository)
		wantErr      bool
		wantErrCode  *int
		wantLat      float64
		wantLng      float64
	}{
		{
			name: "no lat/lng provided - uses geocoded values",
			input: func() *models.CreateLocationInput {
				input := &models.CreateLocationInput{}
				input.Body.AddressLine1 = "123 Broadway"
				input.Body.Subdistrict = "Manhattan"
				input.Body.District = "New York County"
				input.Body.Province = "NY"
				input.Body.PostalCode = "10001"
				input.Body.Country = "USA"
				return input
			}(),
			geocodeSetup: func(m *mockGeocodingService) {
				m.On("Geocode", mock.Anything, "123 Broadway, New York County, NY, USA").
					Return(40.7128, -74.0060, nil)
			},
			repoSetup: func(m *repomocks.MockLocationRepository) {
				m.On("CreateLocation", mock.Anything, mock.AnythingOfType("*models.CreateLocationInput")).Return(&models.Location{
					ID:           uuid.New(),
					Latitude:     40.7128,
					Longitude:    -74.0060,
					AddressLine1: "123 Broadway",
					Subdistrict:  "Manhattan",
					District:     "New York County",
					Province:     "NY",
					PostalCode:   "10001",
					Country:      "USA",
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
				}, nil)
			},
			wantLat: 40.7128,
			wantLng: -74.0060,
		},
		{
			name: "no lat/lng provided - Chiang Mai",
			input: func() *models.CreateLocationInput {
				input := &models.CreateLocationInput{}
				input.Body.AddressLine1 = "Nimman Rd"
				input.Body.Subdistrict = "Suthep"
				input.Body.District = "Mueang Chiang Mai"
				input.Body.Province = "Chiang Mai"
				input.Body.PostalCode = "50200"
				input.Body.Country = "Thailand"
				return input
			}(),
			geocodeSetup: func(m *mockGeocodingService) {
				m.On("Geocode", mock.Anything, "Nimman Rd, Mueang Chiang Mai, Chiang Mai, Thailand").
					Return(18.7883, 98.9853, nil)
			},
			repoSetup: func(m *repomocks.MockLocationRepository) {
				m.On("CreateLocation", mock.Anything, mock.AnythingOfType("*models.CreateLocationInput")).Return(&models.Location{
					ID:           uuid.New(),
					Latitude:     18.7883,
					Longitude:    98.9853,
					AddressLine1: "Nimman Rd",
					Subdistrict:  "Suthep",
					District:     "Mueang Chiang Mai",
					Province:     "Chiang Mai",
					PostalCode:   "50200",
					Country:      "Thailand",
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
				}, nil)
			},
			wantLat: 18.7883,
			wantLng: 98.9853,
		},
		{
			// Provided coords (~0.5 km from geocoded) — within 50 km threshold
			name: "lat/lng within range - accepted, geocoded values used",
			input: func() *models.CreateLocationInput {
				input := &models.CreateLocationInput{}
				input.Body.Latitude = ptr(13.7600)  // ~0.5 km from geocoded 13.7563
				input.Body.Longitude = ptr(100.5100) // ~0.5 km from geocoded 100.5018
				input.Body.AddressLine1 = "1 Sukhumvit Rd"
				input.Body.Subdistrict = "Khlong Toei"
				input.Body.District = "Khlong Toei"
				input.Body.Province = "Bangkok"
				input.Body.PostalCode = "10110"
				input.Body.Country = "Thailand"
				return input
			}(),
			geocodeSetup: func(m *mockGeocodingService) {
				m.On("Geocode", mock.Anything, "1 Sukhumvit Rd, Khlong Toei, Bangkok, Thailand").
					Return(13.7563, 100.5018, nil)
			},
			repoSetup: func(m *repomocks.MockLocationRepository) {
				m.On("CreateLocation", mock.Anything, mock.AnythingOfType("*models.CreateLocationInput")).Return(&models.Location{
					ID:           uuid.New(),
					Latitude:     13.7563,
					Longitude:    100.5018,
					AddressLine1: "1 Sukhumvit Rd",
					District:     "Khlong Toei",
					Province:     "Bangkok",
					PostalCode:   "10110",
					Country:      "Thailand",
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
				}, nil)
			},
			wantLat: 13.7563,
			wantLng: 100.5018,
		},
		{
			// Bangkok coords vs Chiang Mai geocode — ~680 km apart, exceeds 50 km threshold
			name: "lat/lng too far from geocoded address - returns bad request",
			input: func() *models.CreateLocationInput {
				input := &models.CreateLocationInput{}
				input.Body.Latitude = ptr(13.7563)  // Bangkok
				input.Body.Longitude = ptr(100.5018) // Bangkok
				input.Body.AddressLine1 = "Nimman Rd"
				input.Body.Subdistrict = "Suthep"
				input.Body.District = "Mueang Chiang Mai"
				input.Body.Province = "Chiang Mai"
				input.Body.PostalCode = "50200"
				input.Body.Country = "Thailand"
				return input
			}(),
			geocodeSetup: func(m *mockGeocodingService) {
				m.On("Geocode", mock.Anything, "Nimman Rd, Mueang Chiang Mai, Chiang Mai, Thailand").
					Return(18.7883, 98.9853, nil) // Chiang Mai
			},
			repoSetup:   func(*repomocks.MockLocationRepository) {},
			wantErr:     true,
			wantErrCode: func() *int { c := http.StatusBadRequest; return &c }(),
		},
		{
			name: "invalid address - geocoding returns bad request",
			input: func() *models.CreateLocationInput {
				input := &models.CreateLocationInput{}
				input.Body.AddressLine1 = "zzzzz not real"
				input.Body.Subdistrict = "Nowhere"
				input.Body.District = "Nowhere"
				input.Body.Province = "ZZ"
				input.Body.PostalCode = "00000"
				input.Body.Country = "ZZ"
				return input
			}(),
			geocodeSetup: func(m *mockGeocodingService) {
				e := errs.BadRequest("address is invalid or could not be geocoded with sufficient confidence")
				m.On("Geocode", mock.Anything, "zzzzz not real, Nowhere, ZZ, ZZ").
					Return(0.0, 0.0, &e)
			},
			repoSetup:   func(*repomocks.MockLocationRepository) {},
			wantErr:     true,
			wantErrCode: func() *int { c := http.StatusBadRequest; return &c }(),
		},
		{
			name: "geocoding service failure - returns internal server error",
			input: func() *models.CreateLocationInput {
				input := &models.CreateLocationInput{}
				input.Body.AddressLine1 = "123 Broadway"
				input.Body.Subdistrict = "Manhattan"
				input.Body.District = "New York County"
				input.Body.Province = "NY"
				input.Body.PostalCode = "10001"
				input.Body.Country = "USA"
				return input
			}(),
			geocodeSetup: func(m *mockGeocodingService) {
				e := errs.InternalServerError("geocoding failed: connection refused")
				m.On("Geocode", mock.Anything, "123 Broadway, New York County, NY, USA").
					Return(0.0, 0.0, &e)
			},
			repoSetup:   func(*repomocks.MockLocationRepository) {},
			wantErr:     true,
			wantErrCode: func() *int { c := http.StatusInternalServerError; return &c }(),
		},
		{
			name: "repository error - returns internal server error",
			input: func() *models.CreateLocationInput {
				input := &models.CreateLocationInput{}
				input.Body.AddressLine1 = "123 Broadway"
				input.Body.Subdistrict = "Manhattan"
				input.Body.District = "New York County"
				input.Body.Province = "NY"
				input.Body.PostalCode = "10001"
				input.Body.Country = "USA"
				return input
			}(),
			geocodeSetup: func(m *mockGeocodingService) {
				m.On("Geocode", mock.Anything, "123 Broadway, New York County, NY, USA").
					Return(40.7128, -74.0060, nil)
			},
			repoSetup: func(m *repomocks.MockLocationRepository) {
				m.On("CreateLocation", mock.Anything, mock.AnythingOfType("*models.CreateLocationInput")).Return(nil, &errs.HTTPError{
					Code:    errs.InternalServerError("Internal server error").Code,
					Message: "Internal server error",
				})
			},
			wantErr:     true,
			wantErrCode: func() *int { c := http.StatusInternalServerError; return &c }(),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockLocationRepository)
			tt.repoSetup(mockRepo)

			mockGeocoder := new(mockGeocodingService)
			tt.geocodeSetup(mockGeocoder)
			handler := NewHandler(mockRepo, mockGeocoder)

			location, err := handler.CreateLocation(context.Background(), tt.input)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Nil(t, location)
				assert.Equal(t, *tt.wantErrCode, err.Code)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, location)
				assert.Equal(t, tt.input.Body.AddressLine1, location.AddressLine1)
				assert.Equal(t, tt.wantLat, location.Latitude)
				assert.Equal(t, tt.wantLng, location.Longitude)
			}

			mockGeocoder.AssertExpectations(t)
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_GetAllLocations(t *testing.T) {
	statusCodeInternalServerError := http.StatusInternalServerError
	messageSubstringInternalServerError := "Internal server error"
	tests := []struct {
		name             string
		mockSetup        func(*repomocks.MockLocationRepository)
		wantErr          bool
		statusCode       *int
		messageSubstring *string
	}{
		{
			name: "successful get all locations",
			mockSetup: func(m *repomocks.MockLocationRepository) {
				m.On("GetAllLocations", mock.Anything, mock.AnythingOfType("utils.Pagination")).Return([]models.Location{
					{
						ID:           uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"),
						Latitude:     40.7128,
						Longitude:    -74.0060,
						AddressLine1: "123 Broadway",
						District:     "New York County",
						Province:     "NY",
						PostalCode:   "10001",
						Country:      "USA",
					},
					{
						ID:           uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a19"),
						Latitude:     42.3601,
						Longitude:    -71.0589,
						AddressLine1: "600 Boylston Street",
						District:     "Boston",
						Province:     "MA",
						PostalCode:   "02116",
						Country:      "USA",
					},
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "internal server error",
			mockSetup: func(m *repomocks.MockLocationRepository) {
				m.On("GetAllLocations", mock.Anything, mock.AnythingOfType("utils.Pagination")).
					Return(nil, &errs.HTTPError{
						Code:    errs.InternalServerError("Internal server error").Code,
						Message: "Internal server error",
					})
			},
			wantErr:          true,
			statusCode:       &statusCodeInternalServerError,
			messageSubstring: &messageSubstringInternalServerError,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockLocationRepository)
			tt.mockSetup(mockRepo)

			handler := NewHandler(mockRepo, nil)
			ctx := context.Background()

			pagination := utils.Pagination{
				Page:  1,
				Limit: 10,
			}

			locations, err := handler.GetAllLocations(ctx, pagination)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Nil(t, locations)
				assert.Equal(t, *tt.statusCode, err.Code)
				assert.Contains(t, err.Message, *tt.messageSubstring)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, locations)
				assert.Len(t, locations, 2)
				assert.Equal(t, "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11", locations[0].ID.String())
				assert.Equal(t, "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a19", locations[1].ID.String())
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

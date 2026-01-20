package location

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	repomocks "skillspark/internal/storage/repo-mocks"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_GetLocationById(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		mockSetup func(*repomocks.MockLocationRepository)
		wantErr   bool
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
			wantErr: true,
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
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable for parallel
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockLocationRepository)
			tt.mockSetup(mockRepo)

			handler := NewHandler(mockRepo)
			ctx := context.Background()

			input := &models.GetLocationByIDInput{ID: uuid.MustParse(tt.id)}
			location, err := handler.GetLocationById(ctx, input)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Nil(t, location)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, location)
				assert.Equal(t, tt.id, location.ID.String())
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_CreateLocation(t *testing.T) {
	tests := []struct {
		name      string
		input     *models.CreateLocationInput
		mockSetup func(*repomocks.MockLocationRepository)
		wantErr   bool
	}{
		{
			name: "create New York location",
			input: func() *models.CreateLocationInput {
				input := &models.CreateLocationInput{}
				input.Body.Latitude = 40.7128
				input.Body.Longitude = -74.0060
				input.Body.AddressLine1 = "123 Broadway"
				input.Body.AddressLine2 = nil
				input.Body.Subdistrict = "Manhattan"
				input.Body.District = "New York County"
				input.Body.Province = "NY"
				input.Body.PostalCode = "10001"
				input.Body.Country = "USA"
				return input
			}(),
			mockSetup: func(m *repomocks.MockLocationRepository) {
				m.On("CreateLocation", mock.Anything, mock.AnythingOfType("*models.CreateLocationInput")).Return(&models.Location{
					ID:           uuid.New(),
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
			name: "create Boston location",
			input: func() *models.CreateLocationInput {
				input := &models.CreateLocationInput{}
				input.Body.Latitude = 42.3601
				input.Body.Longitude = -71.0589
				input.Body.AddressLine1 = "600 Boylston Street"
				input.Body.District = "Boston"
				input.Body.Subdistrict = "Back Bay"
				input.Body.Province = "MA"
				input.Body.PostalCode = "02116"
				input.Body.Country = "USA"
				return input
			}(),
			mockSetup: func(m *repomocks.MockLocationRepository) {
				m.On("CreateLocation", mock.Anything, mock.AnythingOfType("*models.CreateLocationInput")).Return(&models.Location{
					ID:           uuid.New(),
					Latitude:     42.3601,
					Longitude:    -71.0589,
					AddressLine1: "600 Boylston Street",
					District:     "Boston",
					Subdistrict:  "Back Bay",
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
			name: "internal server error",
			input: func() *models.CreateLocationInput {
				input := &models.CreateLocationInput{}
				input.Body.Latitude = 40.7128
				input.Body.Longitude = -74.0060
				input.Body.AddressLine1 = "123 Broadway"
				input.Body.AddressLine2 = nil
				input.Body.Subdistrict = "Manhattan"
				input.Body.District = "New York County"
				input.Body.Province = "NY"
				input.Body.PostalCode = "10001"
				input.Body.Country = "USA"
				return input
			}(),
			mockSetup: func(m *repomocks.MockLocationRepository) {
				m.On("CreateLocation", mock.Anything, mock.AnythingOfType("*models.CreateLocationInput")).Return(nil, &errs.HTTPError{
					Code:    errs.InternalServerError("Internal server error").Code,
					Message: "Internal server error",
				})
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable for parallel
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := new(repomocks.MockLocationRepository)
			tt.mockSetup(mockRepo)

			handler := NewHandler(mockRepo)
			ctx := context.Background()

			location, err := handler.CreateLocation(ctx, tt.input)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Nil(t, location)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, location)
				assert.Equal(t, tt.input.Body.AddressLine1, location.AddressLine1)
				if tt.input.Body.AddressLine2 != nil {
					assert.Equal(t, *tt.input.Body.AddressLine2, *location.AddressLine2)
				} else {
					assert.Nil(t, location.AddressLine2)
				}
				assert.Equal(t, tt.input.Body.Subdistrict, location.Subdistrict)
				assert.Equal(t, tt.input.Body.District, location.District)
				assert.Equal(t, tt.input.Body.Province, location.Province)
				assert.Equal(t, tt.input.Body.PostalCode, location.PostalCode)
				assert.Equal(t, tt.input.Body.Country, location.Country)
				assert.Equal(t, tt.input.Body.Latitude, location.Latitude)
				assert.Equal(t, tt.input.Body.Longitude, location.Longitude)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

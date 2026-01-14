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
				m.On("GetLocationByID", mock.Anything, uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")).
					Return(&models.Location{
						ID:        uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"),
						Latitude:  40.7128,
						Longitude: -74.0060,
						Address:   "123 Broadway",
						City:      "New York",
						State:     "NY",
						ZipCode:   "10001",
						Country:   "USA",
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
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
						ID:        uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a19"),
						Latitude:  42.3601,
						Longitude: -71.0589,
						Address:   "600 Boylston Street",
						City:      "Boston",
						State:     "MA",
						ZipCode:   "02116",
						Country:   "USA",
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
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
				assert.Error(t, err)
				assert.Nil(t, location)
			} else {
				assert.NoError(t, err)
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
				input.Body.Address = "123 Broadway"
				input.Body.City = "New York"
				input.Body.State = "NY"
				input.Body.ZipCode = "10001"
				input.Body.Country = "USA"
				return input
			}(),
			mockSetup: func(m *repomocks.MockLocationRepository) {
				m.On("CreateLocation", mock.Anything, mock.AnythingOfType("*models.CreateLocationInput")).Return(&models.Location{
					ID:        uuid.New(),
					Latitude:  40.7128,
					Longitude: -74.0060,
					Address:   "123 Broadway",
					City:      "New York",
					State:     "NY",
					ZipCode:   "10001",
					Country:   "USA",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
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
				input.Body.Address = "600 Boylston Street"
				input.Body.City = "Boston"
				input.Body.State = "MA"
				input.Body.ZipCode = "02116"
				input.Body.Country = "USA"
				return input
			}(),
			mockSetup: func(m *repomocks.MockLocationRepository) {
				m.On("CreateLocation", mock.Anything, mock.AnythingOfType("*models.CreateLocationInput")).Return(&models.Location{
					ID:        uuid.New(),
					Latitude:  42.3601,
					Longitude: -71.0589,
					Address:   "600 Boylston Street",
					City:      "Boston",
					State:     "MA",
					ZipCode:   "02116",
					Country:   "USA",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "internal server error",
			input: func() *models.CreateLocationInput {
				input := &models.CreateLocationInput{}
				input.Body.Latitude = 37.7749
				input.Body.Longitude = -122.4194
				input.Body.Address = "700 Market Street"
				input.Body.City = "San Francisco"
				input.Body.State = "CA"
				input.Body.ZipCode = "94102"
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
				assert.Error(t, err)
				assert.Nil(t, location)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, location)
				assert.Equal(t, tt.input.Body.Address, location.Address)
				assert.Equal(t, tt.input.Body.City, location.City)
				assert.Equal(t, tt.input.Body.Latitude, location.Latitude)
				assert.Equal(t, tt.input.Body.Longitude, location.Longitude)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

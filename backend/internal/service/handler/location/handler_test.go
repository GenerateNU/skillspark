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
			name: "successful get location by id",
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
			name: "internal server error",
			id:   "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a19",
			mockSetup: func(m *repomocks.MockLocationRepository) {
				m.On("GetLocationByID", mock.Anything, uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a19")).Return(nil, &errs.HTTPError{
					Code:    errs.InternalServerError("Internal server error").Code,
					Message: "Internal server error",
				})
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(repomocks.MockLocationRepository)
			tt.mockSetup(mockRepo)

			handler := NewHandler(mockRepo)
			ctx := context.Background()

			input := &models.GetLocationByIDInput{
				ID: uuid.MustParse(tt.id),
			}

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
			name: "successful create location",
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
		t.Run(tt.name, func(t *testing.T) {
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

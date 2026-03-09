package repomocks

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/utils"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockLocationRepository struct {
	mock.Mock
}

func (m *MockLocationRepository) GetLocationByID(ctx context.Context, id uuid.UUID) (*models.Location, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return args.Get(0).(*models.Location), nil
}

func (m *MockLocationRepository) GetLocationByOrganizationID(ctx context.Context, orgID uuid.UUID) (*models.Location, error) {
	args := m.Called(ctx, orgID)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return args.Get(0).(*models.Location), nil
}

func (m *MockLocationRepository) CreateLocation(ctx context.Context, location *models.CreateLocationInput) (*models.Location, error) {
	args := m.Called(ctx, location)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return args.Get(0).(*models.Location), nil
}

func (m *MockLocationRepository) GetAllLocations(ctx context.Context, pagination utils.Pagination) ([]models.Location, error) {
	args := m.Called(ctx, pagination)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return args.Get(0).([]models.Location), nil
}

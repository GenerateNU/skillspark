package repomocks

import (
	"context"
	"skillspark/internal/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockEventRepository struct {
	mock.Mock
}

func (m *MockEventRepository) CreateEvent(ctx context.Context, input *models.CreateEventInput) (*models.Event, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return args.Get(0).(*models.Event), nil
}

func (m *MockEventRepository) UpdateEvent(ctx context.Context, input *models.UpdateEventInput) (*models.Event, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return args.Get(0).(*models.Event), nil
}

func (m *MockEventRepository) DeleteEvent(ctx context.Context, id uuid.UUID) (*struct{}, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*struct{}), args.Error(1)
}

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

func (m *MockEventRepository) CreateEvent(ctx context.Context, input *models.CreateEventDBInput, HeaderImageS3Key *string) (*models.Event, error) {
	args := m.Called(ctx, input, HeaderImageS3Key)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return args.Get(0).(*models.Event), nil
}

func (m *MockEventRepository) UpdateEvent(ctx context.Context, input *models.UpdateEventDBInput, HeaderImageS3Key *string) (*models.Event, error) {
	args := m.Called(ctx, input, HeaderImageS3Key)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return args.Get(0).(*models.Event), nil
}

func (m *MockEventRepository) DeleteEvent(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockEventRepository) GetEventOccurrencesByEventID(ctx context.Context, event_id uuid.UUID) ([]models.EventOccurrence, error) {
	args := m.Called(ctx, event_id)
	eventOccurrences := args.Get(0)
	if eventOccurrences == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return eventOccurrences.([]models.EventOccurrence), nil
}

func (m *MockEventRepository) GetEventByID(ctx context.Context, id uuid.UUID) (*models.Event, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return args.Get(0).(*models.Event), nil
}

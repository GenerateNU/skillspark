package repomocks

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/utils"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/mock"
)

type MockEventOccurrenceRepository struct {
	mock.Mock
}

func (m *MockEventOccurrenceRepository) GetAllEventOccurrences(ctx context.Context, pagination utils.Pagination, filter models.GetAllEventOccurrencesFilter) ([]models.EventOccurrence, error) {
	args := m.Called(ctx, pagination, filter)
	eventOccurrences := args.Get(0)
	if eventOccurrences == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return eventOccurrences.([]models.EventOccurrence), nil
}

func (m *MockEventOccurrenceRepository) GetEventOccurrenceByID(ctx context.Context, id uuid.UUID) (*models.EventOccurrence, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return args.Get(0).(*models.EventOccurrence), nil
}

func (m *MockEventOccurrenceRepository) CreateEventOccurrence(ctx context.Context, input *models.CreateEventOccurrenceInput) (*models.EventOccurrence, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return args.Get(0).(*models.EventOccurrence), nil
}

func (m *MockEventOccurrenceRepository) UpdateEventOccurrence(ctx context.Context, input *models.UpdateEventOccurrenceInput, tx *pgx.Tx) (*models.EventOccurrence, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return args.Get(0).(*models.EventOccurrence), nil
}

func (m *MockEventOccurrenceRepository) CancelEventOccurrence(
	ctx context.Context,
	id uuid.UUID,
) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

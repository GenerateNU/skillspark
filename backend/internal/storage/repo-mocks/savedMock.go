package repomocks

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/utils"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockSavedRepository struct {
	mock.Mock
}

func (m *MockSavedRepository) CreateSaved(ctx context.Context, input *models.CreateSavedInput) (*models.Saved, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return args.Get(0).(*models.Saved), args.Error(1)
}

func (m *MockSavedRepository) GetByGuardianID(ctx context.Context, id uuid.UUID, pagination utils.Pagination, acceptLanguage string) ([]models.Saved, error) {
	args := m.Called(ctx, id, pagination, acceptLanguage)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return args.Get(0).([]models.Saved), args.Error(1)
}

func (m *MockSavedRepository) DeleteSaved(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

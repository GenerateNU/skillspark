package repomocks

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/utils"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockReviewRepository struct {
	mock.Mock
}

func (m *MockReviewRepository) CreateReview(ctx context.Context, input *models.CreateReviewDBInput) (*models.Review, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return args.Get(0).(*models.Review), args.Error(1)
}

func (m *MockReviewRepository) GetReviewsByGuardianID(ctx context.Context, id uuid.UUID, acceptLanguage string, pagination utils.Pagination) ([]models.Review, error) {
	args := m.Called(ctx, id, acceptLanguage, pagination)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return args.Get(0).([]models.Review), args.Error(1)
}

func (m *MockReviewRepository) GetReviewsByEventID(ctx context.Context, id uuid.UUID, acceptLanguage string, pagination utils.Pagination) ([]models.Review, error) {
	args := m.Called(ctx, id, acceptLanguage, pagination)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return args.Get(0).([]models.Review), args.Error(1)
}

func (m *MockReviewRepository) DeleteReview(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

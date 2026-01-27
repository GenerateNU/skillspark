package repomocks

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockGuardianRepository struct {
	mock.Mock
}

func (m *MockGuardianRepository) CreateGuardian(ctx context.Context, guardian *models.CreateGuardianInput) (*models.Guardian, error) {
	args := m.Called(ctx, guardian)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*errs.HTTPError)
	}
	return args.Get(0).(*models.Guardian), nil
}

func (m *MockGuardianRepository) GetGuardianByChildID(ctx context.Context, childID uuid.UUID) (*models.Guardian, error) {
	args := m.Called(ctx, childID)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*errs.HTTPError)
	}
	return args.Get(0).(*models.Guardian), nil
}

func (m *MockGuardianRepository) GetGuardianByID(ctx context.Context, id uuid.UUID) (*models.Guardian, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*errs.HTTPError)
	}
	return args.Get(0).(*models.Guardian), nil
}

func (m *MockGuardianRepository) GetGuardianByUserID(ctx context.Context, userID uuid.UUID) (*models.Guardian, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*errs.HTTPError)
	}
	return args.Get(0).(*models.Guardian), nil
}

func (m *MockGuardianRepository) UpdateGuardian(ctx context.Context, guardian *models.UpdateGuardianInput) (*models.Guardian, error) {
	args := m.Called(ctx, guardian)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*errs.HTTPError)
	}
	return args.Get(0).(*models.Guardian), nil
}

func (m *MockGuardianRepository) DeleteGuardian(ctx context.Context, id uuid.UUID) (*models.Guardian, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*errs.HTTPError)
	}
	return args.Get(0).(*models.Guardian), nil
}

func (m *MockGuardianRepository) GetGuardianByAuthID(ctx context.Context, authID string) (*models.Guardian, error) {
	args := m.Called(ctx, authID)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*errs.HTTPError)
	}
	return args.Get(0).(*models.Guardian), nil
}

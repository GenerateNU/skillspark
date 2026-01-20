package repomocks

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockChildRepository struct {
	mock.Mock
}

func (m *MockChildRepository) GetChildByID(ctx context.Context, childID uuid.UUID) (*models.Child, *errs.HTTPError) {
	args := m.Called(ctx, childID)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*errs.HTTPError)
	}
	return args.Get(0).(*models.Child), nil
}

func (m *MockChildRepository) GetChildrenByParentID(ctx context.Context, parentID uuid.UUID) ([]models.Child, *errs.HTTPError) {
	args := m.Called(ctx, parentID)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*errs.HTTPError)
	}
	return args.Get(0).([]models.Child), nil
}

func (m *MockChildRepository) UpdateChildByID(ctx context.Context, child *models.UpdateChildInput) (*models.Child, *errs.HTTPError) {
	args := m.Called(ctx, child)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*errs.HTTPError)
	}
	return args.Get(0).(*models.Child), nil
}

func (m *MockChildRepository) CreateChild(ctx context.Context, child *models.CreateChildInput) (*models.Child, *errs.HTTPError) {
	args := m.Called(ctx, child)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*errs.HTTPError)
	}
	return args.Get(0).(*models.Child), nil
}

func (m *MockChildRepository) DeleteChildByID(ctx context.Context, childID uuid.UUID) (*models.Child, *errs.HTTPError) {
	args := m.Called(ctx, childID)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*errs.HTTPError)
	}
	return args.Get(0).(*models.Child), nil
}

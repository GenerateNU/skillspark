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
	return args.Get(0).(*models.Child), args.Get(1).(*errs.HTTPError)
}

func (m *MockChildRepository) GetChildrenByParentID(ctx context.Context, parentID uuid.UUID) ([]models.Child, *errs.HTTPError) {
	args := m.Called(ctx, parentID)
	return args.Get(0).([]models.Child), args.Get(1).(*errs.HTTPError)
}

func (m *MockChildRepository) UpdateChildByID(ctx context.Context, childID uuid.UUID, child *models.UpdateChildInput) (*models.Child, *errs.HTTPError) {
	args := m.Called(ctx, childID, child)
	return args.Get(0).(*models.Child), args.Get(1).(*errs.HTTPError)
}

func (m *MockChildRepository) CreateChild(ctx context.Context, child *models.CreateChildInput) (*models.Child, *errs.HTTPError) {
	args := m.Called(ctx, child)
	return args.Get(0).(*models.Child), args.Get(1).(*errs.HTTPError)
}

func (m *MockChildRepository) DeleteChildByID(ctx context.Context, childID uuid.UUID) (*models.Child, *errs.HTTPError) {
	args := m.Called(ctx, childID)
	return args.Get(0).(*models.Child), args.Get(1).(*errs.HTTPError)
}

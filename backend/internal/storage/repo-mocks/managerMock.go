package repomocks

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/jackc/pgx/v5"
)

type MockManagerRepository struct {
	mock.Mock
}

func (m *MockManagerRepository) GetManagerByID(ctx context.Context, id uuid.UUID) (*models.Manager, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*errs.HTTPError)
	}
	return args.Get(0).(*models.Manager), nil
}

func (m *MockManagerRepository) GetManagerByUserID(ctx context.Context, userID uuid.UUID) (*models.Manager, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*errs.HTTPError)
	}
	return args.Get(0).(*models.Manager), nil
}

func (m *MockManagerRepository) GetManagerByOrgID(ctx context.Context, org_id uuid.UUID) (*models.Manager, error) {
	args := m.Called(ctx, org_id)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*errs.HTTPError)
	}
	return args.Get(0).(*models.Manager), nil
}

func (m *MockManagerRepository) CreateManager(ctx context.Context, manager *models.CreateManagerInput) (*models.Manager, error) {
	args := m.Called(ctx, manager)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*errs.HTTPError)
	}
	return args.Get(0).(*models.Manager), nil
}

func (m *MockManagerRepository) DeleteManager(ctx context.Context, id uuid.UUID, tx pgx.Tx) (*models.Manager, error) {
	args := m.Called(ctx, id, tx)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*errs.HTTPError)
	}
	return args.Get(0).(*models.Manager), nil
}

func (m *MockManagerRepository) PatchManager(ctx context.Context, manager *models.PatchManagerInput) (*models.Manager, error) {
	args := m.Called(ctx, manager)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*errs.HTTPError)
	}
	return args.Get(0).(*models.Manager), nil
}

func (m *MockManagerRepository) GetManagerByAuthID(ctx context.Context, authID string) (*models.Manager, error) {
	args := m.Called(ctx, authID)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*errs.HTTPError)
	}
	return args.Get(0).(*models.Manager), nil
}

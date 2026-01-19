package repomocks

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockManagerRepository struct {
	mock.Mock
}

func (m *MockManagerRepository) GetManagerByID(ctx context.Context, id uuid.UUID) (*models.Manager, *errs.HTTPError) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*errs.HTTPError)
	}
	return args.Get(0).(*models.Manager), nil
}

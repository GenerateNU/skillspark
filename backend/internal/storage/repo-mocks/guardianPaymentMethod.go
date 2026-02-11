package repomocks

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockGuardianPaymentMethodRepository struct {
	mock.Mock
}

func (m *MockGuardianPaymentMethodRepository) CreateGuardianPaymentMethod(
	ctx context.Context,
	input *models.CreateGuardianPaymentMethodInput,
) (*models.GuardianPaymentMethod, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*errs.HTTPError)
	}
	return args.Get(0).(*models.GuardianPaymentMethod), nil
}

func (m *MockGuardianPaymentMethodRepository) GetPaymentMethodsByGuardianID(
	ctx context.Context,
	guardianID uuid.UUID,
) ([]models.GuardianPaymentMethod, error) {
	args := m.Called(ctx, guardianID)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*errs.HTTPError)
	}
	return args.Get(0).([]models.GuardianPaymentMethod), nil
}

func (m *MockGuardianPaymentMethodRepository) UpdateGuardianPaymentMethod(
	ctx context.Context,
	id uuid.UUID,
	isDefault bool,
) (*models.GuardianPaymentMethod, error) {
	args := m.Called(ctx, id, isDefault)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*errs.HTTPError)
	}
	return args.Get(0).(*models.GuardianPaymentMethod), nil
}

func (m *MockGuardianPaymentMethodRepository) DeleteGuardianPaymentMethod(
	ctx context.Context,
	id uuid.UUID,
) (*models.GuardianPaymentMethod, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*errs.HTTPError)
	}
	return args.Get(0).(*models.GuardianPaymentMethod), nil
}
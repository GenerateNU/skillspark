package repomocks

import (
	"context"
	"skillspark/internal/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockEmergencyContactRepository struct {
	mock.Mock
}

func (m *MockEmergencyContactRepository) CreateEmergencyContact(ctx context.Context, input *models.CreateEmergencyContactInput) (*models.CreateEmergencyContactOutput, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return args.Get(0).(*models.CreateEmergencyContactOutput), args.Error(1)
}

func (m *MockEmergencyContactRepository) UpdateEmergencyContact(ctx context.Context, input *models.UpdateEmergencyContactInput) (*models.UpdateEmergencyContactOutput, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return args.Get(0).(*models.UpdateEmergencyContactOutput), args.Error(1)
}

func (m *MockEmergencyContactRepository) GetEmergencyContactByGuardianID(ctx context.Context, guardian_id uuid.UUID) ([]*models.EmergencyContact, error) {
	args := m.Called(ctx, guardian_id)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return args.Get(0).([]*models.EmergencyContact), args.Error(1)
}

func (m *MockEmergencyContactRepository) GetEmergencyContactByID(ctx context.Context, id uuid.UUID) (*models.EmergencyContact, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return args.Get(0).(*models.EmergencyContact), args.Error(1)
}

func (m *MockEmergencyContactRepository) DeleteEmergencyContact(ctx context.Context, id uuid.UUID) (*models.DeleteEmergencyContactOutput, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return args.Get(0).(*models.DeleteEmergencyContactOutput), args.Error(1)
}

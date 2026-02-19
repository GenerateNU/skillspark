package repomocks

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/registration"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/mock"
)

type MockRegistrationRepository struct {
	mock.Mock
}

func (m *MockRegistrationRepository) CreateRegistration(ctx context.Context, input *models.CreateRegistrationInput) (*models.CreateRegistrationOutput, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.CreateRegistrationOutput), args.Error(1)
}

func (m *MockRegistrationRepository) GetRegistrationByID(ctx context.Context, input *models.GetRegistrationByIDInput, tx *pgx.Tx) (*models.GetRegistrationByIDOutput, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.GetRegistrationByIDOutput), args.Error(1)
}

func (m *MockRegistrationRepository) GetRegistrationsByChildID(ctx context.Context, input *models.GetRegistrationsByChildIDInput) (*models.GetRegistrationsByChildIDOutput, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.GetRegistrationsByChildIDOutput), args.Error(1)
}

func (m *MockRegistrationRepository) GetRegistrationsByGuardianID(ctx context.Context, input *models.GetRegistrationsByGuardianIDInput) (*models.GetRegistrationsByGuardianIDOutput, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.GetRegistrationsByGuardianIDOutput), args.Error(1)
}

func (m *MockRegistrationRepository) UpdateRegistration(ctx context.Context, input *models.UpdateRegistrationInput) (*models.UpdateRegistrationOutput, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.UpdateRegistrationOutput), args.Error(1)
}

func (m *MockRegistrationRepository) GetRegistrationsByEventOccurrenceID(ctx context.Context, input *models.GetRegistrationsByEventOccurrenceIDInput) (*models.GetRegistrationsByEventOccurrenceIDOutput, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.GetRegistrationsByEventOccurrenceIDOutput), args.Error(1)
}

func (m *MockRegistrationRepository) GetUpcomingUnsentRegistrations(ctx context.Context, input *registration.GetUpcomingUnsentRegistrationsInput) (*registration.GetUpcomingUnsentRegistrationsOutput, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*registration.GetUpcomingUnsentRegistrationsOutput), args.Error(1)
}

func (m *MockRegistrationRepository) MarkReminderSent(ctx context.Context, tx pgx.Tx, id uuid.UUID, sent bool) error {
	args := m.Called(ctx, tx, id, sent)
	return args.Error(0)
}

func (m *MockEventOccurrenceRepository) DeleteEventOccurrence(
	ctx context.Context,
	id uuid.UUID,
) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

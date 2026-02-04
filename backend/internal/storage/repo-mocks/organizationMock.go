package repomocks

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/utils"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockOrganizationRepository struct {
	mock.Mock
}

func (m *MockOrganizationRepository) CreateOrganization(ctx context.Context, input *models.CreateOrganizationInput, PfpS3Key *string) (*models.Organization, error) {
	args := m.Called(ctx, input, PfpS3Key)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*errs.HTTPError)
	}
	return args.Get(0).(*models.Organization), nil
}

func (m *MockOrganizationRepository) GetOrganizationByID(ctx context.Context, id uuid.UUID) (*models.Organization, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*errs.HTTPError)
	}
	return args.Get(0).(*models.Organization), nil
}

func (m *MockOrganizationRepository) GetAllOrganizations(ctx context.Context, pagination utils.Pagination) ([]models.Organization, error) {
	args := m.Called(ctx, pagination)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*errs.HTTPError)
	}
	return args.Get(0).([]models.Organization), nil
}

func (m *MockOrganizationRepository) UpdateOrganization(ctx context.Context, input *models.UpdateOrganizationInput, PfpS3Key *string) (*models.Organization, error) {
	args := m.Called(ctx, input, PfpS3Key)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*errs.HTTPError)
	}
	return args.Get(0).(*models.Organization), nil
}

func (m *MockOrganizationRepository) DeleteOrganization(ctx context.Context, id uuid.UUID) (*models.Organization, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(*errs.HTTPError)
	}
	return args.Get(0).(*models.Organization), nil
}

func (m *MockOrganizationRepository) GetEventOccurrencesByOrganizationID(ctx context.Context, organization_id uuid.UUID) ([]models.EventOccurrence, error) {
	args := m.Called(ctx, organization_id)
	eventOccurrences := args.Get(0)
	if eventOccurrences == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return eventOccurrences.([]models.EventOccurrence), nil
}

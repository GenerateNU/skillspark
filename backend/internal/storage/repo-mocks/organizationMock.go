package repomocks

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/utils"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockOrganizationRepository struct {
	mock.Mock
}

func (m *MockOrganizationRepository) CreateOrganization(ctx context.Context, input *models.CreateOrganizationDBInput, PfpS3Key *string) (*models.Organization, error) {
	args := m.Called(ctx, input, PfpS3Key)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return args.Get(0).(*models.Organization), nil
}

func (m *MockOrganizationRepository) GetOrganizationByID(ctx context.Context, id uuid.UUID, acceptLanguage string) (*models.Organization, error) {
	args := m.Called(ctx, id, acceptLanguage)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return args.Get(0).(*models.Organization), nil
}

func (m *MockOrganizationRepository) GetAllOrganizations(ctx context.Context, pagination utils.Pagination, acceptLanguage string) ([]models.Organization, error) {
	args := m.Called(ctx, pagination, acceptLanguage)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return args.Get(0).([]models.Organization), nil
}

func (m *MockOrganizationRepository) UpdateOrganization(ctx context.Context, input *models.UpdateOrganizationDBInput, PfpS3Key *string) (*models.Organization, error) {
	args := m.Called(ctx, input, PfpS3Key)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return args.Get(0).(*models.Organization), nil
}

func (m *MockOrganizationRepository) DeleteOrganization(ctx context.Context, id uuid.UUID, acceptLanguage string) (*models.Organization, error) {
	args := m.Called(ctx, id, acceptLanguage)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return args.Get(0).(*models.Organization), nil
}

func (m *MockOrganizationRepository) GetEventOccurrencesByOrganizationID(ctx context.Context, organization_id uuid.UUID, acceptLanguage string) ([]models.EventOccurrence, error) {
	args := m.Called(ctx, organization_id, acceptLanguage)
	eventOccurrences := args.Get(0)
	if eventOccurrences == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return eventOccurrences.([]models.EventOccurrence), nil
}

func (m *MockOrganizationRepository) SetStripeAccountID(ctx context.Context, orgID uuid.UUID, stripeAccountID string) (*models.Organization, error) {
	args := m.Called(ctx, orgID, stripeAccountID)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return args.Get(0).(*models.Organization), nil
}

func (m *MockOrganizationRepository) SetStripeAccountStatus(ctx context.Context, stripeAccountID string, activated bool) (*models.Organization, error) {
	args := m.Called(ctx, stripeAccountID, activated)
	if args.Get(0) == nil {
		if args.Get(1) == nil {
			return nil, nil
		}
		return nil, args.Get(1).(error)
	}
	return args.Get(0).(*models.Organization), nil
}

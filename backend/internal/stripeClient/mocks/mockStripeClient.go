package stripemocks

import (
	"context"
	"skillspark/internal/models"
	"github.com/stretchr/testify/mock"
	"github.com/stripe/stripe-go/v84"
)

type MockStripeClient struct {
	mock.Mock
}

func (m *MockStripeClient) CreateOrganizationAccount(
	ctx context.Context,
	name string,
	email string,
	country string,
) (*models.CreateOrgStripeAccountClientOutput, error) {
	args := m.Called(ctx, name, email, country)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.CreateOrgStripeAccountClientOutput), args.Error(1)
}

func (m *MockStripeClient) CreateAccountOnboardingLink(
	ctx context.Context,
	input *models.CreateStripeOnboardingLinkClientInput,
) (*models.CreateStripeOnboardingLinkOutput, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.CreateStripeOnboardingLinkOutput), args.Error(1)
}

func (m *MockStripeClient) CreateCustomer(
	ctx context.Context,
	email string,
	name string,
) (*stripe.Customer, error) {
	args := m.Called(ctx, email, name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*stripe.Customer), args.Error(1)
}

func (m *MockStripeClient) CreateSetupIntent(
	ctx context.Context,
	stripeCustomerID string,
) (string, error) {
	args := m.Called(ctx, stripeCustomerID)
	return args.String(0), args.Error(1)
}

func (m *MockStripeClient) CreatePaymentIntent(
	ctx context.Context,
	input *models.CreatePaymentIntentInput,
) (*models.CreatePaymentIntentOutput, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.CreatePaymentIntentOutput), args.Error(1)
}

func (m *MockStripeClient) CapturePaymentIntent(
	ctx context.Context,
	input *models.CapturePaymentIntentInput,
) (*models.CapturePaymentIntentOutput, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.CapturePaymentIntentOutput), args.Error(1)
}

func (m *MockStripeClient) CancelPaymentIntent(
	ctx context.Context,
	input *models.CancelPaymentIntentInput,
) (*models.CancelPaymentIntentOutput, error) {
	args := m.Called(ctx, input)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.CancelPaymentIntentOutput), args.Error(1)
}

func (m *MockStripeClient) GetAccount(
	ctx context.Context,
	accountID string,
) (*stripe.V2CoreAccount, error) {
	args := m.Called(ctx, accountID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*stripe.V2CoreAccount), args.Error(1)
}

func (m *MockStripeClient) DetachPaymentMethod(
	ctx context.Context,
	paymentMethodID string,
) error {
	args := m.Called(ctx, paymentMethodID)
	return args.Error(0)
}

func (m *MockStripeClient) CreateLoginLink(
	ctx context.Context,
	accountID string,
) (string, error) {
	args := m.Called(ctx, accountID)
	return args.String(0), args.Error(1)
}
package stripeClient

import (
	"context"
	"skillspark/internal/models"

	"github.com/stripe/stripe-go/v84"
)

type StripeClientInterface interface {
	CreateOrganizationAccount(ctx context.Context, name string, email string, country string) (*models.CreateOrgStripeAccountClientOutput, error)
	CreateAccountOnboardingLink(ctx context.Context, input *models.CreateStripeOnboardingLinkClientInput) (*models.CreateStripeOnboardingLinkOutput, error)
	CreateCustomer(ctx context.Context, email string, name string) (*stripe.Customer, error)
	CreateSetupIntent(ctx context.Context, stripeCustomerID string) (string, error)
	CreatePaymentIntent(ctx context.Context, input *models.CreatePaymentIntentInput) (*models.CreatePaymentIntentOutput, error)
	GetAccount(ctx context.Context, accountID string) (*stripe.V2CoreAccount, error)
	DetachPaymentMethod(ctx context.Context, paymentMethodID string) error
	CreateLoginLink(ctx context.Context, accountID string) (string, error)
	CancelPaymentIntent(ctx context.Context, input *models.CancelPaymentIntentInput) (*models.CancelPaymentIntentOutput, error)
	CapturePaymentIntent(ctx context.Context, input *models.CapturePaymentIntentInput) (*models.CapturePaymentIntentOutput, error)
}
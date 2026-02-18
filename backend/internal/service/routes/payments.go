package routes

import (
	"context"
	"net/http"
	"skillspark/internal/models"
	"skillspark/internal/service/handler/payment"
	"skillspark/internal/storage"
	"skillspark/internal/stripeClient"

	"github.com/danielgtaylor/huma/v2"
)

func SetupPaymentRoutes(api huma.API, repo *storage.Repository, sc stripeClient.StripeClientInterface) {
	paymentHandler := payment.NewHandler(
		repo.Organization,
		repo.Manager,
		repo.Registration,
		repo.Location,
		repo.Guardian,
		sc,
	)

	huma.Register(api, huma.Operation{
		OperationID:   "create-org-stripe-account",
		Method:        http.MethodPost,
		Path:          "/api/v1/stripe/orgaccount",
		Summary:       "Create a new Stripe account for an organization",
		Description:   "Create a new Stripe account for an organization",
		Tags:          []string{"Payments"},
	}, func(ctx context.Context, input *models.CreateOrgStripeAccountInput) (*models.CreateOrgStripeAccountOutput, error) {
		return paymentHandler.CreateOrgStripeAccount(ctx, input)
	})

	huma.Register(api, huma.Operation{
	OperationID:   "create-org-stripe-onboarding-link",
	Method:        http.MethodPost,
	Path:          "/api/v1/stripe/onboarding/{organization_id}",
	Summary:       "Creates an onboarding link for a Stripe account",
	Description:   "Creates an onboarding link for a Stripe account",
	Tags:          []string{"Payments"},
}, func(ctx context.Context, input *models.CreateStripeOnboardingLinkInput) (*models.CreateStripeOnboardingLinkOutput, error) {
	return paymentHandler.CreateAccountOnboardingLink(ctx, input)
})

	huma.Register(api, huma.Operation{
		OperationID:   "create-org-login-link",
		Method:        http.MethodPost,
		Path:          "/api/v1/stripe/login/{organization_id}",
		Summary:       "Create Stripe dashboard login link for organization",
		Description:   "Generates a login link for organization to access their Stripe Express dashboard",
		Tags:          []string{"Payments"},
	}, func(ctx context.Context, input *models.CreateOrgLoginLinkInput) (*models.CreateOrgLoginLinkOutput, error) {
		return paymentHandler.CreateOrgLoginLink(ctx, input)
	})

	huma.Register(api, huma.Operation{
	OperationID:   "create-guardian-setup-intent",
	Method:        http.MethodPost,
	Path:          "/api/v1/stripe/setup-intent/{guardian_id}",
	Summary:       "Create a SetupIntent for guardian to add payment method",
	Description:   "Creates a Stripe SetupIntent and returns client_secret for frontend to collect card details",
	Tags:          []string{"Payments"},
	}, func(ctx context.Context, input *models.CreateSetupIntentInput) (*models.CreateSetupIntentOutput, error) {
	return paymentHandler.CreateSetupIntent(ctx, input)
	})

	huma.Register(api, huma.Operation{
		OperationID:   "create-stripe-customer",
		Method:        http.MethodPost,
		Path:          "/api/v1/stripe/customer/{guardian_id}",
		Summary:       "Create a new Stripe Customer",
		Description:   "Create a new Stripe Customer",
		Tags:          []string{"Payments"},
	}, func(ctx context.Context, input *models.CreateStripeCustomerInput) (*models.CreateStripeCustomerOutput, error) {
		return paymentHandler.CreateStripeCustomer(ctx, input)
	})

	// huma.Register(api, huma.Operation{
	// 	OperationID: "stripe-webhook",
	// 	Method:      http.MethodPost,
	// 	Path:        "/webhooks/stripe",
	// 	Summary:     "Handle Stripe webhooks",
	// 	Tags:        []string{"Webhooks"},
	// 	Security:    []map[string][]string{},
	// }, paymentHandler.HandleStripeWebhook)
}
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

func SetupPaymentRoutes(api huma.API, repo *storage.Repository, sc *stripeClient.StripeClient) {
	paymentHandler := payment.NewHandler(repo.Organization, repo.Manager, repo.Registration, repo.Location, *sc)


	huma.Register(api, huma.Operation{
		OperationID:   "create-org-stripe-account",
		Method:        http.MethodPost,
		Path:          "/api/v1/stripe/orgaccount",
		Summary:       "Create a new Stripe account for an organization",
		Description:   "Create a new Stripe account for an organization",
		Tags:          []string{"StripeClient"},
	}, func(ctx context.Context, input *models.CreateOrgStripeAccountInput) (*models.CreateOrgStripeAccountOutput, error) {
		return paymentHandler.CreateOrgStripeAccount(ctx, input);
	})

	huma.Register(api, huma.Operation{
		OperationID:   "create-org-stripe-onboarding-link",
		Method:        http.MethodPost,
		Path:          "/api/v1/stripe/onboarding",
		Summary:       "Creates an onboarding link for a Stripe account",
		Description:   "Creates an onboarding link for a Stripe account",
		Tags:          []string{"StripeClient"},
	}, func(ctx context.Context, input *models.CreateStripeOnboardingLinkInput) (*models.CreateStripeOnboardingLinkOutput, error) {
		return sc.CreateAccountOnboardingLink(ctx, input);
	})
}
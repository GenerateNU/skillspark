package models

import (
	"time"

	// "github.com/google/uuid"
	"github.com/google/uuid"
	"github.com/stripe/stripe-go/v84"
)

type CreateOrgStripeAccountInput struct {
	Body struct {
		OrganizationID uuid.UUID `json:"organization_id" doc:"UUID of the existing organization"`
	}
}

type CreateOrgStripeAccountOutput struct {
	Body struct {
		Account stripe.V2CoreAccount `json:"account" doc:"Stripe account details"`
	}
}

type CreateStripeOnboardingLinkInput struct {
	Body struct {
		AccountID  string `json:"account_id" doc:"Stripe account ID (e.g., acct_123)"`
		RefreshURL string `json:"refresh_url" doc:"URL to redirect if onboarding is exited early"`
		ReturnURL  string `json:"return_url" doc:"URL to redirect after successful onboarding"`
	}
}

type CreateStripeOnboardingLinkOutput struct {
	Body struct {
		OnboardingURL string `json:"onboarding_url" doc:"Stripe-hosted onboarding page URL"`
	}
}

type CreatePaymentIntentInput struct {
	Body struct {
		// From frontend
		RegistrationID   uuid.UUID  `json:"registration_id"`
		GuardianID       uuid.UUID  `json:"guardian_id"` 
		ProviderOrgID    uuid.UUID  `json:"provider_org_id"`
		Amount           int64      `json:"amount"`
		Currency         string     `json:"currency"`
		EventDate        time.Time  `json:"event_date"`
		PaymentMethodID  *string    `json:"payment_method_id,omitempty"`
		
		// Populated by handler from DB (not from frontend!)
		GuardianStripeID string
		OrgStripeID      string
	}
}

type CreatePaymentIntentOutput struct {
	Body struct {
		PaymentIntentID string `json:"payment_intent_id" doc:"Stripe payment intent ID"`
		ClientSecret    string `json:"client_secret" doc:"Client secret for frontend to confirm payment"`
		Status          string `json:"status" doc:"Payment intent status (e.g., requires_confirmation)"`
	}
}


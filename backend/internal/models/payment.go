package models

import (
	"time"

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

type CreateSetupIntentInput struct {
	GuardianID uuid.UUID `path:"guardian_id" doc:"Guardian ID"`
}

type CreateSetupIntentOutput struct {
	Body struct {
		ClientSecret string `json:"client_secret" doc:"Stripe SetupIntent client_secret for frontend"`
	}
}

type CreateOrgLoginLinkInput struct {
	OrganizationID uuid.UUID `path:"organization_id" doc:"Organization ID"`
}

type CreateOrgLoginLinkOutput struct {
	Body struct {
		LoginURL string `json:"login_url" doc:"Stripe Express dashboard login URL"`
	}
}

type CreatePaymentIntentInput struct {
	Body struct {
		RegistrationID  uuid.UUID `json:"registration_id" doc:"Registration/booking ID"`
		GuardianID      uuid.UUID `json:"guardian_id" doc:"Guardian ID"`
		ProviderOrgID   uuid.UUID `json:"provider_org_id" doc:"Provider organization ID"`
		Amount          int64     `json:"amount" doc:"Total amount in cents" minimum:"1"`
		Currency        string    `json:"currency" doc:"Currency code (e.g., thb, usd)" pattern:"^[a-z]{3}$"`
		EventDate       time.Time `json:"event_date" doc:"Event date and time"`
		PaymentMethodID *string   `json:"payment_method_id,omitempty" doc:"Stripe payment method ID (required for bookings)"`
		
		GuardianStripeID string
		OrgStripeID      string
	}
}

type CreatePaymentIntentOutput struct {
	Body struct {
		PaymentIntentID string `json:"payment_intent_id" doc:"Stripe payment intent ID"`
		ClientSecret    string `json:"client_secret" doc:"Client secret for frontend to confirm payment"`
		Status          string `json:"status" doc:"Payment intent status"`
	}
}
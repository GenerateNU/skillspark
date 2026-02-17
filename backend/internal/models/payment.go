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
		Account Organization `json:"account" doc:"Stripe account details"`
	} `json:"body" doc:"Updated organization with Stripe account ID"`
}

type CreateOrgStripeAccountClientOutput struct {
	Body struct {
		Account stripe.V2CoreAccount `json:"account" doc:"Stripe account details"`
	} `json:"body" doc:"Updated organization with Stripe account ID"`
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
	} `json:"body"`
}

type CreateSetupIntentInput struct {
	GuardianID uuid.UUID `path:"guardian_id" doc:"Guardian ID"`
}

type CreateSetupIntentOutput struct {
	Body struct {
		ClientSecret string `json:"client_secret" doc:"Stripe SetupIntent client_secret for frontend"`
	} `json:"body"`
}

type CreateOrgLoginLinkInput struct {
	OrganizationID uuid.UUID `path:"organization_id" doc:"Organization ID"`
}

type CreateOrgLoginLinkOutput struct {
	Body struct {
		LoginURL string `json:"login_url" doc:"Stripe Express dashboard login URL"`
	} `json:"body"`
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
		PaymentIntentID   string `json:"payment_intent_id" doc:"Stripe payment intent ID"`
		ClientSecret      string `json:"client_secret" doc:"Client secret for frontend to confirm payment"`
		Status            string `json:"status" doc:"Payment intent status"`
		TotalAmount       int    `json:"total_amount" doc:"Total amount in cents"`
		ProviderAmount    int    `json:"provider_amount" doc:"Amount provider receives in cents"`
		PlatformFeeAmount int    `json:"platform_fee_amount" doc:"Platform fee in cents"`
		Currency          string `json:"currency" doc:"Currency code"`
	} `json:"body"`
}

type CancelPaymentIntentInput struct {
	PaymentIntentID string `json:"payment_intent_id" doc:"Stripe payment intent ID to cancel/refund"`
	StripeAccountID string `json:"stripe_account_id" doc:"Organization's Stripe account ID"`
}

type CancelPaymentIntentOutput struct {
	Body struct {
		PaymentIntentID string `json:"payment_intent_id" doc:"Cancelled payment intent ID"`
		Status          string `json:"status" doc:"Payment intent status after cancellation"`
		Amount          int64  `json:"amount" doc:"Amount that was cancelled/refunded in cents"`
		Currency        string `json:"currency" doc:"Currency code"`
	} `json:"body" doc:"Cancellation result"`
}

type CapturePaymentIntentInput struct {
	PaymentIntentID string `json:"payment_intent_id" doc:"Stripe payment intent ID to capture"`
	StripeAccountID string `json:"stripe_account_id" doc:"Organization's Stripe account ID"`
}

type CapturePaymentIntentOutput struct {
	Body struct {
		PaymentIntentID string `json:"payment_intent_id" doc:"Captured payment intent ID"`
		Status          string `json:"status" doc:"Payment intent status (should be 'succeeded')"`
		Amount          int64  `json:"amount" doc:"Amount captured in cents"`
		Currency        string `json:"currency" doc:"Currency code"`
	} `json:"body" doc:"Capture result"`
}
package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/stripe/stripe-go/v84"
)

type CreateOrgStripeAccountInput struct {
	OrganizationID uuid.UUID `path:"organization_id" doc:"UUID of the existing organization"`
}

type CreateOrgStripeAccountOutput struct {
	Body struct {
		Account Organization `json:"account" doc:"Stripe account details"`
	} `json:"body" doc:"Updated organization with Stripe account ID"`
}

type CreateOrgStripeAccountClientOutput struct {
	Body struct {
		Account stripe.V2CoreAccount
	}
}

type CreateStripeOnboardingLinkInput struct {
	OrganizationID uuid.UUID `path:"organization_id" doc:"Organization ID"`
	Body           struct {
		RefreshURL string `json:"refresh_url" doc:"URL to redirect if onboarding is exited early"`
		ReturnURL  string `json:"return_url" doc:"URL to redirect after successful onboarding"`
	} `json:"body"`
}

type CreateStripeOnboardingLinkClientInput struct {
	Body struct {
		StripeAccountID string
		RefreshURL      string
		ReturnURL       string
	}
}

type GetPaymentMethodsByGuardianIDInput struct {
	GuardianID uuid.UUID `path:"guardian_id" doc:"Guardian ID"`
}

type PaymentMethodCard struct {
	Brand    string `json:"brand"`
	Last4    string `json:"last4"`
	ExpMonth int64  `json:"exp_month"`
	ExpYear  int64  `json:"exp_year"`
}

type PaymentMethod struct {
	ID   string            `json:"id"`
	Type string            `json:"type"`
	Card PaymentMethodCard `json:"card"`
}

type GetPaymentMethodsByGuardianIDOutput struct {
	Body struct {
		PaymentMethods []PaymentMethod `json:"payment_methods"`
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

type CreateCustomerPaymentMethodInput struct {
	GuardianID uuid.UUID `path:"guardian_id"`
}

type CreateCustomerPaymentMethodOutput struct {
	PaymentMethod stripe.PaymentMethod `json:"payment_method_id"`
}

type CreatePaymentIntentInput struct {
	Body struct {
		RegistrationID   uuid.UUID `json:"registration_id" doc:"Registration/booking ID"`
		GuardianID       uuid.UUID `json:"guardian_id" doc:"Guardian ID"`
		ProviderOrgID    uuid.UUID `json:"provider_org_id" doc:"Provider organization ID"`
		Amount           int64     `json:"amount" doc:"Total amount in cents" minimum:"1"` // Stripe requires int64
		Currency         string    `json:"currency" doc:"Currency code (e.g., thb, usd)" pattern:"^[a-z]{3}$"`
		EventDate        time.Time `json:"event_date" doc:"Event date and time"`
		PaymentMethodID  string    `json:"payment_method_id,omitempty" doc:"Stripe payment method ID (required for bookings)"`
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
}

type CapturePaymentIntentOutput struct {
	Body struct {
		PaymentIntentID string `json:"payment_intent_id" doc:"Captured payment intent ID"`
		Status          string `json:"status" doc:"Payment intent status (should be 'succeeded')"`
		Amount          int64  `json:"amount" doc:"Amount captured in cents"`
		Currency        string `json:"currency" doc:"Currency code"`
	} `json:"body" doc:"Capture result"`
}

type CancelRegistrationWithPaymentInput struct {
	ID                  uuid.UUID          `json:"id"`
	CancelledAt         time.Time          `json:"cancelled_at"`
	Status              RegistrationStatus `json:"status"`
	PaymentIntentStatus string             `json:"payment_intent_status"`
}

type DetachPaymentMethodInput struct {
	Body struct {
		PaymentMethodID string    `json:"payment_method_id" doc:"Payment Method ID"`
		GuardianID      uuid.UUID `json:"guardian_id" doc:"Guardian ID"`
	} `json:"body"`
}

type GetRegistrationByPaymentIntentIDOutput struct {
	ID                    uuid.UUID          `json:"id"`
	EventOccurrenceID     uuid.UUID          `json:"event_occurrence_id"`
	GuardianID            uuid.UUID          `json:"guardian_id"`
	ChildID               uuid.UUID          `json:"child_id"`
	Status                RegistrationStatus `json:"status"`
	StripePaymentIntentID string             `json:"stripe_payment_intent_id"`
	OrgStripeAccountID    string             `json:"org_stripe_account_id"`
}

type RefundPaymentInput struct {
	PaymentIntentID string `json:"stripe_payment_intent_id"`
}

type RefundPaymentOutput struct {
	Body struct {
		RefundID string `json:"refund_id"`
		Status   string `json:"status"`
		Amount   int64  `json:"amount"`
		Currency string `json:"currency"`
	}
}

type AttachPaymentMethodInput struct {
	GuardianID uuid.UUID `path:"guardian_id" doc:"Guardian ID"`
	Body       struct {
		PaymentMethodID string `json:"payment_method_id" doc:"Stripe payment method ID to attach (e.g. pm_card_visa)"`
	}
}

type AttachPaymentMethodOutput struct {
	Body struct {
		PaymentMethodID string `path:"payment_method_id" doc:"Attached payment method ID"`
		CustomerID      string `json:"customer_id" doc:"Stripe customer ID the payment method was attached to"`
	}
}

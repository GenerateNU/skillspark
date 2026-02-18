package models

import (
	"time"

	"github.com/google/uuid"
)

type Registration struct {
	ID                    uuid.UUID          `json:"id" db:"id" doc:"Unique registration identifier"`
	ChildID               uuid.UUID          `json:"child_id" db:"child_id" doc:"ID of the registered child"`
	GuardianID            uuid.UUID          `json:"guardian_id" db:"guardian_id" doc:"ID of the child's guardian"`
	EventOccurrenceID     uuid.UUID          `json:"event_occurrence_id" db:"event_occurrence_id" doc:"ID of the event occurrence"`
	Status                RegistrationStatus `json:"status" db:"status" doc:"Current status of the registration" enum:"registered,cancelled"`
	StripePaymentIntentID string             `json:"stripe_payment_intent_id" db:"stripe_payment_intent_id" doc:"Stripe payment intent ID"`
	StripeCustomerID      string             `json:"stripe_customer_id" db:"stripe_customer_id" doc:"Stripe customer ID"`
	OrgStripeAccountID    string             `json:"org_stripe_account_id" db:"org_stripe_account_id" doc:"Organization's Stripe account ID"`
	StripePaymentMethodID string             `json:"stripe_payment_method_id" db:"stripe_payment_method_id" doc:"Stripe payment method ID"`
	TotalAmount           int                `json:"total_amount" db:"total_amount" doc:"Total amount in cents"`
	ProviderAmount        int                `json:"provider_amount" db:"provider_amount" doc:"Amount provider receives in cents"`
	PlatformFeeAmount     int                `json:"platform_fee_amount" db:"platform_fee_amount" doc:"Platform fee amount in cents"`
	Currency              string             `json:"currency" db:"currency" doc:"Currency code (e.g., thb, usd)"`
	PaymentIntentStatus   string             `json:"payment_intent_status" db:"payment_intent_status" doc:"Stripe payment intent status"`
	PaidAt                *time.Time         `json:"paid_at,omitempty" db:"paid_at" doc:"Timestamp when payment was completed"`
	CancelledAt           *time.Time         `json:"cancelled_at,omitempty" db:"cancelled_at" doc:"Timestamp when registration was cancelled"`
	CreatedAt             time.Time          `json:"created_at" db:"created_at" doc:"Timestamp when registration was created"`
	UpdatedAt             time.Time          `json:"updated_at" db:"updated_at" doc:"Timestamp when registration was last updated"`
	EventName             string             `json:"event_name" db:"event_name" doc:"Name of the event"`
	OccurrenceStartTime   time.Time          `json:"occurrence_start_time" db:"occurrence_start_time" doc:"Start time of the event occurrence"`
}

type RegistrationStatus string

const (
	RegistrationStatusRegistered RegistrationStatus = "registered"
	RegistrationStatusCancelled  RegistrationStatus = "cancelled"
)

func (rs RegistrationStatus) IsValid() bool {
	return rs == RegistrationStatusRegistered || rs == RegistrationStatusCancelled
}

type CreateRegistrationInput struct {
	Body struct {
		ChildID           uuid.UUID          `json:"child_id" doc:"ID of the child to register" format:"uuid" required:"true"`
		GuardianID        uuid.UUID          `json:"guardian_id" doc:"ID of the guardian registering the child" format:"uuid" required:"true"`
		EventOccurrenceID uuid.UUID          `json:"event_occurrence_id" doc:"ID of the event occurrence to register for" format:"uuid" required:"true"`
		PaymentMethodID   *string            `json:"payment_method_id" doc:"Stripe payment method ID to use"`
		Currency          string             `json:"currency" doc:"Currency code (e.g., thb, usd)" default:"thb"`
		Status            RegistrationStatus `json:"status" doc:"Initial status of the registration" default:"registered" enum:"registered,cancelled"`
	} `json:"body"`
}

type CreateRegistrationWithPaymentData struct {
	ChildID                  uuid.UUID
	GuardianID               uuid.UUID
	EventOccurrenceID        uuid.UUID
	Status                   RegistrationStatus
	StripePaymentIntentID    string
	StripeCustomerID         string
	OrgStripeAccountID       string
	StripePaymentMethodID    string
	TotalAmount              int
	ProviderAmount           int
	PlatformFeeAmount        int
	Currency                 string
	PaymentIntentStatus      string
}

type CreateRegistrationOutput struct {
	Body Registration `json:"body" doc:"The newly created registration with full details"`
}

type UpdateRegistrationInput struct {
	ID   uuid.UUID `path:"id" format:"uuid" doc:"Registration ID to update"`
	Body struct {
		ChildID uuid.UUID `json:"child_id,omitempty" doc:"Updated child ID (optional)"`
	} `json:"body"`
}

type UpdateRegistrationOutput struct {
	Body Registration `json:"body" doc:"The updated registration with full details"`
}

type CancelRegistrationInput struct {
	ID uuid.UUID `path:"id" format:"uuid" doc:"Registration ID to cancel"`
}

type CancelRegistrationOutput struct {
	Body struct {
		Message      string             `json:"message" doc:"Success message"`
		RefundStatus string             `json:"refund_status,omitempty" doc:"Refund status if applicable"`
		Registration Registration `json:"registration" doc:"Updated registration"`
	} `json:"body"`
}

type UpdateRegistrationPaymentStatusInput struct {
	ID uuid.UUID `path:"id" format:"uuid" doc:"Registration ID"`
	Body struct {
		PaymentIntentStatus string     `json:"payment_intent_status" doc:"New payment intent status from Stripe"`
	} `json:"body"`
}

type UpdateRegistrationPaymentStatusOutput struct {
	Body Registration `json:"body" doc:"The updated registration with full details"`
}

type GetRegistrationByIDInput struct {
	ID uuid.UUID `path:"id" format:"uuid" doc:"Registration ID to retrieve" required:"true"`
}

type GetRegistrationByIDOutput struct {
	Body Registration `json:"body" doc:"The requested registration with full details"`
}

type GetRegistrationsByChildIDInput struct {
	ChildID uuid.UUID `path:"child_id" format:"uuid" doc:"Child ID to retrieve registrations for" required:"true"`
}

type GetRegistrationsByChildIDOutput struct {
	Body struct {
		Registrations []Registration `json:"registrations" doc:"List of registrations for the child"`
	} `json:"body"`
}

type GetRegistrationsByGuardianIDInput struct {
	GuardianID uuid.UUID `path:"guardian_id" format:"uuid" doc:"Guardian ID to retrieve registrations for" required:"true"`
}

type GetRegistrationsByGuardianIDOutput struct {
	Body struct {
		Registrations []Registration `json:"registrations" doc:"List of registrations for the guardian"`
	} `json:"body"`
}

type GetRegistrationsByEventOccurrenceIDInput struct {
	EventOccurrenceID uuid.UUID `path:"event_occurrence_id" format:"uuid" doc:"Event Occurrence ID to retrieve registrations for" required:"true"`
}

type GetRegistrationsByEventOccurrenceIDOutput struct {
	Body struct {
		Registrations []Registration `json:"registrations" doc:"List of registrations for the event occurrence"`
	} `json:"body"`
}

type DeleteRegistrationInput struct {
	ID uuid.UUID `path:"id" format:"uuid" doc:"Registration ID to delete" required:"true"`
}

type DeleteRegistrationOutput struct {
	Body Registration `json:"body" doc:"The deleted registration details"`
}

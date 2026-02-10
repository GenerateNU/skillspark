package models

import (
	"time"

	"github.com/google/uuid"
)

type GuardianPaymentMethod struct {
	ID                    uuid.UUID  `json:"id" db:"id"`
	GuardianID            uuid.UUID  `json:"guardian_id" db:"guardian_id"`
	StripePaymentMethodID string     `json:"stripe_payment_method_id" db:"stripe_payment_method_id"`
	CardBrand             *string    `json:"card_brand,omitempty" db:"card_brand"`
	CardLast4             *string    `json:"card_last4,omitempty" db:"card_last4"`
	CardExpMonth          *int       `json:"card_exp_month,omitempty" db:"card_exp_month"`
	CardExpYear           *int       `json:"card_exp_year,omitempty" db:"card_exp_year"`
	IsDefault             bool       `json:"is_default" db:"is_default"`
	CreatedAt             time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at" db:"updated_at"`
}

type CreateGuardianPaymentMethodInput struct {
	Body struct {
		GuardianID            uuid.UUID `json:"guardian_id" doc:"Guardian ID"`
		StripePaymentMethodID string    `json:"stripe_payment_method_id" doc:"Stripe payment method ID (pm_...)"`
		CardBrand             *string   `json:"card_brand,omitempty" doc:"Card brand (visa, mastercard, etc.)"`
		CardLast4             *string   `json:"card_last4,omitempty" doc:"Last 4 digits of card"`
		CardExpMonth          *int      `json:"card_exp_month,omitempty" doc:"Card expiration month" minimum:"1" maximum:"12"`
		CardExpYear           *int      `json:"card_exp_year,omitempty" doc:"Card expiration year" minimum:"2026"`
		IsDefault             bool      `json:"is_default" doc:"Whether this is the default payment method"`
	}
}

type CreateGuardianPaymentMethodOutput struct {
	Body GuardianPaymentMethod `json:"body" doc:"Created payment method"`
}

type GetGuardianPaymentMethodsByGuardianIDInput struct {
	GuardianID uuid.UUID `path:"guardian_id" doc:"Guardian ID"`
}

type GetGuardianPaymentMethodsByGuardianIDOutput struct {
	Body []GuardianPaymentMethod `json:"body" doc:"List of guardian's payment methods"`
}

type DeleteGuardianPaymentMethodInput struct {
	ID uuid.UUID `path:"id" doc:"Payment method ID"`
}

type DeleteGuardianPaymentMethodOutput struct {
	Body struct {
		Message string `json:"message" doc:"Success message"`
	}
}

type SetDefaultPaymentMethodInput struct {
	GuardianID        uuid.UUID `path:"guardian_id" doc:"Guardian ID"`
	PaymentMethodID   uuid.UUID `path:"payment_method_id" doc:"Payment method ID to set as default"`
}

type SetDefaultPaymentMethodOutput struct {
	Body GuardianPaymentMethod `json:"body" doc:"Updated payment method"`
}
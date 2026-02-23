package webhook

import (
	"context"
	"log"

	"skillspark/internal/models"

	"github.com/stripe/stripe-go/v84"
)

func (h *Handler) handlePaymentIntentFailed(ctx context.Context, event stripe.Event) error {
	pi, err := unmarshalEvent[stripe.PaymentIntent](event)
	if err != nil {
		log.Printf("Failed to unmarshal payment_intent.payment_failed: %v", err)
		return nil // return nil so Stripe gets 200 and doesn't retry
	}

	registration, err := h.repo.Registration.GetRegistrationByPaymentIntentID(ctx, pi.ID)
	if err != nil {
		log.Printf("Registration not found for payment intent %s: %v", pi.ID, err)
		return nil
	}

	cancelledStatus := models.RegistrationStatusCancelled
	piStatus := string(pi.Status)
	input := &models.CancelRegistrationInput{
		ID:                  registration.ID,
		Status:              &cancelledStatus,
		PaymentIntentStatus: &piStatus,
	}

	if _, err := h.repo.Registration.CancelRegistration(ctx, input); err != nil {
		log.Printf("Failed to cancel registration %s: %v", registration.ID, err)
		return nil
	}

	log.Printf("Cancelled registration %s due to failed payment intent %s", registration.ID, pi.ID)
	return nil
}

func (h *Handler) handleSetupIntentSucceeded(ctx context.Context, event stripe.Event) error {
	si, err := unmarshalEvent[stripe.SetupIntent](event)
	if err != nil {
		log.Printf("Failed to unmarshal setup_intent.succeeded: %v", err)
		return nil
	}

	if si.Customer == nil || si.PaymentMethod == nil {
		log.Printf("Setup intent %s missing customer or payment method", si.ID)
		return nil
	}

	log.Printf("Setup intent %s succeeded for customer %s, payment method %s attached",
		si.ID, si.Customer.ID, si.PaymentMethod.ID)

	return nil
}

func (h *Handler) handleAccountUpdated(ctx context.Context, event stripe.Event) error {
	account, err := unmarshalEvent[stripe.Account](event)
	if err != nil {
		log.Printf("Failed to unmarshal account.updated: %v", err)
		return nil
	}

	activated := account.ChargesEnabled && account.PayoutsEnabled

	if _, err := h.repo.Organization.SetStripeAccountStatus(ctx, account.ID, activated); err != nil {
		log.Printf("Failed to update stripe account activation for %s: %v", account.ID, err)
		return nil
	}

	log.Printf("Account %s activation status updated to %v (charges_enabled=%v, payouts_enabled=%v)",
		account.ID, activated, account.ChargesEnabled, account.PayoutsEnabled)

	return nil
}
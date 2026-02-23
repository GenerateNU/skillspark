package webhook

import (
	"context"
	"log"

	"skillspark/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v84"
	"github.com/stripe/stripe-go/v84/webhook"
)

func (h *Handler) HandlePlatformWebhook(c *fiber.Ctx) error {
	payload := c.Body()
	signature := c.Get("Stripe-Signature")

	event, err := webhook.ConstructEvent(payload, signature, h.webhookSecret)
	if err != nil {
		log.Printf("Webhook signature verification failed: %v", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid signature",
		})
	}

	switch event.Type {
	case "payment_intent.payment_failed":
		return h.handlePaymentIntentFailed(c.Context(), event)
	case "setup_intent.succeeded":
		return h.handleSetupIntentSucceeded(c.Context(), event)
	default:
		log.Printf("Unhandled platform event type: %s", event.Type)
	}

	return c.SendStatus(fiber.StatusOK)
}

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
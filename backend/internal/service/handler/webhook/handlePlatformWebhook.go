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
	case "payment_method.attached":
		return h.handlePaymentMethodAdditionSuccess(c.Context(), event)
	default:
		log.Printf("Unhandled platform event type: %s", event.Type)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *Handler) handlePaymentIntentFailed(ctx context.Context, event stripe.Event) error {
	pi, err := unmarshalEvent[stripe.PaymentIntent](event)
	if err != nil {
		log.Printf("Failed to unmarshal payment_intent.payment_failed: %v", err)
		return err
	}

	registration, err := h.repo.Registration.GetRegistrationByPaymentIntentID(ctx, pi.ID, "en-US")
	if err != nil {
		log.Printf("Registration not found for payment intent %s: %v", pi.ID, err)
		return err
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
		return err
	}

	log.Printf("Cancelled registration %s due to failed payment intent %s", registration.ID, pi.ID)
	return nil
}

func (h *Handler) handlePaymentMethodAdditionSuccess(ctx context.Context, event stripe.Event) error {
	method, err := unmarshalEvent[stripe.PaymentMethod](event)
	if err != nil {
		log.Printf("Failed to unmarshal payment_mehtod.attached: %v", err)
		return err
	}

	customer := method.Customer
	if customer == nil {
		log.Printf("Payment method %s has no associated customer, skipping", method.ID)
		return nil
	}

	existing, err := h.stripeClient.GetPaymentMethodsByCustomerID(ctx, customer.ID)

	if err != nil {
		return err
	}

	for _, pm := range existing.Body.PaymentMethods {
		if pm.ID != method.ID {
			if err := h.stripeClient.DetachPaymentMethod(ctx, pm.ID); err != nil {
				return err
			}
		}
	}
	return nil
}

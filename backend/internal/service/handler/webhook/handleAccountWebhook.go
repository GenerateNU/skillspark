package webhook

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v84"
	"github.com/stripe/stripe-go/v84/webhook"
)

func (h *Handler) HandleAccountWebhook(c *fiber.Ctx) error {
	payload := c.Body()
	signature := c.Get("Stripe-Signature")

	event, err := webhook.ConstructEvent(payload, signature, h.connectWebhookSecret)
	if err != nil {
		log.Printf("Connect webhook signature verification failed: %v", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid signature",
		})
	}

	switch event.Type {
	case "account.updated":
		return h.handleAccountUpdated(c.Context(), event)
	default:
		log.Printf("Unhandled connect event type: %s", event.Type)
	}

	return c.SendStatus(fiber.StatusOK)
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

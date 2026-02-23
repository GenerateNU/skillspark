package webhook

import (
	"encoding/json"
	"log"
	"skillspark/internal/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v84"
	"github.com/stripe/stripe-go/v84/webhook"
)

type Handler struct {
	repo                 *storage.Repository
	webhookSecret        string
	connectWebhookSecret string
}

func NewHandler(repo *storage.Repository, webhookSecret string, connectWebhookSecret string) *Handler {
	return &Handler{
		repo:                 repo,
		webhookSecret:        webhookSecret,
		connectWebhookSecret: connectWebhookSecret,
	}
}

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

func unmarshalEvent[T any](event stripe.Event) (*T, error) {
	var obj T
	if err := json.Unmarshal(event.Data.Raw, &obj); err != nil {
		return nil, err
	}
	return &obj, nil
}
package routes

import (
	"skillspark/internal/service/handler/webhook"
	"skillspark/internal/storage"
	"skillspark/internal/stripeClient"

	"github.com/gofiber/fiber/v2"
)

func SetupWebhookRoutes(app *fiber.App, repo *storage.Repository, webhookSecret string, connectWebhookSecret string, sc stripeClient.StripeClientInterface) {
	handler := webhook.NewHandler(repo, webhookSecret, connectWebhookSecret, sc)

	app.Post("/api/v1/webhooks/stripe", handler.HandlePlatformWebhook)
	app.Post("/api/v1/webhooks/stripe/account", handler.HandleAccountWebhook)
}

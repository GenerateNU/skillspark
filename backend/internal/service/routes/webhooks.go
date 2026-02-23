package routes

import (
	"skillspark/internal/service/handler/webhook"
	"skillspark/internal/storage"

	"github.com/gofiber/fiber/v2"
)

func SetupWebhookRoutes(app *fiber.App, repo *storage.Repository, webhookSecret string, connectWebhookSecret string) {
	handler := webhook.NewHandler(repo, webhookSecret, connectWebhookSecret)

	app.Post("/api/v1/webhooks/stripe", handler.HandlePlatformWebhook)
	app.Post("/api/v1/webhooks/stripe/account", handler.HandleAccountWebhook)
}
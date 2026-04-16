package webhook

import (
	"encoding/json"
	"skillspark/internal/storage"
	"skillspark/internal/stripeClient"

	"github.com/stripe/stripe-go/v84"
)

type Handler struct {
	repo                 *storage.Repository
	stripeClient         stripeClient.StripeClientInterface
	webhookSecret        string
	connectWebhookSecret string
}

func NewHandler(repo *storage.Repository, webhookSecret string, connectWebhookSecret string, sc stripeClient.StripeClientInterface) *Handler {
	return &Handler{
		repo:                 repo,
		webhookSecret:        webhookSecret,
		connectWebhookSecret: connectWebhookSecret,
		stripeClient:         sc,
	}
}

func unmarshalEvent[T any](event stripe.Event) (*T, error) {
	var obj T
	if err := json.Unmarshal(event.Data.Raw, &obj); err != nil {
		return nil, err
	}
	return &obj, nil
}

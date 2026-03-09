package webhook

import (
	"encoding/json"
	"skillspark/internal/storage"

	"github.com/stripe/stripe-go/v84"
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

func unmarshalEvent[T any](event stripe.Event) (*T, error) {
	var obj T
	if err := json.Unmarshal(event.Data.Raw, &obj); err != nil {
		return nil, err
	}
	return &obj, nil
}
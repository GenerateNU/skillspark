package stripeClient

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStripeClient_DetachPaymentMethod(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Stripe integration test in short mode")
	}

	apiKey := getTestStripeAPIKey(t)
	client,_ := NewStripeClient(apiKey)
	ctx := context.Background()

	t.Run("Successfully detaches payment method from customer", func(t *testing.T) {
		t.Skip("Skipping - requires real card confirmation flow")
	})

	t.Run("Fails with invalid payment method ID", func(t *testing.T) {
		err := client.DetachPaymentMethod(ctx, "pm_invalid123")

		assert.Error(t, err)
	})

	t.Run("Fails with empty payment method ID", func(t *testing.T) {
		err := client.DetachPaymentMethod(ctx, "")

		assert.Error(t, err)
	})

	t.Run("Fails when payment method not attached", func(t *testing.T) {
		paymentMethodID := "pm_card_visa"

		err := client.DetachPaymentMethod(ctx, paymentMethodID)
		
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "not attached")
	})
}
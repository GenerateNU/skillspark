package stripeClient

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStripeClient_DetachPaymentMethod(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Stripe integration test in short mode")
	}

	apiKey := getTestStripeAPIKey(t)
	client, _ := NewStripeClient(apiKey)
	ctx := context.Background()

	t.Run("successfully detaches an attached payment method", func(t *testing.T) {
		err := client.AttachPaymentMethod(ctx, "pm_card_visa", testStripeCustomerID)
		require.NoError(t, err)

		// get the attached payment method ID
		pms, err := client.GetPaymentMethodsByCustomerID(ctx, testStripeCustomerID)
		require.NoError(t, err)
		require.NotEmpty(t, pms.Body.PaymentMethods)

		pmID := pms.Body.PaymentMethods[0].ID

		err = client.DetachPaymentMethod(ctx, pmID)

		assert.NoError(t, err)

		// verify it's gone
		pmsAfter, err := client.GetPaymentMethodsByCustomerID(ctx, testStripeCustomerID)
		require.NoError(t, err)
		for _, pm := range pmsAfter.Body.PaymentMethods {
			assert.NotEqual(t, pmID, pm.ID)
		}
	})

	t.Run("fails with invalid payment method ID", func(t *testing.T) {
		err := client.DetachPaymentMethod(ctx, "pm_invalid123")

		assert.Error(t, err)
	})

	t.Run("fails with empty payment method ID", func(t *testing.T) {
		err := client.DetachPaymentMethod(ctx, "")

		assert.Error(t, err)
	})

	t.Run("fails when payment method not attached to any customer", func(t *testing.T) {
		err := client.DetachPaymentMethod(ctx, "pm_card_visa")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "not attached")
	})
}

package stripeClient

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStripeClient_CreateSetupIntent(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Stripe integration test in short mode")
	}

	apiKey := getTestStripeAPIKey(t)
	client := NewStripeClient(apiKey)
	ctx := context.Background()

	t.Run("Successfully creates setup intent for valid customer", func(t *testing.T) {
		// First create a customer
		customer, err := client.CreateCustomer(ctx, "setuptest@example.com", "Setup Test User")
		require.NoError(t, err)
		require.NotNil(t, customer)

		// Create setup intent for the customer
		clientSecret, err := client.CreateSetupIntent(ctx, customer.ID)

		require.NoError(t, err)
		assert.NotEmpty(t, clientSecret)
		assert.Contains(t, clientSecret, "seti_") // Setup intent secret starts with seti_
		assert.Contains(t, clientSecret, "_secret_") // Contains _secret_ in the middle

		t.Logf("Created setup intent with client secret: %s", clientSecret[:20]+"...")
	})

	t.Run("Successfully creates multiple setup intents for same customer", func(t *testing.T) {
		// Create a customer
		customer, err := client.CreateCustomer(ctx, "multiple@example.com", "Multiple Setup User")
		require.NoError(t, err)

		// Create first setup intent
		clientSecret1, err := client.CreateSetupIntent(ctx, customer.ID)
		require.NoError(t, err)
		assert.NotEmpty(t, clientSecret1)

		// Create second setup intent for same customer
		clientSecret2, err := client.CreateSetupIntent(ctx, customer.ID)
		require.NoError(t, err)
		assert.NotEmpty(t, clientSecret2)

		// Should be different setup intents
		assert.NotEqual(t, clientSecret1, clientSecret2)
	})

	t.Run("Fails with invalid customer ID", func(t *testing.T) {
		clientSecret, err := client.CreateSetupIntent(ctx, "cus_invalid123")

		assert.Error(t, err)
		assert.Empty(t, clientSecret)
		assert.Contains(t, err.Error(), "No such customer")
	})

	t.Run("Fails with empty customer ID", func(t *testing.T) {
		clientSecret, err := client.CreateSetupIntent(ctx, "")

		assert.Error(t, err)
		assert.Empty(t, clientSecret)
	})

	t.Run("Successfully creates setup intent for Thai customer", func(t *testing.T) {
		// Create Thai customer
		customer, err := client.CreateCustomer(ctx, "ผู้ปกครอง@example.com", "สมชาย ใจดี")
		require.NoError(t, err)

		// Create setup intent
		clientSecret, err := client.CreateSetupIntent(ctx, customer.ID)

		require.NoError(t, err)
		assert.NotEmpty(t, clientSecret)
		assert.Contains(t, clientSecret, "seti_")
	})
}
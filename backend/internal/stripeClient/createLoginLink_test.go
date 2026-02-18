package stripeClient

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStripeClient_CreateLoginLink(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Stripe integration test in short mode")
	}

	apiKey := getTestStripeAPIKey(t)
	client := NewStripeClient(apiKey)
	ctx := context.Background()

	t.Run("Successfully creates login link for onboarded account", func(t *testing.T) {
		loginURL, err := client.CreateLoginLink(ctx, testStripeAccountID)

		require.NoError(t, err)
		assert.NotEmpty(t, loginURL)
		assert.Contains(t, loginURL, "https://connect.stripe.com")
		assert.Contains(t, loginURL, "/express/")

		t.Logf("Created login link: %s", loginURL)
	})

	t.Run("Can create multiple login links for same account", func(t *testing.T) {
		loginURL1, err := client.CreateLoginLink(ctx, testStripeAccountID)
		require.NoError(t, err)
		assert.NotEmpty(t, loginURL1)

		loginURL2, err := client.CreateLoginLink(ctx, testStripeAccountID)
		require.NoError(t, err)
		assert.NotEmpty(t, loginURL2)

		assert.NotEqual(t, loginURL1, loginURL2)
	})

	t.Run("Fails with invalid account ID", func(t *testing.T) {
		loginURL, err := client.CreateLoginLink(ctx, "acct_invalid123")

		assert.Error(t, err)
		assert.Empty(t, loginURL)
	})

	t.Run("Fails with empty account ID", func(t *testing.T) {
		loginURL, err := client.CreateLoginLink(ctx, "")

		assert.Error(t, err)
		assert.Empty(t, loginURL)
	})

	t.Run("Fails with non-Express account type", func(t *testing.T) {
		loginURL, err := client.CreateLoginLink(ctx, "acct_1SuwEe2SjsBcKnNj")

		assert.Error(t, err)
		assert.Empty(t, loginURL)
	})
}
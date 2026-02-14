package stripeClient

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStripeClient_CreateOrganizationAccount(t *testing.T) {
	t.Skip("V2 API account creation is slow and times out - skip for now")
	
	if testing.Short() {
		t.Skip("Skipping Stripe integration test in short mode")
	}

	apiKey := getTestStripeAPIKey(t)
	client := NewStripeClient(apiKey)
	ctx := context.Background()

	t.Run("Successfully creates Express account", func(t *testing.T) {
		output, err := client.CreateOrganizationAccount(
			ctx,
			"Test Bangkok Soccer Academy",
			"test+soccer@example.com",
			"TH",
		)

		require.NoError(t, err)
		require.NotNil(t, output)
		assert.NotEmpty(t, output.Body.Account.ID)
		assert.Equal(t, "express", string(output.Body.Account.Dashboard))
		assert.Equal(t, "test+soccer@example.com", output.Body.Account.ContactEmail)
		assert.Equal(t, "Test Bangkok Soccer Academy", output.Body.Account.DisplayName)

		assert.NotNil(t, output.Body.Account.Configuration)
		assert.NotNil(t, output.Body.Account.Configuration.Merchant)
		assert.NotNil(t, output.Body.Account.Configuration.Recipient)

		merchantCaps := output.Body.Account.Configuration.Merchant.Capabilities
		assert.NotNil(t, merchantCaps.CardPayments)
		assert.NotEqual(t, "", merchantCaps.CardPayments.Status)

		t.Logf("Created test account: %s", output.Body.Account.ID)
	})

	t.Run("Successfully creates account for US organization", func(t *testing.T) {
		output, err := client.CreateOrganizationAccount(
			ctx,
			"Test US Sports Center",
			"test+us@example.com",
			"US",
		)

		require.NoError(t, err)
		require.NotNil(t, output)
		assert.NotEmpty(t, output.Body.Account.ID)
		assert.Equal(t, "US", output.Body.Account.Identity.Country)
	})

	t.Run("Fails with invalid country code", func(t *testing.T) {
		output, err := client.CreateOrganizationAccount(
			ctx,
			"Test Invalid Country",
			"test+invalid@example.com",
			"INVALID",
		)

		assert.Error(t, err)
		assert.Nil(t, output)
		assert.Contains(t, err.Error(), "country")
	})

	t.Run("Fails with invalid email", func(t *testing.T) {
		output, err := client.CreateOrganizationAccount(
			ctx,
			"Test Invalid Email",
			"not-an-email",
			"TH",
		)

		assert.Error(t, err)
		assert.Nil(t, output)
	})
}

func getTestStripeAPIKey(t *testing.T) string {
	_ = godotenv.Load("../../.env")
	
	apiKey := os.Getenv("STRIPE_SECRET_TEST_KEY")
	if apiKey == "" {
		t.Skip("STRIPE_SECRET_TEST_KEY not set, skipping Stripe integration tests")
	}
	return apiKey
}
package stripeClient

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stripe/stripe-go/v84"
)

func TestGetAccount(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Stripe integration test in short mode")
	}

	apiKey := getTestStripeAPIKey(t)
	client,_ := NewStripeClient(apiKey)
	
	// Add timeout to prevent hanging
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Run("Successfully retrieves existing account", func(t *testing.T) {
		// First create an account
		createdAccount, err := client.CreateOrganizationAccount(
			ctx,
			"Get Test Org",
			"gettest@example.com",
			"TH",
		)
		require.NoError(t, err)
		require.NotNil(t, createdAccount)

		// Now retrieve it
		account, err := client.GetAccount(ctx, createdAccount.Body.Account.ID)

		require.NoError(t, err)
		require.NotNil(t, account)
		assert.Equal(t, createdAccount.Body.Account.ID, account.ID)
		assert.Equal(t, "gettest@example.com", account.ContactEmail)
		assert.Equal(t, "Get Test Org", account.DisplayName)

		// Verify included data is present
		assert.NotNil(t, account.Configuration)
		assert.NotNil(t, account.Configuration.Merchant)
		assert.NotNil(t, account.Configuration.Recipient)
		assert.NotNil(t, account.Identity)
		assert.NotNil(t, account.Defaults)

		t.Logf("Retrieved account: %s", account.ID)
	})

	t.Run("Successfully retrieves account configuration details", func(t *testing.T) {
		// Create account
		createdAccount, err := client.CreateOrganizationAccount(
			ctx,
			"Config Test Org",
			"configtest@example.com",
			"US",
		)
		require.NoError(t, err)

		// Retrieve it
		account, err := client.GetAccount(ctx, createdAccount.Body.Account.ID)

		require.NoError(t, err)

		// Verify merchant capabilities
		assert.NotNil(t, account.Configuration.Merchant.Capabilities.CardPayments)
		assert.NotEmpty(t, account.Configuration.Merchant.Capabilities.CardPayments.Status)

		// Verify recipient capabilities
		assert.NotNil(t, account.Configuration.Recipient.Capabilities.StripeBalance)
		assert.NotNil(t, account.Configuration.Recipient.Capabilities.StripeBalance.StripeTransfers)
		assert.NotEmpty(t, account.Configuration.Recipient.Capabilities.StripeBalance.StripeTransfers.Status)

		t.Logf("Card Payments Status: %s", account.Configuration.Merchant.Capabilities.CardPayments.Status)
		t.Logf("Stripe Transfers Status: %s", account.Configuration.Recipient.Capabilities.StripeBalance.StripeTransfers.Status)
	})

	t.Run("Successfully retrieves identity information", func(t *testing.T) {
		// Create account
		createdAccount, err := client.CreateOrganizationAccount(
			ctx,
			"Identity Test Org",
			"identity@example.com",
			"TH",
		)
		require.NoError(t, err)

		// Retrieve it
		account, err := client.GetAccount(ctx, createdAccount.Body.Account.ID)

		require.NoError(t, err)
		assert.NotNil(t, account.Identity)
		assert.Equal(t, "TH", account.Identity.Country)
	})

	t.Run("Successfully retrieves defaults", func(t *testing.T) {
		// Create account
		createdAccount, err := client.CreateOrganizationAccount(
			ctx,
			"Defaults Test Org",
			"defaults@example.com",
			"US",
		)
		require.NoError(t, err)

		// Retrieve it
		account, err := client.GetAccount(ctx, createdAccount.Body.Account.ID)

		require.NoError(t, err)
		assert.NotNil(t, account.Defaults)
		assert.NotEmpty(t, account.Defaults.Currency)
		
		// Verify responsibilities using Stripe constants
		assert.NotNil(t, account.Defaults.Responsibilities)
		assert.Equal(t, 
			stripe.V2CoreAccountDefaultsResponsibilitiesLossesCollectorApplication, 
			account.Defaults.Responsibilities.LossesCollector)
		assert.Equal(t, 
			stripe.V2CoreAccountDefaultsResponsibilitiesFeesCollectorApplication, 
			account.Defaults.Responsibilities.FeesCollector)

		t.Logf("Default Currency: %s", account.Defaults.Currency)
		t.Logf("Losses Collector: %s", account.Defaults.Responsibilities.LossesCollector)
		t.Logf("Fees Collector: %s", account.Defaults.Responsibilities.FeesCollector)
	})

	t.Run("Fails with invalid account ID", func(t *testing.T) {
		account, err := client.GetAccount(ctx, "acct_invalid123")

		assert.Error(t, err)
		assert.Nil(t, account)
	})

	t.Run("Fails with empty account ID", func(t *testing.T) {
		account, err := client.GetAccount(ctx, "")

		assert.Error(t, err)
		assert.Nil(t, account)
	})

	t.Run("Can retrieve same account multiple times", func(t *testing.T) {
		// Create account once
		createdAccount, err := client.CreateOrganizationAccount(
			ctx,
			"Multiple Retrieve Org",
			"multiretrieve@example.com",
			"TH",
		)
		require.NoError(t, err)

		accountID := createdAccount.Body.Account.ID

		// Retrieve multiple times
		account1, err := client.GetAccount(ctx, accountID)
		require.NoError(t, err)

		account2, err := client.GetAccount(ctx, accountID)
		require.NoError(t, err)

		// Should return same data
		assert.Equal(t, account1.ID, account2.ID)
		assert.Equal(t, account1.ContactEmail, account2.ContactEmail)
	})
}
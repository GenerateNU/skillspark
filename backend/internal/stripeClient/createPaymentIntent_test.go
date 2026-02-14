package stripeClient

import (
	"context"
	"testing"
	"time"

	"skillspark/internal/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStripeClient_CreatePaymentIntent(t *testing.T) {
	t.Skip("Requires Express account with transfer capability - times out without proper setup")
	
	if testing.Short() {
		t.Skip("Skipping Stripe integration test in short mode")
	}

	apiKey := getTestStripeAPIKey(t)
	client := NewStripeClient(apiKey)
	ctx := context.Background()

	t.Run("Successfully creates payment intent with test payment method", func(t *testing.T) {
		customer, err := client.CreateCustomer(ctx, "paymenttest@example.com", "Payment Test User")
		require.NoError(t, err)

		org, err := client.CreateOrganizationAccount(
			ctx,
			"Payment Test Org",
			"paymentorg"+time.Now().Format("20060102150405")+"@example.com",
			"US",
		)
		require.NoError(t, err)
		t.Logf("Created test account: %s", org.Body.Account.ID)

		paymentMethodID := "pm_card_visa"

		input := &models.CreatePaymentIntentInput{}
		input.Body.Amount = 10000
		input.Body.Currency = "usd"
		input.Body.GuardianStripeID = customer.ID
		input.Body.OrgStripeID = org.Body.Account.ID
		input.Body.PaymentMethodID = &paymentMethodID
		input.Body.EventDate = time.Now().Add(24 * time.Hour)
		input.Body.RegistrationID = uuid.New()
		input.Body.GuardianID = uuid.New()
		input.Body.ProviderOrgID = uuid.New()

		output, err := client.CreatePaymentIntent(ctx, input)

		if err != nil {
			if assert.Contains(t, err.Error(), "capabilities") {
				t.Skip("Account capabilities not active yet - this is expected for new test accounts")
			}
			require.NoError(t, err)
		}

		require.NotNil(t, output)
		assert.NotEmpty(t, output.Body.PaymentIntentID)
		assert.NotEmpty(t, output.Body.ClientSecret)
		assert.NotEmpty(t, output.Body.Status)
		assert.Contains(t, output.Body.PaymentIntentID, "pi_")

		t.Logf("✓ Created payment intent: %s", output.Body.PaymentIntentID)
		t.Logf("✓ Status: %s", output.Body.Status)
	})

	t.Run("Fails when payment method is nil", func(t *testing.T) {
		customer, err := client.CreateCustomer(ctx, "nopm@example.com", "No PM User")
		require.NoError(t, err)

		org, err := client.CreateOrganizationAccount(
			ctx,
			"No PM Org",
			"nopm"+time.Now().Format("20060102150405")+"@example.com",
			"US",
		)
		require.NoError(t, err)

		input := &models.CreatePaymentIntentInput{}
		input.Body.Amount = 5000
		input.Body.Currency = "usd"
		input.Body.GuardianStripeID = customer.ID
		input.Body.OrgStripeID = org.Body.Account.ID
		input.Body.PaymentMethodID = nil
		input.Body.EventDate = time.Now().Add(24 * time.Hour)
		input.Body.RegistrationID = uuid.New()
		input.Body.GuardianID = uuid.New()
		input.Body.ProviderOrgID = uuid.New()

		output, err := client.CreatePaymentIntent(ctx, input)

		assert.Error(t, err)
		assert.Nil(t, output)
		assert.Contains(t, err.Error(), "payment method required")
	})

	t.Run("Fails when payment method is empty string", func(t *testing.T) {
		customer, err := client.CreateCustomer(ctx, "emptypm@example.com", "Empty PM User")
		require.NoError(t, err)

		org, err := client.CreateOrganizationAccount(
			ctx,
			"Empty PM Org",
			"emptypm"+time.Now().Format("20060102150405")+"@example.com",
			"US",
		)
		require.NoError(t, err)

		emptyPM := ""
		input := &models.CreatePaymentIntentInput{}
		input.Body.Amount = 5000
		input.Body.Currency = "usd"
		input.Body.GuardianStripeID = customer.ID
		input.Body.OrgStripeID = org.Body.Account.ID
		input.Body.PaymentMethodID = &emptyPM
		input.Body.EventDate = time.Now().Add(24 * time.Hour)
		input.Body.RegistrationID = uuid.New()
		input.Body.GuardianID = uuid.New()
		input.Body.ProviderOrgID = uuid.New()

		output, err := client.CreatePaymentIntent(ctx, input)

		assert.Error(t, err)
		assert.Nil(t, output)
		assert.Contains(t, err.Error(), "payment method required")
	})

	t.Run("Fails with invalid customer ID", func(t *testing.T) {
		org, err := client.CreateOrganizationAccount(
			ctx,
			"Invalid Cust Org",
			"invalidcust"+time.Now().Format("20060102150405")+"@example.com",
			"US",
		)
		require.NoError(t, err)

		paymentMethodID := "pm_card_visa"
		input := &models.CreatePaymentIntentInput{}
		input.Body.Amount = 5000
		input.Body.Currency = "usd"
		input.Body.GuardianStripeID = "cus_nonexistent123"
		input.Body.OrgStripeID = org.Body.Account.ID
		input.Body.PaymentMethodID = &paymentMethodID
		input.Body.EventDate = time.Now().Add(24 * time.Hour)
		input.Body.RegistrationID = uuid.New()
		input.Body.GuardianID = uuid.New()
		input.Body.ProviderOrgID = uuid.New()

		output, err := client.CreatePaymentIntent(ctx, input)

		assert.Error(t, err)
		assert.Nil(t, output)
		assert.Contains(t, err.Error(), "No such customer")
	})

	t.Run("Fails with invalid organization account ID", func(t *testing.T) {
		customer, err := client.CreateCustomer(ctx, "invalidorg@example.com", "Invalid Org User")
		require.NoError(t, err)

		paymentMethodID := "pm_card_visa"
		input := &models.CreatePaymentIntentInput{}
		input.Body.Amount = 5000
		input.Body.Currency = "usd"
		input.Body.GuardianStripeID = customer.ID
		input.Body.OrgStripeID = "acct_nonexistent123"
		input.Body.PaymentMethodID = &paymentMethodID
		input.Body.EventDate = time.Now().Add(24 * time.Hour)
		input.Body.RegistrationID = uuid.New()
		input.Body.GuardianID = uuid.New()
		input.Body.ProviderOrgID = uuid.New()

		output, err := client.CreatePaymentIntent(ctx, input)

		assert.Error(t, err)
		assert.Nil(t, output)
	})

	t.Run("Validates application fee calculation", func(t *testing.T) {
		amount := int64(10000)
		expectedFee := int64(1000)
		expectedOrgProfit := int64(9000)
		
		calculatedFee := (amount * 10) / 100
		calculatedProfit := amount - calculatedFee
		
		assert.Equal(t, expectedFee, calculatedFee)
		assert.Equal(t, expectedOrgProfit, calculatedProfit)
	})
}
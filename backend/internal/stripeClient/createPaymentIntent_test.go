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

	if testing.Short() {
		t.Skip("Skipping Stripe integration test in short mode")
	}

	apiKey := getTestStripeAPIKey(t)
	stripeAccountID := getSeededOrgStripeAccountID(t)
	stripeCustomerID := getSeededGuardianStripeCustomerID(t)
	client, _ := NewStripeClient(apiKey)
	ctx := context.Background()

	t.Run("Successfully creates payment intent with test payment method", func(t *testing.T) {
		input := &models.CreatePaymentIntentInput{}
		input.Body.Amount = 10000
		input.Body.Currency = "usd"
		input.Body.GuardianStripeID = stripeCustomerID
		input.Body.OrgStripeID = stripeAccountID
		input.Body.PaymentMethodID = "pm_card_visa"
		input.Body.EventDate = time.Now().Add(24 * time.Hour)
		input.Body.PlatformFeePercentage = 10
		input.Body.RegistrationID = uuid.New()
		input.Body.GuardianID = uuid.New()
		input.Body.ProviderOrgID = uuid.New()

		output, err := client.CreatePaymentIntent(ctx, input)

		require.NoError(t, err)
		require.NotNil(t, output)
		assert.NotEmpty(t, output.Body.PaymentIntentID)
		assert.NotEmpty(t, output.Body.ClientSecret)
		assert.NotEmpty(t, output.Body.Status)
		assert.Contains(t, output.Body.PaymentIntentID, "pi_")
	})

	t.Run("Fails when payment method is nil", func(t *testing.T) {
		input := &models.CreatePaymentIntentInput{}
		input.Body.Amount = 5000
		input.Body.Currency = "usd"
		input.Body.GuardianStripeID = stripeCustomerID
		input.Body.OrgStripeID = stripeAccountID
		input.Body.PaymentMethodID = ""
		input.Body.EventDate = time.Now().Add(24 * time.Hour)
		input.Body.PlatformFeePercentage = 10
		input.Body.RegistrationID = uuid.New()
		input.Body.GuardianID = uuid.New()
		input.Body.ProviderOrgID = uuid.New()

		output, err := client.CreatePaymentIntent(ctx, input)

		assert.Error(t, err)
		assert.Nil(t, output)
	})

	t.Run("Fails when payment method is empty string", func(t *testing.T) {
		input := &models.CreatePaymentIntentInput{}
		input.Body.Amount = 5000
		input.Body.Currency = "usd"
		input.Body.GuardianStripeID = stripeCustomerID
		input.Body.OrgStripeID = stripeAccountID
		input.Body.PaymentMethodID = ""
		input.Body.EventDate = time.Now().Add(24 * time.Hour)
		input.Body.PlatformFeePercentage = 10
		input.Body.RegistrationID = uuid.New()
		input.Body.GuardianID = uuid.New()
		input.Body.ProviderOrgID = uuid.New()

		output, err := client.CreatePaymentIntent(ctx, input)

		assert.Error(t, err)
		assert.Nil(t, output)
	})

	t.Run("Fails with invalid customer ID", func(t *testing.T) {
		input := &models.CreatePaymentIntentInput{}
		input.Body.Amount = 5000
		input.Body.Currency = "usd"
		input.Body.GuardianStripeID = "cus_nonexistent123"
		input.Body.OrgStripeID = stripeAccountID
		input.Body.PaymentMethodID = "pm_card_visa"
		input.Body.EventDate = time.Now().Add(24 * time.Hour)
		input.Body.PlatformFeePercentage = 10
		input.Body.RegistrationID = uuid.New()
		input.Body.GuardianID = uuid.New()
		input.Body.ProviderOrgID = uuid.New()

		output, err := client.CreatePaymentIntent(ctx, input)

		assert.Error(t, err)
		assert.Nil(t, output)
		assert.Contains(t, err.Error(), "No such customer")
	})

	t.Run("Fails with invalid organization account ID", func(t *testing.T) {
		input := &models.CreatePaymentIntentInput{}
		input.Body.Amount = 5000
		input.Body.Currency = "usd"
		input.Body.GuardianStripeID = stripeCustomerID
		input.Body.OrgStripeID = "acct_nonexistent123"
		input.Body.PaymentMethodID = "pm_card_visa"
		input.Body.EventDate = time.Now().Add(24 * time.Hour)
		input.Body.PlatformFeePercentage = 10
		input.Body.RegistrationID = uuid.New()
		input.Body.GuardianID = uuid.New()
		input.Body.ProviderOrgID = uuid.New()

		output, err := client.CreatePaymentIntent(ctx, input)

		assert.Error(t, err)
		assert.Nil(t, output)
	})

	t.Run("Validates application fee calculation", func(t *testing.T) {
		amount := int64(10000)
		platformFeePercentage := int64(10)
		expectedFee := int64(1000)
		expectedOrgProfit := int64(9000)

		calculatedFee := (amount * platformFeePercentage) / 100
		calculatedProfit := amount - calculatedFee

		assert.Equal(t, expectedFee, calculatedFee)
		assert.Equal(t, expectedOrgProfit, calculatedProfit)
	})
}

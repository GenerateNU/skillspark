package stripeClient

import (
	"context"
	"skillspark/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStripeClient_RefundPayment(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Stripe integration test")
	}

	apiKey := getTestStripeAPIKey(t)
	client, _ := NewStripeClient(apiKey)
	ctx := context.Background()

	createAndCapture := func(t *testing.T) string {
		t.Helper()

		createInput := &models.CreatePaymentIntentInput{}
		createInput.Body.Amount = 10000
		createInput.Body.Currency = "usd"
		createInput.Body.GuardianStripeID = testStripeCustomerID
		createInput.Body.OrgStripeID = testStripeAccountID
		createInput.Body.PaymentMethodID = "pm_card_visa"

		created, err := client.CreatePaymentIntent(ctx, createInput)
		require.NoError(t, err)
		require.Equal(t, "requires_capture", created.Body.Status)

		captured, err := client.CapturePaymentIntent(ctx, &models.CapturePaymentIntentInput{
			PaymentIntentID: created.Body.PaymentIntentID,
		})
		require.NoError(t, err)
		require.Equal(t, "succeeded", captured.Body.Status)

		return created.Body.PaymentIntentID
	}

	t.Run("successfully refunds a captured payment", func(t *testing.T) {
		paymentIntentID := createAndCapture(t)

		result, err := client.RefundPayment(ctx, &models.RefundPaymentInput{
			PaymentIntentID: paymentIntentID,
		})

		require.NoError(t, err)
		require.NotNil(t, result)
		assert.NotEmpty(t, result.Body.RefundID)
		assert.Contains(t, result.Body.RefundID, "re_")
		assert.Equal(t, "succeeded", result.Body.Status)
		assert.Equal(t, int64(10000), result.Body.Amount)
		assert.Equal(t, "usd", result.Body.Currency)
	})

	t.Run("fails when already refunded", func(t *testing.T) {
		paymentIntentID := createAndCapture(t)

		_, err := client.RefundPayment(ctx, &models.RefundPaymentInput{PaymentIntentID: paymentIntentID})
		require.NoError(t, err)

		result, err := client.RefundPayment(ctx, &models.RefundPaymentInput{PaymentIntentID: paymentIntentID})

		require.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "charge_already_refunded")
	})

	t.Run("fails when intent requires capture", func(t *testing.T) {
		createInput := &models.CreatePaymentIntentInput{}
		createInput.Body.Amount = 10000
		createInput.Body.Currency = "usd"
		createInput.Body.GuardianStripeID = testStripeCustomerID
		createInput.Body.OrgStripeID = testStripeAccountID
		createInput.Body.PaymentMethodID = "pm_card_visa"

		created, err := client.CreatePaymentIntent(ctx, createInput)
		require.NoError(t, err)
		require.Equal(t, "requires_capture", created.Body.Status)

		result, err := client.RefundPayment(ctx, &models.RefundPaymentInput{
			PaymentIntentID: created.Body.PaymentIntentID,
		})

		require.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("fails with invalid payment intent ID", func(t *testing.T) {
		result, err := client.RefundPayment(ctx, &models.RefundPaymentInput{
			PaymentIntentID: "pi_invalid_id",
		})

		require.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "resource_missing")
	})

	t.Run("fails with empty payment intent ID", func(t *testing.T) {
		result, err := client.RefundPayment(ctx, &models.RefundPaymentInput{
			PaymentIntentID: "",
		})

		require.Error(t, err)
		assert.Nil(t, result)
	})
}

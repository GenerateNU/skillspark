package stripeClient

import (
	"context"
	"skillspark/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStripeClient_CancelPaymentIntent_RequiresCapture(t *testing.T) {
	t.Skip("Requires Express account with transfer capability - use handler mocks instead")
	
	if testing.Short() {
		t.Skip("Skipping Stripe integration test")
	}

	apiKey := getTestStripeAPIKey(t)
	client,_ := NewStripeClient(apiKey)
	ctx := context.Background()

	paymentMethodID := "pm_card_visa"

	createPIInput := &models.CreatePaymentIntentInput{}
	createPIInput.Body.Amount = 10000
	createPIInput.Body.Currency = "usd"
	createPIInput.Body.GuardianStripeID = testStripeCustomerID
	createPIInput.Body.OrgStripeID = testStripeAccountID
	createPIInput.Body.PaymentMethodID = &paymentMethodID

	createdPI, err := client.CreatePaymentIntent(ctx, createPIInput)
	require.NoError(t, err)
	require.Equal(t, "requires_capture", createdPI.Body.Status)

	cancelInput := &models.CancelPaymentIntentInput{
		PaymentIntentID: createdPI.Body.PaymentIntentID,
		StripeAccountID: testStripeAccountID,
	}

	cancelled, err := client.CancelPaymentIntent(ctx, cancelInput)

	require.NoError(t, err)
	require.NotNil(t, cancelled)
	assert.Equal(t, createdPI.Body.PaymentIntentID, cancelled.Body.PaymentIntentID)
	assert.Equal(t, "canceled", cancelled.Body.Status)
	assert.Equal(t, int64(10000), cancelled.Body.Amount)
	assert.Equal(t, "usd", cancelled.Body.Currency)
}

func TestStripeClient_CancelPaymentIntent_Succeeded(t *testing.T) {
	t.Skip("Requires Express account with transfer capability - use handler mocks instead")
	
	if testing.Short() {
		t.Skip("Skipping Stripe integration test")
	}

	apiKey := getTestStripeAPIKey(t)
	client, _ := NewStripeClient(apiKey)
	ctx := context.Background()

	paymentMethodID := "pm_card_visa"

	createPIInput := &models.CreatePaymentIntentInput{}
	createPIInput.Body.Amount = 10000
	createPIInput.Body.Currency = "usd"
	createPIInput.Body.GuardianStripeID = testStripeCustomerID
	createPIInput.Body.OrgStripeID = testStripeAccountID
	createPIInput.Body.PaymentMethodID = &paymentMethodID

	createdPI, err := client.CreatePaymentIntent(ctx, createPIInput)
	require.NoError(t, err)

	captureInput := &models.CapturePaymentIntentInput{
		PaymentIntentID: createdPI.Body.PaymentIntentID,
		StripeAccountID: testStripeAccountID,
	}
	captured, err := client.CapturePaymentIntent(ctx, captureInput)
	require.NoError(t, err)
	require.Equal(t, "succeeded", captured.Body.Status)

	cancelInput := &models.CancelPaymentIntentInput{
		PaymentIntentID: createdPI.Body.PaymentIntentID,
		StripeAccountID: testStripeAccountID,
	}

	cancelled, err := client.CancelPaymentIntent(ctx, cancelInput)

	require.NoError(t, err)
	require.NotNil(t, cancelled)
	assert.Equal(t, createdPI.Body.PaymentIntentID, cancelled.Body.PaymentIntentID)
	assert.Equal(t, "succeeded", cancelled.Body.Status)
	assert.Equal(t, int64(10000), cancelled.Body.Amount)
	assert.Equal(t, "usd", cancelled.Body.Currency)
}

func TestStripeClient_CancelPaymentIntent_InvalidID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Stripe integration test")
	}

	apiKey := getTestStripeAPIKey(t)
	client, _ := NewStripeClient(apiKey)
	ctx := context.Background()

	cancelInput := &models.CancelPaymentIntentInput{
		PaymentIntentID: "pi_invalid_id",
		StripeAccountID: testStripeAccountID,
	}

	cancelled, err := client.CancelPaymentIntent(ctx, cancelInput)

	require.Error(t, err)
	assert.Nil(t, cancelled)
}

func TestStripeClient_CancelPaymentIntent_AlreadyCanceled(t *testing.T) {
	t.Skip("Requires Express account with transfer capability - use handler mocks instead")
	
	if testing.Short() {
		t.Skip("Skipping Stripe integration test")
	}

	apiKey := getTestStripeAPIKey(t)
	client, _ := NewStripeClient(apiKey)
	ctx := context.Background()

	paymentMethodID := "pm_card_visa"

	createPIInput := &models.CreatePaymentIntentInput{}
	createPIInput.Body.Amount = 10000
	createPIInput.Body.Currency = "usd"
	createPIInput.Body.GuardianStripeID = testStripeCustomerID
	createPIInput.Body.OrgStripeID = testStripeAccountID
	createPIInput.Body.PaymentMethodID = &paymentMethodID

	createdPI, err := client.CreatePaymentIntent(ctx, createPIInput)
	require.NoError(t, err)

	cancelInput := &models.CancelPaymentIntentInput{
		PaymentIntentID: createdPI.Body.PaymentIntentID,
		StripeAccountID: testStripeAccountID,
	}

	cancelled1, err := client.CancelPaymentIntent(ctx, cancelInput)
	require.NoError(t, err)
	require.Equal(t, "canceled", cancelled1.Body.Status)

	cancelled2, err := client.CancelPaymentIntent(ctx, cancelInput)

	require.Error(t, err)
	assert.Nil(t, cancelled2)
}
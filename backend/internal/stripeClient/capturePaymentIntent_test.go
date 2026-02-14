package stripeClient

import (
	"context"
	"skillspark/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testStripeAccountID = "acct_1T0lX12SjspRdSkp"
const testStripeCustomerID = "cus_TwTqKNe9HwjesR"

func TestStripeClient_CapturePaymentIntent_Success(t *testing.T) {
	t.Skip("Requires Express account with transfer capability - use handler mocks instead")
	
	if testing.Short() {
		t.Skip("Skipping Stripe integration test")
	}

	apiKey := getTestStripeAPIKey(t)
	client := NewStripeClient(apiKey)
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

	captureInput := &models.CapturePaymentIntentInput{
		PaymentIntentID: createdPI.Body.PaymentIntentID,
		StripeAccountID: testStripeAccountID,
	}

	captured, err := client.CapturePaymentIntent(ctx, captureInput)

	require.NoError(t, err)
	require.NotNil(t, captured)
	assert.Equal(t, createdPI.Body.PaymentIntentID, captured.Body.PaymentIntentID)
	assert.Equal(t, "succeeded", captured.Body.Status)
	assert.Equal(t, int64(10000), captured.Body.Amount)
	assert.Equal(t, "usd", captured.Body.Currency)
}

func TestStripeClient_CapturePaymentIntent_InvalidID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Stripe integration test")
	}

	apiKey := getTestStripeAPIKey(t)
	client := NewStripeClient(apiKey)
	ctx := context.Background()

	captureInput := &models.CapturePaymentIntentInput{
		PaymentIntentID: "pi_invalid_id",
		StripeAccountID: testStripeAccountID,
	}

	captured, err := client.CapturePaymentIntent(ctx, captureInput)

	require.Error(t, err)
	assert.Nil(t, captured)
}

func TestStripeClient_CapturePaymentIntent_AlreadyCaptured(t *testing.T) {
	t.Skip("Requires Express account with transfer capability - use handler mocks instead")
	
	if testing.Short() {
		t.Skip("Skipping Stripe integration test")
	}

	apiKey := getTestStripeAPIKey(t)
	client := NewStripeClient(apiKey)
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

	captured1, err := client.CapturePaymentIntent(ctx, captureInput)
	require.NoError(t, err)
	require.Equal(t, "succeeded", captured1.Body.Status)

	captured2, err := client.CapturePaymentIntent(ctx, captureInput)

	require.Error(t, err)
	assert.Nil(t, captured2)
}

func TestStripeClient_CapturePaymentIntent_CanceledIntent(t *testing.T) {
	t.Skip("Requires Express account with transfer capability - use handler mocks instead")
	
	if testing.Short() {
		t.Skip("Skipping Stripe integration test")
	}

	apiKey := getTestStripeAPIKey(t)
	client := NewStripeClient(apiKey)
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
	_, err = client.CancelPaymentIntent(ctx, cancelInput)
	require.NoError(t, err)

	captureInput := &models.CapturePaymentIntentInput{
		PaymentIntentID: createdPI.Body.PaymentIntentID,
		StripeAccountID: testStripeAccountID,
	}

	captured, err := client.CapturePaymentIntent(ctx, captureInput)

	require.Error(t, err)
	assert.Nil(t, captured)
}
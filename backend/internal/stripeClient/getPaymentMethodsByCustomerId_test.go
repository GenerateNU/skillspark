package stripeClient

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStripeClient_GetPaymentMethodsByCustomerID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Stripe integration test in short mode")
	}

	apiKey := getTestStripeAPIKey(t)
	client, _ := NewStripeClient(apiKey)
	ctx := context.Background()

	t.Run("customer with no payment methods", func(t *testing.T) {
		customer, err := client.CreateCustomer(ctx, "nopayment@example.com", "No Payment User")
		require.NoError(t, err)
		require.NotNil(t, customer)

		result, err := client.GetPaymentMethodsByCustomerID(ctx, customer.ID)

		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Empty(t, result.Body.PaymentMethods)
	})

	t.Run("customer with one visa card", func(t *testing.T) {
		customer, err := client.CreateCustomer(ctx, "onevisa@example.com", "One Visa User")
		require.NoError(t, err)

		err = client.AttachPaymentMethod(ctx, "pm_card_visa", customer.ID)
		require.NoError(t, err)

		result, err := client.GetPaymentMethodsByCustomerID(ctx, customer.ID)

		require.NoError(t, err)
		require.NotNil(t, result)
		require.Len(t, result.Body.PaymentMethods, 1)

		pm := result.Body.PaymentMethods[0]
		assert.NotEmpty(t, pm.ID)
		assert.Equal(t, "card", pm.Type)
		assert.Equal(t, "visa", pm.Card.Brand)
		assert.Equal(t, "4242", pm.Card.Last4)
		assert.NotZero(t, pm.Card.ExpMonth)
		assert.NotZero(t, pm.Card.ExpYear)

		t.Logf("Retrieved payment method: %s ending in %s", pm.Card.Brand, pm.Card.Last4)
	})

	t.Run("customer with multiple cards", func(t *testing.T) {
		customer, err := client.CreateCustomer(ctx, "multicards@example.com", "Multi Card User")
		require.NoError(t, err)

		err = client.AttachPaymentMethod(ctx, "pm_card_visa", customer.ID)
		require.NoError(t, err)
		err = client.AttachPaymentMethod(ctx, "pm_card_mastercard", customer.ID)
		require.NoError(t, err)

		result, err := client.GetPaymentMethodsByCustomerID(ctx, customer.ID)

		require.NoError(t, err)
		require.NotNil(t, result)
		require.Len(t, result.Body.PaymentMethods, 2)

		brands := []string{
			result.Body.PaymentMethods[0].Card.Brand,
			result.Body.PaymentMethods[1].Card.Brand,
		}
		assert.Contains(t, brands, "visa")
		assert.Contains(t, brands, "mastercard")
	})

	t.Run("all returned payment methods have required fields", func(t *testing.T) {
		customer, err := client.CreateCustomer(ctx, "fieldcheck@example.com", "Field Check User")
		require.NoError(t, err)

		err = client.AttachPaymentMethod(ctx, "pm_card_visa", customer.ID)
		require.NoError(t, err)

		result, err := client.GetPaymentMethodsByCustomerID(ctx, customer.ID)
		require.NoError(t, err)

		for _, pm := range result.Body.PaymentMethods {
			assert.NotEmpty(t, pm.ID)
			assert.Contains(t, pm.ID, "pm_")
			assert.NotEmpty(t, pm.Type)
			assert.NotEmpty(t, pm.Card.Brand)
			assert.Len(t, pm.Card.Last4, 4)
			assert.NotZero(t, pm.Card.ExpMonth)
			assert.NotZero(t, pm.Card.ExpYear)
		}
	})

	t.Run("invalid customer ID", func(t *testing.T) {
		result, err := client.GetPaymentMethodsByCustomerID(ctx, "cus_invalid123")

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "No such customer")
	})

	t.Run("empty customer ID", func(t *testing.T) {
		result, err := client.GetPaymentMethodsByCustomerID(ctx, "")

		assert.Error(t, err)
		assert.Nil(t, result)
	})
}
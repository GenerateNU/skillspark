package stripeClient

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStripeClient_CreateCustomer(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Stripe integration test in short mode")
	}

	apiKey := getTestStripeAPIKey(t)
	client := NewStripeClient(apiKey)
	ctx := context.Background()

	t.Run("Successfully creates customer with email and name", func(t *testing.T) {
		customer, err := client.CreateCustomer(
			ctx,
			"parent@example.com",
			"John Doe",
		)

		require.NoError(t, err)
		require.NotNil(t, customer)
		assert.NotEmpty(t, customer.ID)
		assert.Equal(t, "parent@example.com", customer.Email)
		assert.Equal(t, "John Doe", customer.Name)
		assert.Equal(t, "customer", customer.Object)

		t.Logf("Created test customer: %s", customer.ID)
	})

	t.Run("Successfully creates customer with Thai name", func(t *testing.T) {
		customer, err := client.CreateCustomer(
			ctx,
			"สมชาย@example.com",
			"สมชาย ใจดี",
		)

		require.NoError(t, err)
		require.NotNil(t, customer)
		assert.NotEmpty(t, customer.ID)
		assert.Equal(t, "สมชาย@example.com", customer.Email)
		assert.Equal(t, "สมชาย ใจดี", customer.Name)
	})

	t.Run("Successfully creates customer with email only", func(t *testing.T) {
		customer, err := client.CreateCustomer(
			ctx,
			"emailonly@example.com",
			"",
		)

		require.NoError(t, err)
		require.NotNil(t, customer)
		assert.NotEmpty(t, customer.ID)
		assert.Equal(t, "emailonly@example.com", customer.Email)
		assert.Empty(t, customer.Name)
	})

	t.Run("Fails with invalid email", func(t *testing.T) {
		customer, err := client.CreateCustomer(
			ctx,
			"not-an-email",
			"Jane Smith",
		)

		assert.Error(t, err)
		assert.Nil(t, customer)
		assert.Contains(t, err.Error(), "email")
	})

	t.Run("Successfully creates customer with duplicate email", func(t *testing.T) {
		// Stripe allows duplicate emails (unlike accounts)
		email := "duplicate@example.com"
		
		customer1, err := client.CreateCustomer(ctx, email, "First Customer")
		require.NoError(t, err)
		require.NotNil(t, customer1)

		customer2, err := client.CreateCustomer(ctx, email, "Second Customer")
		require.NoError(t, err)
		require.NotNil(t, customer2)

		// Different customer IDs even with same email
		assert.NotEqual(t, customer1.ID, customer2.ID)
		assert.Equal(t, email, customer1.Email)
		assert.Equal(t, email, customer2.Email)
	})

	t.Run("Successfully creates customer with long name", func(t *testing.T) {
		longName := "Somchai Withveryverylonglastnamethatexceedsthirtychars"
		
		customer, err := client.CreateCustomer(
			ctx,
			"longname@example.com",
			longName,
		)

		require.NoError(t, err)
		require.NotNil(t, customer)
		assert.Equal(t, longName, customer.Name)
	})
}
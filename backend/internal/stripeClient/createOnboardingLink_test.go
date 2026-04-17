package stripeClient

import (
	"context"
	"testing"

	"skillspark/internal/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStripeClient_CreateAccountOnboardingLink(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Stripe integration test in short mode")
	}

	apiKey := getTestStripeAPIKey(t)
	stripeAccountID := getSeededOrgStripeAccountID(t)
	client, _ := NewStripeClient(apiKey)
	ctx := context.Background()

	t.Run("Successfully creates onboarding link for valid account", func(t *testing.T) {
		input := &models.CreateStripeOnboardingLinkClientInput{}
		input.Body.StripeAccountID = stripeAccountID
		input.Body.RefreshURL = "http://localhost:8080/onboarding/refresh"
		input.Body.ReturnURL = "http://localhost:8080/onboarding/success"

		output, err := client.CreateAccountOnboardingLink(ctx, input)

		require.NoError(t, err)
		require.NotNil(t, output)
		assert.NotEmpty(t, output.Body.OnboardingURL)
		assert.Contains(t, output.Body.OnboardingURL, "https://connect.stripe.com")

		t.Logf("Created onboarding link: %s", output.Body.OnboardingURL)
	})

	t.Run("Successfully creates onboarding link with different URLs", func(t *testing.T) {
		input := &models.CreateStripeOnboardingLinkClientInput{}
		input.Body.StripeAccountID = stripeAccountID
		input.Body.RefreshURL = "https://furever.com/setup/retry"
		input.Body.ReturnURL = "https://furever.com/dashboard"

		output, err := client.CreateAccountOnboardingLink(ctx, input)

		require.NoError(t, err)
		assert.NotEmpty(t, output.Body.OnboardingURL)
	})

	t.Run("Can create multiple onboarding links for same account", func(t *testing.T) {
		input := &models.CreateStripeOnboardingLinkClientInput{}
		input.Body.StripeAccountID = stripeAccountID
		input.Body.RefreshURL = "http://localhost:8080/refresh"
		input.Body.ReturnURL = "http://localhost:8080/return"

		output1, err := client.CreateAccountOnboardingLink(ctx, input)
		require.NoError(t, err)
		assert.NotEmpty(t, output1.Body.OnboardingURL)

		output2, err := client.CreateAccountOnboardingLink(ctx, input)
		require.NoError(t, err)
		assert.NotEmpty(t, output2.Body.OnboardingURL)

		assert.NotEqual(t, output1.Body.OnboardingURL, output2.Body.OnboardingURL)
	})

	t.Run("Fails with invalid account ID", func(t *testing.T) {
		input := &models.CreateStripeOnboardingLinkClientInput{}
		input.Body.StripeAccountID = "acct_invalid123"
		input.Body.RefreshURL = "http://localhost:8080/refresh"
		input.Body.ReturnURL = "http://localhost:8080/return"

		output, err := client.CreateAccountOnboardingLink(ctx, input)

		assert.Error(t, err)
		assert.Nil(t, output)
		assert.Contains(t, err.Error(), "No such account")
	})

	t.Run("Fails with empty account ID", func(t *testing.T) {
		input := &models.CreateStripeOnboardingLinkClientInput{}
		input.Body.StripeAccountID = ""
		input.Body.RefreshURL = "http://localhost:8080/refresh"
		input.Body.ReturnURL = "http://localhost:8080/return"

		output, err := client.CreateAccountOnboardingLink(ctx, input)

		assert.Error(t, err)
		assert.Nil(t, output)
	})

	t.Run("Fails with invalid refresh URL format", func(t *testing.T) {
		input := &models.CreateStripeOnboardingLinkClientInput{}
		input.Body.StripeAccountID = stripeAccountID
		input.Body.RefreshURL = "not-a-valid-url"
		input.Body.ReturnURL = "http://localhost:8080/return"

		output, err := client.CreateAccountOnboardingLink(ctx, input)

		assert.Error(t, err)
		assert.Nil(t, output)
		assert.Contains(t, err.Error(), "url")
	})
}

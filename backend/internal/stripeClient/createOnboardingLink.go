package stripeClient

import (
	"context"
	"skillspark/internal/models"

	"github.com/stripe/stripe-go/v84"
)



func (sc *StripeClient) CreateAccountOnboardingLink(
    ctx context.Context,
    input *models.CreateStripeOnboardingLinkClientInput,
) (*models.CreateStripeOnboardingLinkOutput, error) {
	
    params := &stripe.AccountLinkCreateParams{
        Account:    &input.Body.StripeAccountID,
        RefreshURL: &input.Body.RefreshURL,
        ReturnURL:  &input.Body.ReturnURL,
        Type:       stripe.String("account_onboarding"),
        CollectionOptions: &stripe.AccountLinkCreateCollectionOptionsParams{
            Fields: stripe.String("eventually_due"),  // collects details needed currently and ones that will be needed in the future
            },
    }
    
    link, err := sc.client.V1AccountLinks.Create(ctx, params)
    if err != nil {
        return nil, err
    }

    output := &models.CreateStripeOnboardingLinkOutput{}
    output.Body.OnboardingURL = link.URL // Doubt we need the other fields for link
    
    return output, nil
}
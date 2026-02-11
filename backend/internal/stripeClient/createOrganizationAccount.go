package stripeClient

import (
	"context"

	"skillspark/internal/models"

	"github.com/stripe/stripe-go/v84"
)

func (sc *StripeClient) CreateOrganizationAccount(
	ctx context.Context, name string, email string, country string) (*models.CreateOrgStripeAccountOutput, error) {
	params := &stripe.V2CoreAccountCreateParams{
		Identity: &stripe.V2CoreAccountCreateIdentityParams{
			Country: stripe.String(country),
		},
		DisplayName:  stripe.String(name),
		ContactEmail: stripe.String(email),
		Configuration: &stripe.V2CoreAccountCreateConfigurationParams{
			Recipient: &stripe.V2CoreAccountCreateConfigurationRecipientParams{
				Capabilities: &stripe.V2CoreAccountCreateConfigurationRecipientCapabilitiesParams{
					StripeBalance: &stripe.V2CoreAccountCreateConfigurationRecipientCapabilitiesStripeBalanceParams{
						StripeTransfers: &stripe.V2CoreAccountCreateConfigurationRecipientCapabilitiesStripeBalanceStripeTransfersParams{
							Requested: stripe.Bool(true),
						},
					},
				},
			},
			Merchant: &stripe.V2CoreAccountCreateConfigurationMerchantParams{
				Capabilities: &stripe.V2CoreAccountCreateConfigurationMerchantCapabilitiesParams{
					CardPayments: &stripe.V2CoreAccountCreateConfigurationMerchantCapabilitiesCardPaymentsParams{
						Requested: stripe.Bool(true),
					},
					// Add PromptPay here later if still wanted
				},
			},
		},
		Defaults: &stripe.V2CoreAccountCreateDefaultsParams{
			Responsibilities: &stripe.V2CoreAccountCreateDefaultsResponsibilitiesParams{
				LossesCollector: stripe.String("application"),
				FeesCollector:   stripe.String("application"),
			},
		},
		Dashboard: stripe.String("express"),
		Include: []*string{
			stripe.String("configuration.merchant"),
			stripe.String("configuration.recipient"),
			stripe.String("identity"),
			stripe.String("defaults"),
		},
	}

	acct, err := sc.client.V2CoreAccounts.Create(ctx, params)
	if err != nil {
		return nil, err
	}

    output := &models.CreateOrgStripeAccountOutput{}
    output.Body.Account = *acct

	return output, nil
}

package stripeClient

import (
	"context"
	"skillspark/internal/models"
	"time"

	"github.com/stripe/stripe-go/v84"
)

func (sc *StripeClient) CreatePaymentIntent(ctx context.Context, input *models.CreatePaymentIntentInput) (*models.CreatePaymentIntentOutput, error) {

	applicationFeeTotal := (input.Body.Amount * int64(input.Body.PlatformFeePercentage)) / 100

	params := &stripe.PaymentIntentCreateParams{
		Amount:               stripe.Int64(input.Body.Amount),
		Currency:             stripe.String(input.Body.Currency),
		Customer:             stripe.String(input.Body.GuardianStripeID),
		PaymentMethod:        stripe.String(input.Body.PaymentMethodID),
		ApplicationFeeAmount: stripe.Int64(applicationFeeTotal),
		TransferData: &stripe.PaymentIntentCreateTransferDataParams{
			Destination: stripe.String(input.Body.OrgStripeID),
		},
		OffSession: stripe.Bool(true),
		Metadata: map[string]string{
			"event_date": input.Body.EventDate.Format(time.RFC3339),
		},
		Confirm:       stripe.Bool(true),
		CaptureMethod: stripe.String("manual"),
	}

	intent, err := sc.client.V1PaymentIntents.Create(ctx, params)
	if err != nil {
		return nil, err
	}

	output := &models.CreatePaymentIntentOutput{}
	output.Body.ClientSecret = intent.ClientSecret
	output.Body.PaymentIntentID = intent.ID
	output.Body.Status = string(intent.Status)
	output.Body.TotalAmount = int(intent.Amount)
	output.Body.ProviderAmount = int(intent.Amount) - int(intent.ApplicationFeeAmount)
	output.Body.PlatformFeeAmount = int(intent.ApplicationFeeAmount)
	output.Body.Currency = input.Body.Currency

	return output, nil
}

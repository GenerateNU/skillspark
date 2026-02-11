package routes

import (
	"context"
	"net/http"
	"skillspark/internal/models"
	guardianpaymentmethod "skillspark/internal/service/handler/guardian-payment-method"
	"skillspark/internal/storage"
	"skillspark/internal/stripeClient"

	"github.com/danielgtaylor/huma/v2"
)

func SetupGuardianPaymentMethodRoutes(api huma.API, repo *storage.Repository, sc stripeClient.StripeClientInterface) {
	handler := guardianpaymentmethod.NewHandler(
		repo.Guardian,
		repo.GuardianPaymentMethod,
		sc,
	)

	huma.Register(api, huma.Operation{
		OperationID:   "get-guardian-payment-methods",
		Method:        http.MethodGet,
		Path:          "/api/v1/guardians/{guardian_id}/payment-methods",
		Summary:       "Get all payment methods for a guardian",
		Description:   "Returns list of saved payment methods for the guardian",
		Tags:          []string{"Guardian Payment Methods"},
	}, func(ctx context.Context, input *models.GetGuardianPaymentMethodsByGuardianIDInput) (*models.GetGuardianPaymentMethodsByGuardianIDOutput, error) {
		return handler.GetPaymentMethodsByGuardianID(ctx, input)
	})

	huma.Register(api, huma.Operation{
		OperationID:   "create-guardian-payment-method",
		Method:        http.MethodPost,
		Path:          "/api/v1/guardians/payment-methods",
		Summary:       "Save a new payment method for guardian",
		Description:   "Saves payment method details after frontend confirms with Stripe.js",
		Tags:          []string{"Guardian Payment Methods"},
	}, func(ctx context.Context, input *models.CreateGuardianPaymentMethodInput) (*models.CreateGuardianPaymentMethodOutput, error) {
		return handler.CreateGuardianPaymentMethod(ctx, input)
	})

	huma.Register(api, huma.Operation{
		OperationID:   "set-default-payment-method",
		Method:        http.MethodPatch,
		Path:          "/api/v1/guardians/{guardian_id}/payment-methods/{payment_method_id}/default",
		Summary:       "Set payment method as default",
		Description:   "Sets the specified payment method as the default for this guardian",
		Tags:          []string{"Guardian Payment Methods"},
	}, func(ctx context.Context, input *models.SetDefaultPaymentMethodInput) (*models.SetDefaultPaymentMethodOutput, error) {
		return handler.SetDefaultGuardianPaymentMethod(ctx, input)
	})

	huma.Register(api, huma.Operation{
		OperationID:   "delete-guardian-payment-method",
		Method:        http.MethodDelete,
		Path:          "/api/v1/guardians/payment-methods/{id}",
		Summary:       "Delete a saved payment method",
		Description:   "Removes payment method from guardian and detaches from Stripe",
		Tags:          []string{"Guardian Payment Methods"},
	}, func(ctx context.Context, input *models.DeleteGuardianPaymentMethodInput) (*models.DeleteGuardianPaymentMethodOutput, error) {
		return handler.DeleteGuardianPaymentMethod(ctx, input)
	})
}
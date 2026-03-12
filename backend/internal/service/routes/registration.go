package routes

import (
	"context"
	"net/http"
	"skillspark/internal/models"
	"skillspark/internal/service/handler/registration"
	"skillspark/internal/storage"
	"skillspark/internal/stripeClient"

	"github.com/danielgtaylor/huma/v2"
)

func SetupRegistrationRoutes(api huma.API, repo *storage.Repository, sc stripeClient.StripeClientInterface) {
	registrationHandler := registration.NewHandler(repo.Registration, repo.Child, repo.Guardian, repo.EventOccurrence, repo.Organization, sc)

	huma.Register(api, huma.Operation{
		OperationID: "create-registration",
		Method:      http.MethodPost,
		Path:        "/api/v1/registrations",
		Summary:     "Create a new registration",
		Description: "Create a new registration for a child to attend an event occurrence",
		Tags:        []string{"Registrations"},
	}, func(ctx context.Context, input *models.CreateRegistrationInput) (*models.CreateRegistrationOutput, error) {
		return registrationHandler.CreateRegistration(ctx, input)
	})

	huma.Register(api, huma.Operation{
		OperationID: "get-registration-by-id",
		Method:      http.MethodGet,
		Path:        "/api/v1/registrations/{id}",
		Summary:     "Get registration by ID",
		Description: "Retrieve a specific registration by its unique identifier",
		Tags:        []string{"Registrations"},
	}, func(ctx context.Context, input *models.GetRegistrationByIDInput) (*models.GetRegistrationByIDOutput, error) {
		return registrationHandler.GetRegistrationByID(ctx, input)
	})

	huma.Register(api, huma.Operation{
		OperationID: "get-registrations-by-child-id",
		Method:      http.MethodGet,
		Path:        "/api/v1/registrations/child/{child_id}",
		Summary:     "Get registrations by child ID",
		Description: "Retrieve all registrations for a specific child",
		Tags:        []string{"Registrations"},
	}, func(ctx context.Context, input *models.GetRegistrationsByChildIDInput) (*models.GetRegistrationsByChildIDOutput, error) {
		return registrationHandler.GetRegistrationsByChildID(ctx, input)
	})

	huma.Register(api, huma.Operation{
		OperationID: "get-registrations-by-guardian-id",
		Method:      http.MethodGet,
		Path:        "/api/v1/registrations/guardian/{guardian_id}",
		Summary:     "Get registrations by guardian ID",
		Description: "Retrieve all registrations for children under a specific guardian",
		Tags:        []string{"Registrations"},
	}, func(ctx context.Context, input *models.GetRegistrationsByGuardianIDInput) (*models.GetRegistrationsByGuardianIDOutput, error) {
		return registrationHandler.GetRegistrationsByGuardianID(ctx, input)
	})

	huma.Register(api, huma.Operation{
		OperationID: "get-registrations-by-event-occurrence-id",
		Method:      http.MethodGet,
		Path:        "/api/v1/registrations/event_occurrence/{event_occurrence_id}",
		Summary:     "Get registrations by event occurrence ID",
		Description: "Retrieve all registrations for a specific event occurrence",
		Tags:        []string{"Registrations"},
	}, func(ctx context.Context, input *models.GetRegistrationsByEventOccurrenceIDInput) (*models.GetRegistrationsByEventOccurrenceIDOutput, error) {
		return registrationHandler.GetRegistrationsByEventOccurrenceID(ctx, input)
	})

	huma.Register(api, huma.Operation{
		OperationID: "update-registration",
		Method:      http.MethodPatch,
		Path:        "/api/v1/registrations/{id}",
		Summary:     "Update a registration",
		Description: "Update the child associated with a registration",
		Tags:        []string{"Registrations"},
	}, func(ctx context.Context, input *models.UpdateRegistrationInput) (*models.UpdateRegistrationOutput, error) {
		return registrationHandler.UpdateRegistration(ctx, input)
	})

	huma.Register(api, huma.Operation{
		OperationID: "cancel-registration",
		Method:      http.MethodPost,
		Path:        "/api/v1/registrations/{id}/cancel",
		Summary:     "Cancel a registration",
		Description: "Cancel a registration and process refund if applicable",
		Tags:        []string{"Registrations"},
	}, func(ctx context.Context, input *models.CancelRegistrationInput) (*models.CancelRegistrationOutput, error) {
		return registrationHandler.CancelRegistration(ctx, input)
	})

	huma.Register(api, huma.Operation{
		OperationID: "update-registration-payment-status",
		Method:      http.MethodPatch,
		Path:        "/api/v1/registrations/{id}/payment-status",
		Summary:     "Update registration payment status",
		Description: "Update the payment intent status for a registration (typically called by webhooks)",
		Tags:        []string{"Registrations"},
	}, func(ctx context.Context, input *models.UpdateRegistrationPaymentStatusInput) (*models.UpdateRegistrationPaymentStatusOutput, error) {
		return registrationHandler.UpdateRegistrationPaymentStatus(ctx, input)
	})
}
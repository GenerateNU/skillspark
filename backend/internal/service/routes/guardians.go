package routes

import (
	"context"
	"net/http"
	"skillspark/internal/config"
	"skillspark/internal/models"
	"skillspark/internal/service/handler/guardian"
	"skillspark/internal/storage"
	"skillspark/internal/stripeClient"

	"github.com/danielgtaylor/huma/v2"
)

func SetupGuardiansRoutes(api huma.API, repo *storage.Repository, sc stripeClient.StripeClientInterface, config config.Config) {
	guardianHandler := guardian.NewHandler(repo.Guardian, repo.GetDB(), sc, config.Supabase)
	huma.Register(api, huma.Operation{
		OperationID: "get-guardian-by-id",
		Method:      http.MethodGet,
		Path:        "/api/v1/guardians/{id}",
		Summary:     "Get a guardian by id",
		Description: "Returns a guardian by id",
		Tags:        []string{"Guardians"},
	}, func(ctx context.Context, input *models.GetGuardianByIDInput) (*models.GetGuardianByIDOutput, error) {
		guardian, err := guardianHandler.GetGuardianById(ctx, input)
		if err != nil {
			return nil, err
		}

		return &models.GetGuardianByIDOutput{
			Body: guardian,
		}, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "delete-guardian",
		Method:      http.MethodDelete,
		Path:        "/api/v1/guardians/{id}",
		Summary:     "Delete a guardian by id",
		Description: "Deletes a guardian by id",
		Tags:        []string{"Guardians"},
	}, func(ctx context.Context, input *models.DeleteGuardianInput) (*models.DeleteGuardianOutput, error) {
		guardian, err := guardianHandler.DeleteGuardian(ctx, input)
		if err != nil {
			return nil, err
		}

		return &models.DeleteGuardianOutput{
			Body: guardian,
		}, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "get-guardian-by-child-id",
		Method:      http.MethodGet,
		Path:        "/api/v1/guardians/child/{child_id}",
		Summary:     "Get a guardian by child id",
		Description: "Returns a guardian by child id",
		Tags:        []string{"Guardians"},
	}, func(ctx context.Context, input *models.GetGuardianByChildIDInput) (*models.GetGuardianByChildIDOutput, error) {
		guardian, err := guardianHandler.GetGuardianByChildId(ctx, input)
		if err != nil {
			return nil, err
		}

		return &models.GetGuardianByChildIDOutput{
			Body: guardian,
		}, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "update-guardian",
		Method:      http.MethodPut,
		Path:        "/api/v1/guardians/{id}",
		Summary:     "Update a guardian by id",
		Description: "Updates a guardian by id",
		Tags:        []string{"Guardians"},
	}, func(ctx context.Context, input *models.UpdateGuardianInput) (*models.UpdateGuardianOutput, error) {
		guardian, err := guardianHandler.UpdateGuardian(ctx, input)
		if err != nil {
			return nil, err
		}

		return &models.UpdateGuardianOutput{
			Body: guardian,
		}, nil
	})
}

package routes

import (
	"context"
	"net/http"
	"skillspark/internal/config"
	"skillspark/internal/models"
	"skillspark/internal/service/handler/manager"
	"skillspark/internal/storage"

	"github.com/danielgtaylor/huma/v2"
)

func SetupManagerRoutes(api huma.API, repo *storage.Repository, config config.Config) {
	managerHandler := manager.NewHandler(repo.Manager, repo.Guardian, config.Supabase)

	huma.Register(api, huma.Operation{
		OperationID: "get-manager-by-id",
		Method:      http.MethodGet,
		Path:        "/api/v1/manager/{id}",
		Summary:     "Get a manager by id",
		Description: "Returns a manager by id",
		Tags:        []string{"Managers"},
	}, func(ctx context.Context, input *models.GetManagerByIDInput) (*models.GetManagerByIDOutput, error) {
		manager, err := managerHandler.GetManagerByID(ctx, input)
		if err != nil {
			return nil, err
		}

		return &models.GetManagerByIDOutput{
			Body: manager,
		}, nil
	})

	// get by org id here
	huma.Register(api, huma.Operation{
		OperationID: "get-manager-by-org-id",
		Method:      http.MethodGet,
		Path:        "/api/v1/manager/org/{organization_id}",
		Summary:     "Get a manager by organization id",
		Description: "Returns a manager by organization id",
		Tags:        []string{"Managers"},
	}, func(ctx context.Context, input *models.GetManagerByOrgIDInput) (*models.GetManagerByOrgIDOutput, error) {
		manager, err := managerHandler.GetManagerByOrgID(ctx, input)
		if err != nil {
			return nil, err
		}

		return &models.GetManagerByOrgIDOutput{
			Body: manager,
		}, nil
	})

	// create manager here

	huma.Register(api, huma.Operation{
		OperationID: "post-manager",
		Method:      http.MethodPost,
		Path:        "/api/v1/manager",
		Summary:     "posts a manager",
		Description: "Returns a manager by id",
		Tags:        []string{"Managers"},
	}, func(ctx context.Context, input *models.CreateManagerInput) (*models.CreateManagerOutput, error) {
		manager, err := managerHandler.CreateManager(ctx, input)
		if err != nil {
			return nil, err
		}

		return &models.CreateManagerOutput{
			Body: manager,
		}, nil
	})

	// delete manager here

	huma.Register(api, huma.Operation{
		OperationID: "delete-manager",
		Method:      http.MethodDelete,
		Path:        "/api/v1/manager/{id}",
		Summary:     "Deletes a manager by id",
		Description: "Returns a manager by id",
		Tags:        []string{"Managers"},
	}, func(ctx context.Context, input *models.DeleteManagerInput) (*models.DeleteManagerOutput, error) {
		manager, err := managerHandler.DeleteManager(ctx, input)
		if err != nil {
			return nil, err
		}

		return &models.DeleteManagerOutput{
			Body: manager,
		}, nil
	})

	// patch manager here

	huma.Register(api, huma.Operation{
		OperationID: "patch-manager",
		Method:      http.MethodPatch,
		Path:        "/api/v1/manager",
		Summary:     "Updates a manager by id",
		Description: "Returns a manager by id",
		Tags:        []string{"Managers"},
	}, func(ctx context.Context, input *models.PatchManagerInput) (*models.PatchManagerOutput, error) {
		manager, err := managerHandler.PatchManager(ctx, input)
		if err != nil {
			return nil, err
		}

		return &models.PatchManagerOutput{
			Body: manager,
		}, nil
	})

}

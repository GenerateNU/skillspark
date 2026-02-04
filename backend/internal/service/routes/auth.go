package routes

import (
	"context"
	"net/http"
	"skillspark/internal/config"
	"skillspark/internal/models"
	"skillspark/internal/service/handler/auth"
	"skillspark/internal/storage"

	"github.com/danielgtaylor/huma/v2"
)

func SetupAuthRoutes(api huma.API, repo *storage.Repository, config config.Config) {
	authHandler := auth.NewHandler(config.Supabase, repo.User, repo.Guardian, repo.Manager)

	huma.Register(api, huma.Operation{
		OperationID: "signup-guardian",
		Method:      http.MethodPost,
		Path:        "/api/v1/auth/signup/guardian",
		Summary:     "Register a guardian",
		Description: "Registers a guardian",
		Tags:        []string{"Auth"},
	}, func(ctx context.Context, input *models.GuardianSignUpInput) (*models.GuardianSignUpOutput, error) {
		guardianRes, err := authHandler.GuardianSignUp(ctx, input)

		if err != nil {
			return nil, err
		}

		return guardianRes, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "signup-manager",
		Method:      http.MethodPost,
		Path:        "/api/v1/auth/signup/manager",
		Summary:     "Register a manager",
		Description: "Registers a manager",
		Tags:        []string{"Auth"},
	}, func(ctx context.Context, input *models.ManagerSignUpInput) (*models.ManagerSignUpOutput, error) {
		managerRes, err := authHandler.ManagerSignUp(ctx, input)

		if err != nil {
			return nil, err
		}

		return managerRes, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "login-guardian",
		Method:      http.MethodPost,
		Path:        "/api/v1/auth/login/guardian",
		Summary:     "Login a guardian",
		Description: "Logs in a guardian",
		Tags:        []string{"Auth"},
	}, func(ctx context.Context, input *models.LoginInput) (*models.GuardianLoginOutput, error) {
		guardianRes, err := authHandler.GuardianLogin(ctx, input)

		if err != nil {
			return nil, err
		}

		return guardianRes, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "login-manager",
		Method:      http.MethodPost,
		Path:        "/api/v1/auth/login/manager",
		Summary:     "Login a manager",
		Description: "Logs in a manager",
		Tags:        []string{"Auth"},
	}, func(ctx context.Context, input *models.LoginInput) (*models.ManagerLoginOutput, error) {
		managerRes, err := authHandler.ManagerLogin(ctx, input)

		if err != nil {
			return nil, err
		}

		return managerRes, nil
	})

}

package routes

import (
	"context"
	"net/http"
	"skillspark/internal/models"
	"skillspark/internal/service/handler/user"
	"skillspark/internal/storage"

	"github.com/danielgtaylor/huma/v2"
)

func SetupUserRoutes(api huma.API, repo *storage.Repository) {

	userHandler := user.NewHandler(repo.User)

	huma.Register(api, huma.Operation{
		OperationID: "get-user-by-username",
		Method:      http.MethodGet,
		Path:        "/api/v1/user/{username}",
		Summary:     "Get a user by username",
		Description: "Returns a user by username",
		Tags:        []string{"User"},
	}, func(ctx context.Context, input *models.GetUserByUsernameInput) (*models.GetUserByUsernameOutput, error) {
		user, err := userHandler.GetUserByUsername(ctx, input)
		if err != nil {
			return nil, err
		}

		return &models.GetUserByUsernameOutput{
			Body: user,
		}, nil
	})
}

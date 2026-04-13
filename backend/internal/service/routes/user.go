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
		OperationID: "username-exists",
		Method:      http.MethodGet,
		Path:        "/api/v1/user/{username}",
		Summary:     "Check if a username exists",
		Description: "Returns whether a user with the given username exists",
		Tags:        []string{"User"},
	}, func(ctx context.Context, input *models.GetUserByUsernameInput) (*models.UsernameExistsOutput, error) {
		exists, err := userHandler.GetUserByUsername(ctx, input)
		if err != nil {
			return nil, err
		}

		out := &models.UsernameExistsOutput{}
		out.Body.Exists = exists
		return out, nil
	})
}

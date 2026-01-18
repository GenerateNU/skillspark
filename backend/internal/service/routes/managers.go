package routes

import (
	"context"
	"net/http"
	"skillspark/internal/models"
	"skillspark/internal/service/handler/manager"
	"skillspark/internal/storage"

	"github.com/danielgtaylor/huma/v2"
)

// idk what ctx does
// idk what the tags feature does
// dont rlly know what huma does in general
// how was the manager table created in the schema???
func SetupManagerRoutes(api huma.API, repo *storage.Repository) {
	managerHandler := manager.NewHandler(repo.Manager)

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
			Body: *manager,
		}, nil
	})

	// get by org id here

	// create manager here

	// delete manager here

	// patch manager here

}

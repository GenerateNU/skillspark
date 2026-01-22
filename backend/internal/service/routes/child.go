package routes

import (
	"context"
	"net/http"
	"skillspark/internal/models"
	"skillspark/internal/service/handler/child"
	"skillspark/internal/storage"

	"github.com/danielgtaylor/huma/v2"
)

func SetupChildRoutes(api huma.API, repo *storage.Repository) {

	// childHandler is a very suspicious name...
	childHandler := child.NewHandler(repo.Child)

	huma.Register(api, huma.Operation{
		OperationID: "get-child-by-id",
		Method:      http.MethodGet,
		Path:        "/api/v1/child/{id}",
		Summary:     "Get a child by id",
		Description: "Returns a child by id",
		Tags:        []string{"Child"},
	}, func(ctx context.Context, input *models.ChildIDInput) (*models.ChildOutput, error) {
		child, err := childHandler.GetChildByID(ctx, input)
		if err != nil {
			return nil, err
		}

		return &models.ChildOutput{
			Body: child,
		}, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "get-children-by-guardian-id",
		Method:      http.MethodGet,
		Path:        "/api/v1/children/{id}",
		Summary:     "Get all children for a guardian",
		Description: "Returns the list of children associated with a given guardian ID",
		Tags:        []string{"Child"},
	}, func(ctx context.Context, input *models.GuardianIDInput) (*models.ChildrenOutput, error) {
		children, err := childHandler.GetChildrenByParentID(ctx, input)
		if err != nil {
			return nil, err
		}

		return &models.ChildrenOutput{
			Body: children,
		}, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "delete-child",
		Method:      http.MethodDelete,
		Path:        "/api/v1/child/{id}",
		Summary:     "Delete a child",
		Description: "Deletes a child by ID",
		Tags:        []string{"Child"},
	}, func(ctx context.Context, input *models.ChildIDInput) (*models.ChildOutput, error) {
		child, err := childHandler.DeleteChildByID(ctx, input)
		if err != nil {
			return nil, err
		}

		return &models.ChildOutput{
			Body: child,
		}, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "update-child",
		Method:      http.MethodPatch,
		Path:        "/api/v1/child/{id}",
		Summary:     "Updates a child",
		Description: "Updates a child by ID",
		Tags:        []string{"Child"},
	}, func(ctx context.Context, input *models.UpdateChildInput) (*models.ChildOutput, error) {
		child, err := childHandler.UpdateChildByID(ctx, input.ID, input)
		if err != nil {
			return nil, err
		}

		return &models.ChildOutput{
			Body: child,
		}, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "create-child",
		Method:      http.MethodPost,
		Path:        "/api/v1/child",
		Summary:     "Creates a child",
		Description: "Creates a child",
		Tags:        []string{"Child"},
	}, func(ctx context.Context, input *models.CreateChildInput) (*models.ChildOutput, error) {
		child, err := childHandler.CreateChild(ctx, input)
		if err != nil {
			return nil, err
		}

		return &models.ChildOutput{
			Body: child,
		}, nil
	})
}

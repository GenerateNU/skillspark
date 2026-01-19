package routes

import (
	"context"
	"net/http"
	"skillspark/internal/models"
	"skillspark/internal/service/handler/event"
	"skillspark/internal/storage"

	"github.com/danielgtaylor/huma/v2"
)

func SetupEventRoutes(api huma.API, repo *storage.Repository) {
	eventHandler := event.NewHandler(repo.Event)

	// POST /api/v1/events
	huma.Register(api, huma.Operation{
		OperationID: "create-event",
		Method:      http.MethodPost,
		Path:        "/api/v1/events",
		Summary:     "Create a new event",
		Description: "Creates a new event",
		Tags:        []string{"Events"},
	}, func(ctx context.Context, input *models.CreateEventInput) (*models.CreateEventOutput, error) {
		event, err := eventHandler.CreateEvent(ctx, input)
		if err != nil {
			return nil, err
		}

		return &models.CreateEventOutput{
			Body: event,
		}, nil
	})

	// PATCH /api/v1/events/{id}
	huma.Register(api, huma.Operation{
		OperationID: "patch-event",
		Method:      http.MethodPatch,
		Path:        "/api/v1/events/{id}",
		Summary:     "Update an existing event",
		Description: "Updates an existing event",
		Tags:        []string{"Events"},
	}, func(ctx context.Context, input *models.PatchEventInput) (*models.PatchEventOutput, error) {
		event, err := eventHandler.PatchEvent(ctx, input)
		if err != nil {
			return nil, err
		}

		return &models.PatchEventOutput{
			Body: event,
		}, nil
	})

	// DELETE /api/v1/events/{id}
	huma.Register(api, huma.Operation{
		OperationID: "delete-event",
		Method:      http.MethodDelete,
		Path:        "/api/v1/events/{id}",
		Summary:     "Delete an existing event by id",
		Description: "Deletes an existing event by id",
		Tags:        []string{"Events"},
	}, func(ctx context.Context, input *models.DeleteEventInput) (*struct{}, error) {
		err := eventHandler.DeleteEvent(ctx, input.ID)
		if err != nil {
			return nil, err
		}

		return nil, nil
	})

}

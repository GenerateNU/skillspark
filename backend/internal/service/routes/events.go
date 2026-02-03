package routes

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"skillspark/internal/models"
	"skillspark/internal/s3_client"
	"skillspark/internal/service/handler/event"
	"skillspark/internal/storage"

	"github.com/danielgtaylor/huma/v2"
)

func SetupEventRoutes(api huma.API, repo *storage.Repository, s3Client s3_client.S3Interface) {
	eventHandler := event.NewHandler(repo.Event, s3Client)

	// POST /api/v1/events
	huma.Register(api, huma.Operation{
		OperationID: "create-event",
		Method:      http.MethodPost,
		Path:        "/api/v1/events",
		Summary:     "Create a new event",
		Description: "Creates a new event",
		Tags:        []string{"Events"},
	}, func(ctx context.Context, input *models.CreateEventRouteInput) (*models.CreateEventOutput, error) {

		formData := input.RawBody.Data()

		eventBody := models.CreateEventBody{
			Title:          formData.Title,
			Description:    formData.Description,
			OrganizationID: formData.OrganizationID,
			AgeRangeMin:    &formData.AgeRangeMin,
			AgeRangeMax:    &formData.AgeRangeMax,
			Category:       formData.Category,
		}

		eventModel := models.CreateEventInput{
			Body: eventBody,
		}

		updateBody := models.UpdateEventBody{
			Title:          &formData.Title,
			Description:    &formData.Description,
			OrganizationID: &formData.OrganizationID,
			AgeRangeMin:    &formData.AgeRangeMin,
			AgeRangeMax:    &formData.AgeRangeMax,
			Category:       &formData.Category,
		}

		image_data, err := io.ReadAll(formData.HeaderImage)
		if err != nil {
			return nil, err
		}

		// io.readall on input
		event, err := eventHandler.CreateEvent(ctx, &eventModel, &updateBody, &image_data, s3Client)

		if err != nil {
			return nil, err
		}

		return &models.CreateEventOutput{
			Body: event,
		}, nil
	})

	// PATCH /api/v1/events/{id}
	huma.Register(api, huma.Operation{
		OperationID: "update-event",
		Method:      http.MethodPatch,
		Path:        "/api/v1/events/{id}",
		Summary:     "Update an existing event",
		Description: "Updates an existing event",
		Tags:        []string{"Events"},
	}, func(ctx context.Context, input *models.UpdateEventRouteInput) (*models.UpdateEventOutput, error) {

		formData := input.RawBody.Data()

		eventBody := models.UpdateEventBody{
			Title:          &formData.Title,
			Description:    &formData.Description,
			OrganizationID: &formData.OrganizationID,
			AgeRangeMin:    &formData.AgeRangeMin,
			AgeRangeMax:    &formData.AgeRangeMax,
			Category:       &formData.Category,
		}

		eventModel := models.UpdateEventInput{
			ID:   input.ID,
			Body: eventBody,
		}

		image_data, err := io.ReadAll(formData.HeaderImage)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		// io.readall on input
		event, err := eventHandler.UpdateEvent(ctx, &eventModel, &image_data, s3Client)

		if err != nil {
			return nil, err
		}

		return &models.UpdateEventOutput{
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
	}, func(ctx context.Context, input *models.DeleteEventInput) (*models.DeleteEventOutput, error) {
		msg, err := eventHandler.DeleteEvent(ctx, input.ID)
		if err != nil {
			return nil, err
		}

		return &models.DeleteEventOutput{
			Body: struct {
				Message string `json:"message" doc:"Success message"`
			}{
				Message: msg,
			},
		}, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "get-event-occurrences-by-event-id",
		Method:      http.MethodGet,
		Path:        "/api/v1/events/{event_id}/event-occurrences/",
		Summary:     "Get event occurrences by event ID",
		Description: "Returns event occurrences that match the event ID",
		Tags:        []string{"Events"},
	}, func(ctx context.Context, input *models.GetEventOccurrencesByEventIDInput) (*models.GetEventOccurrencesByEventIDOutput, error) {
		eventOccurrences, err := eventHandler.GetEventOccurrencesByEventID(ctx, input, s3Client)
		if err != nil {
			return nil, err
		}

		return &models.GetEventOccurrencesByEventIDOutput{
			Body: eventOccurrences,
		}, nil
	})
}

package routes

import (
	"context"
	"net/http"
	"skillspark/internal/models"
	"skillspark/internal/service/handler/event-occurrence"
	"skillspark/internal/storage"
	"skillspark/internal/utils"

	"github.com/danielgtaylor/huma/v2"
)

func SetupEventOccurrencesRoutes(api huma.API, repo *storage.Repository) {
	eventOccurrenceHandler := eventoccurrence.NewHandler(repo.EventOccurrence)
	huma.Register(api, huma.Operation{
		OperationID: "get-all-event-occurrences",
		Method:      http.MethodGet,
		Path:        "/api/v1/event-occurrences",
		Summary:     "Get all event occurrences",
		Description: "Returns a list of all event occurrences in the database",
		Tags:        []string{"Event Occurrences"},
	}, func(ctx context.Context, input *models.GetAllEventOccurrencesInput) (*models.GetAllEventOccurrencesOutput, error) {
		page := input.Page
		if page == 0 {
			page = 1
		}
		limit := input.Limit
		if limit == 0 {
			limit = 10
		}

		pagination := utils.Pagination{
			Page:  page,
			Limit: limit,
		}
		eventOccurrences, err := eventOccurrenceHandler.GetAllEventOccurrences(ctx, pagination)
		if err != nil {
			return nil, err
		}

		return &models.GetAllEventOccurrencesOutput{
			Body: eventOccurrences,
		}, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "get-event-occurrences-by-id",
		Method:      http.MethodGet,
		Path:        "/api/v1/event-occurrences/{id}",
		Summary:     "Get an event occurrence by ID",
		Description: "Returns an event occurrence that matches the ID",
		Tags:        []string{"Event Occurrences"},
	}, func(ctx context.Context, input *models.GetEventOccurrenceByIDInput) (*models.GetEventOccurrenceByIDOutput, error) {
		eventOccurrence, err := eventOccurrenceHandler.GetEventOccurrenceByID(ctx, input)
		if err != nil {
			return nil, err
		}

		return &models.GetEventOccurrenceByIDOutput{
			Body: eventOccurrence,
		}, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "post-event-occurrence",
		Method:      http.MethodPost,
		Path:        "/api/v1/event-occurrences",
		Summary:     "Create an event occurrence",
		Description: "Creates a new event occurrence in the database",
		Tags:        []string{"Event Occurrences"},
	}, func(ctx context.Context, input *models.CreateEventOccurrenceInput) (*models.CreateEventOccurrenceOutput, error) {
		eventOccurrence, err := eventOccurrenceHandler.CreateEventOccurrence(ctx, input)
		if err != nil {
			return nil, err
		}

		return &models.CreateEventOccurrenceOutput{
			Body: eventOccurrence,
		}, nil
	})
}
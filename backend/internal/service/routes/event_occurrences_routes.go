package routes

import (
	"context"
	"net/http"
	"skillspark/internal/models"
	eventoccurrence "skillspark/internal/service/handler/event-occurrence"
	"skillspark/internal/storage"
	"skillspark/internal/utils"

	"github.com/danielgtaylor/huma/v2"
)

func SetupEventOccurrencesRoutes(api huma.API, repo *storage.Repository) {
	eventOccurrenceHandler := eventoccurrence.NewHandler(repo.EventOccurrence, repo.Manager, repo.Event, repo.Location)

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

		// all-or-none validation
		if (input.Latitude.Set || input.Longitude.Set || input.RadiusKm != 0) &&
			!(input.Latitude.Set && input.Longitude.Set && input.RadiusKm != 0) {
			return nil, huma.Error400BadRequest("lat, lng, and radius_km must all be provided together")
		}

		// optional: enforce positive radius
		if input.RadiusKm < 0 {
			return nil, huma.Error400BadRequest("radius_km must be positive")
		}

		// map to DB-level filter object
		var filters models.GetAllEventOccurrencesFilter

		if input.Search != "" {
			filters.Search = &input.Search
		}

		if input.Latitude.Set && input.Longitude.Set && input.RadiusKm != 0 {
			filters.Latitude = &input.Latitude.Value
			filters.Longitude = &input.Longitude.Value
			filters.RadiusKm = &input.RadiusKm
		}

		if input.MinDuration != 0 {
			filters.MinDurationMinutes = &input.MinDuration
		}

		if input.MaxDuration != 0 {
			filters.MaxDurationMinutes = &input.MaxDuration
		}

		if input.MinDuration != 0 && input.MaxDuration != 0 && input.MinDuration > input.MaxDuration {
			return nil, huma.Error400BadRequest("min_duration cannot be larger than max_duration")
		}

		if input.PriceTier != "" {
			if input.PriceTier != "$" && input.PriceTier != "$$" && input.PriceTier != "$$$" {
				return nil, huma.Error400BadRequest("price tier must be one of $, $$, $$$")
			} else {
				filters.PriceTier = &input.PriceTier
			}
		}

		eventOccurrences, err := eventOccurrenceHandler.GetAllEventOccurrences(ctx, pagination, filters)
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

	huma.Register(api, huma.Operation{
		OperationID: "patch-event-occurrence",
		Method:      http.MethodPatch,
		Path:        "/api/v1/event-occurrences/{id}",
		Summary:     "Update an event occurrence",
		Description: "Updates an event occurrence in the database",
		Tags:        []string{"Event Occurrences"},
	}, func(ctx context.Context, input *models.UpdateEventOccurrenceInput) (*models.UpdateEventOccurrenceOutput, error) {
		eventOccurrence, err := eventOccurrenceHandler.UpdateEventOccurrence(ctx, input)
		if err != nil {
			return nil, err
		}

		return &models.UpdateEventOccurrenceOutput{
			Body: eventOccurrence,
		}, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "cancel-event-occurrence",
		Method:      http.MethodDelete,
		Path:        "/api/v1/event-occurrences/{id}",
		Summary:     "Cancel an event occurrence and cancel its associated registrations",
		Description: "Cancel an event occurrence and cancel its associated registrations",
		Tags:        []string{"Event Occurrences"},
	}, func(ctx context.Context, input *models.CancelEventOccurrenceInput) (*models.CancelEventOccurrenceOutput, error) {

		msg, err := eventOccurrenceHandler.CancelEventOccurrence(ctx, input.ID)
		if err != nil {
			return nil, err
		}

		return &models.CancelEventOccurrenceOutput{
			Body: struct {
				Message string `json:"message" doc:"Success message"`
			}{
				Message: msg,
			},
		}, nil
	})
}

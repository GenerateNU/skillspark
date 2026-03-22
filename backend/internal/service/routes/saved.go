package routes

import (
	"context"
	"net/http"
	"skillspark/internal/models"
	"skillspark/internal/service/handler/saved"
	"skillspark/internal/storage"
	"skillspark/internal/utils"

	"github.com/danielgtaylor/huma/v2"
)

func SetUpSavedRoutes(api huma.API, repo *storage.Repository) {

	savedHandler := saved.NewHandler(repo.Saved, repo.Guardian)

	huma.Register(api, huma.Operation{
		OperationID: "get-saved-by-guardian-id",
		Method:      http.MethodGet,
		Path:        "/api/v1/saved/{id}",
		Summary:     "Get saved by guardian ID",
		Description: "Returns all saved events with the given guardian ID",
		Tags:        []string{"Saved"},
	}, func(ctx context.Context, input *models.GetSavedInput) (*models.GetSavedOutput, error) {

		page := input.Page
		if page == 0 {
			page = 1
		}
		limit := input.PageSize
		if limit == 0 {
			limit = 10
		}

		pagination := utils.Pagination{
			Page:  page,
			Limit: limit,
		}

		saveds, err := savedHandler.SavedRepository.GetByGuardianID(ctx, input.ID, pagination)
		if err != nil {
			return nil, err
		}

		return &models.GetSavedOutput{
			Body: saveds,
		}, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "delete-saved",
		Method:      http.MethodDelete,
		Path:        "/api/v1/saved/{id}",
		Summary:     "Delete an existing saved by id",
		Description: "Deletes an existing saved by id",
		Tags:        []string{"Saved"},
	}, func(ctx context.Context, input *models.DeleteSavedInput) (*models.DeleteSavedOutput, error) {
		msg, err := savedHandler.DeleteSaved(ctx, input)
		if err != nil {
			return nil, err
		}

		return &models.DeleteSavedOutput{
			Body: struct {
				Message string `json:"message" doc:"Success message"`
			}{
				Message: msg,
			},
		}, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "create-saved",
		Method:      http.MethodPost,
		Path:        "/api/v1/saved",
		Summary:     "Creates a saved event",
		Description: "Creates a saved event",
		Tags:        []string{"Saved"},
	}, func(ctx context.Context, input *models.CreateSavedInput) (*models.CreateSavedOutput, error) {

		savedOutput, err := savedHandler.CreateSaved(ctx, input)
		if err != nil {
			return nil, err
		}

		return savedOutput, nil
	})
}

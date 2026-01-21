package routes

import (
	"context"
	"net/http"
	"skillspark/internal/models"
	"skillspark/internal/service/handler/school"
	"skillspark/internal/storage"
	"skillspark/internal/utils"

	"github.com/danielgtaylor/huma/v2"
)

func SetupSchoolsRoutes(api huma.API, repo *storage.Repository) {
	schoolHandler := school.NewHandler(repo.School)

	huma.Register(api, huma.Operation{
		OperationID: "get-all-schools",
		Method:      http.MethodGet,
		Path:        "/api/v1/schools",
		Summary:     "Get all schools",
		Description: "Returns all schools",
		Tags:        []string{"Schools"},
	}, func(ctx context.Context, input *models.GetAllSchoolsInput) (*models.GetAllSchoolsOutput, error) {
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

		schools, err := schoolHandler.GetAllSchools(ctx, pagination)
		if err != nil {
			return nil, err
		}

		return &models.GetAllSchoolsOutput{
			Body: schools,
		}, nil
	})
}

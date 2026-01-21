package routes

import (
	"context"
	"net/http"
	"skillspark/internal/service/handler/organization"
	"skillspark/internal/models"
	"skillspark/internal/storage"

	"github.com/danielgtaylor/huma/v2"
)

func SetupOrganizationRoutes(api huma.API, repo *storage.Repository) {
	orgHandler := organization.NewHandler(repo.Organization, repo.Location)

	huma.Register(api, huma.Operation{
		OperationID: "create-organization",
		Method:      http.MethodPost,
		Path:        "/api/v1/organizations",
		Summary:     "Create a new organization",
		Description: "Creates a new organization with the provided information",
		Tags:        []string{"Organizations"},
	}, func(ctx context.Context, input *models.CreateOrganizationInput) (*models.CreateOrganizationOutput, error) {
		return orgHandler.CreateOrganization(ctx, input)
	})

	huma.Register(api, huma.Operation{
		OperationID: "get-organization",
		Method:      http.MethodGet,
		Path:        "/api/v1/organizations/{id}",
		Summary:     "Get organization by ID",
		Description: "Returns a single organization by their ID",
		Tags:        []string{"Organizations"},
	}, func(ctx context.Context, input *models.GetOrganizationByIDInput) (*models.GetOrganizationByIDOutput, error) {
		return orgHandler.GetOrganizationById(ctx, input)
	})

	huma.Register(api, huma.Operation{
		OperationID: "list-organizations",
		Method:      http.MethodGet,
		Path:        "/api/v1/organizations",
		Summary:     "List all organizations",
		Description: "Returns a paginated list of organizations with optional filtering",
		Tags:        []string{"Organizations"},
	}, func(ctx context.Context, input *models.GetAllOrganizationsInput) (*models.GetAllOrganizationsOutput, error) {
		return orgHandler.GetAllOrganizations(ctx, input)
	})

	huma.Register(api, huma.Operation{
		OperationID: "update-organization",
		Method:      http.MethodPatch,
		Path:        "/api/v1/organizations/{id}",
		Summary:     "Update an organization",
		Description: "Updates an existing organization with the provided fields (partial update)",
		Tags:        []string{"Organizations"},
	}, func(ctx context.Context, input *models.UpdateOrganizationInput) (*models.UpdateOrganizationOutput, error) {
		return orgHandler.UpdateOrganization(ctx, input)
	})

	huma.Register(api, huma.Operation{
		OperationID: "delete-organization",
		Method:      http.MethodDelete,
		Path:        "/api/v1/organizations/{id}",
		Summary:     "Delete an organization",
		Description: "Deletes an organization by ID",
		Tags:        []string{"Organizations"},
	}, func(ctx context.Context, input *models.DeleteOrganizationInput) (*models.DeleteOrganizationOutput, error) {
		return orgHandler.DeleteOrganization(ctx, input)
	})
}
package routes

import (
	"context"
	"io"
	"net/http"
	"skillspark/internal/models"
	"skillspark/internal/s3_client"
	"skillspark/internal/service/handler/organization"
	"skillspark/internal/storage"
	translations "skillspark/internal/translation"
	"skillspark/internal/utils"

	"encoding/json"

	"github.com/danielgtaylor/huma/v2"
)

func SetupOrganizationRoutes(api huma.API, repo *storage.Repository, s3Client s3_client.S3Interface, translateClient translations.TranslationInterface) {
	orgHandler := organization.NewHandler(repo.Organization, repo.Location, repo.Review, s3Client, translateClient)

	huma.Register(api, huma.Operation{
		OperationID: "create-organization",
		Method:      http.MethodPost,
		Path:        "/api/v1/organizations",
		Summary:     "Create a new organization",
		Description: "Creates a new organization with the provided information",
		Tags:        []string{"Organizations"},
	}, func(ctx context.Context, input *models.CreateOrganizationRouteInput) (*models.CreateOrganizationOutput, error) {

		formData := input.RawBody.Data()

		var links []models.OrgLink

		if formData.Links != "" {
			if err := json.Unmarshal([]byte(formData.Links), &links); err != nil {
				return nil, err
			}
		}

		var about *string
		if formData.About != "" {
			about = &formData.About
		}

		organizationBody := models.CreateOrganizationBody{
			Name:       formData.Name,
			About:      about,
			Active:     &formData.Active,
			LocationID: &formData.LocationID,
			Links:      links,
		}

		organizationModel := models.CreateOrganizationInput{
			AcceptLanguage: input.AcceptLanguage,
			Body:           organizationBody,
		}

		updateBody := models.UpdateOrganizationBody{
			Name:       &formData.Name,
			About:      about,
			Active:     &formData.Active,
			LocationID: &formData.LocationID,
		}

		var image_data []byte
		var err error
		if formData.ProfileImage.Size > 0 {
			image_data, err = io.ReadAll(formData.ProfileImage)
			if err != nil {
				return nil, err
			}
		}

		organization, err := orgHandler.CreateOrganization(ctx, &organizationModel, &updateBody, &image_data, s3Client)

		if err != nil {
			return nil, err
		}

		return &models.CreateOrganizationOutput{
			Body: *organization,
		}, nil

	})

	huma.Register(api, huma.Operation{
		OperationID: "get-organization",
		Method:      http.MethodGet,
		Path:        "/api/v1/organizations/{id}",
		Summary:     "Get organization by ID",
		Description: "Returns a single organization by their ID",
		Tags:        []string{"Organizations"},
	}, func(ctx context.Context, input *models.GetOrganizationByIDInput) (*models.GetOrganizationByIDOutput, error) {

		organization, err := orgHandler.GetOrganizationById(ctx, input, s3Client)
		if err != nil {
			return nil, err
		}

		return &models.GetOrganizationByIDOutput{
			Body: *organization,
		}, nil

	})

	huma.Register(api, huma.Operation{
		OperationID: "list-organizations",
		Method:      http.MethodGet,
		Path:        "/api/v1/organizations",
		Summary:     "List all organizations",
		Description: "Returns a paginated list of organizations",
		Tags:        []string{"Organizations"},
	}, func(ctx context.Context, input *models.GetAllOrganizationsInput) (*models.GetAllOrganizationsOutput, error) {
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

		organizations, err := orgHandler.GetAllOrganizations(ctx, pagination, input.AcceptLanguage, s3Client)
		if err != nil {
			return nil, err
		}

		return &models.GetAllOrganizationsOutput{
			Body: organizations,
		}, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "update-organization",
		Method:      http.MethodPatch,
		Path:        "/api/v1/organizations/{id}",
		Summary:     "Update an organization",
		Description: "Updates an existing organization with the provided fields (partial update)",
		Tags:        []string{"Organizations"},
	}, func(ctx context.Context, input *models.UpdateOrganizationRouteInput) (*models.UpdateOrganizationOutput, error) {
		formData := input.RawBody.Data()

		var links []models.OrgLink

		if formData.Links != "" {
			if err := json.Unmarshal([]byte(formData.Links), &links); err != nil {
				return nil, err
			}
		}

		var about *string
		if formData.About != "" {
			about = &formData.About
		}

		organizationBody := models.UpdateOrganizationBody{
			Name:       &formData.Name,
			About:      about,
			Active:     &formData.Active,
			LocationID: &formData.LocationID,
			Links:      &links,
		}

		organizationModel := models.UpdateOrganizationInput{
			AcceptLanguage: input.AcceptLanguage,
			Body:           organizationBody,
			ID:             input.ID,
		}

		var image_data []byte
		var err error
		if formData.ProfileImage.Size > 0 {
			image_data, err = io.ReadAll(formData.ProfileImage)
			if err != nil {
				return nil, err
			}
		}

		organization, err := orgHandler.UpdateOrganization(ctx, &organizationModel, &image_data, s3Client)

		if err != nil {
			return nil, err
		}

		return &models.UpdateOrganizationOutput{
			Body: *organization,
		}, nil
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

	huma.Register(api, huma.Operation{
		OperationID: "get-event-occurrences-by-organization-id",
		Method:      http.MethodGet,
		Path:        "/api/v1/organizations/{organization_id}/event-occurrences/",
		Summary:     "Get event occurrences by organization ID",
		Description: "Returns event occurrences that match the organization ID",
		Tags:        []string{"Organizations"},
	}, func(ctx context.Context, input *models.GetEventOccurrencesByOrganizationIDInput) (*models.GetEventOccurrencesByOrganizationIDOutput, error) {
		eventOccurrences, err := orgHandler.GetEventOccurrencesByOrganizationID(ctx, input)
		if err != nil {
			return nil, err
		}

		return &models.GetEventOccurrencesByOrganizationIDOutput{
			Body: eventOccurrences,
		}, nil
	})
}

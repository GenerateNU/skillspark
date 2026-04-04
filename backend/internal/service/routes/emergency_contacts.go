package routes

import (
	"context"
	"fmt"
	"net/http"
	"skillspark/internal/models"
	emergencycontact "skillspark/internal/service/handler/emergency-contact"
	"skillspark/internal/storage"

	"github.com/danielgtaylor/huma/v2"
)

func SetupEmergencyContactRoutes(api huma.API, repo *storage.Repository) {

	emergencyContactHandler := emergencycontact.NewHandler(repo.EmergencyContact)

	huma.Register(api, huma.Operation{
		OperationID: "get-emergency-contacts-by-guardian-id",
		Method:      http.MethodGet,
		Path:        "/api/v1/emergency-contact/{guardian_id}",
		Summary:     "Get emergency contacts by guardian id",
		Description: "Returns emergency contacts by guardian id",
		Tags:        []string{"Emergency Contacts"},
	}, func(ctx context.Context, input *models.GetEmergencyContactByGuardianIDInput) (*models.GetEmergencyContactByGuardianIDOutput, error) {
		emergencyContact, err := emergencyContactHandler.GetEmergencyContactByGuardianID(ctx, input.GuardianID)
		if err != nil {
			return nil, err
		}

		return &models.GetEmergencyContactByGuardianIDOutput{
			Body: emergencyContact,
		}, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "delete-emergency-contact",
		Method:      http.MethodDelete,
		Path:        "/api/v1/emergency-contact/{id}",
		Summary:     "Delete an emergency contact",
		Description: "Deletes an emergency contact",
		Tags:        []string{"Emergency Contacts"},
	}, func(ctx context.Context, input *models.DeleteEmergencyContactInput) (*models.DeleteEmergencyContactOutput, error) {
		_, err := emergencyContactHandler.DeleteEmergencyContact(ctx, input.ID)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		return &models.DeleteEmergencyContactOutput{
			SuccessMessage: "nice",
		}, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "update-emergency-contact",
		Method:      http.MethodPatch,
		Path:        "/api/v1/emergency-contact/{id}",
		Summary:     "Updates an emergency contact",
		Description: "Update an emergency contact",
		Tags:        []string{"Emergency Contacts"},
	}, func(ctx context.Context, input *models.UpdateEmergencyContactInput) (*models.UpdateEmergencyContactOutput, error) {
		UpdateEmergencyContact, err := emergencyContactHandler.UpdateEmergencyContact(ctx, input)
		if err != nil {
			return nil, err
		}

		return &models.UpdateEmergencyContactOutput{
			Body: UpdateEmergencyContact.Body,
		}, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "create-emergency-contact",
		Method:      http.MethodPost,
		Path:        "/api/v1/emergency-contact",
		Summary:     "Creates an emergency contact",
		Description: "Creates an emergency contact",
		Tags:        []string{"Emergency Contacts"},
	}, func(ctx context.Context, input *models.CreateEmergencyContactInput) (*models.CreateEmergencyContactOutput, error) {
		emergencyContact, err := emergencyContactHandler.CreateEmergencyContact(ctx, input)
		if err != nil {

			return nil, err
		}

		return &models.CreateEmergencyContactOutput{
			Body: emergencyContact.Body,
		}, nil
	})
}

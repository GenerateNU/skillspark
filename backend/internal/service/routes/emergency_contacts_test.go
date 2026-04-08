package routes_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"skillspark/internal/models"
	"skillspark/internal/service/routes"
	"skillspark/internal/storage"
	repomocks "skillspark/internal/storage/repo-mocks"
	"testing"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupEmergencyContactTestAPI(
	emergencyContactRepo *repomocks.MockEmergencyContactRepository,
) (*fiber.App, huma.API) {
	app := fiber.New()
	api := humafiber.New(app, huma.DefaultConfig("Test Emergency Contacts API", "1.0.0"))
	repo := &storage.Repository{
		EmergencyContact: emergencyContactRepo,
	}
	routes.SetupEmergencyContactRoutes(api, repo)
	return app, api
}

func TestGetEmergencyContactByGuardianID_Success(t *testing.T) {
	t.Parallel()

	mockRepo := new(repomocks.MockEmergencyContactRepository)

	guardianID := uuid.New()
	now := time.Now()

	expected := []*models.EmergencyContact{
		{
			ID:          uuid.New(),
			Name:        "Jane Doe",
			GuardianID:  guardianID,
			PhoneNumber: "+16462996961",
			CreatedAt:   now,
			UpdatedAt:   now,
		},
	}

	mockRepo.On(
		"GetEmergencyContactByGuardianID",
		mock.Anything,
		guardianID,
	).Return(expected, nil)

	app, _ := setupEmergencyContactTestAPI(mockRepo)

	req, err := http.NewRequest(
		http.MethodGet,
		"/api/v1/emergency-contact/"+guardianID.String(),
		nil,
	)
	assert.NoError(t, err)

	resp, err := app.Test(req)
	assert.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var decoded []*models.EmergencyContact
	err = json.NewDecoder(resp.Body).Decode(&decoded)
	assert.NoError(t, err)

	assert.Len(t, decoded, 1)
	assert.Equal(t, expected[0].GuardianID, decoded[0].GuardianID)
	assert.Equal(t, expected[0].Name, decoded[0].Name)

	mockRepo.AssertExpectations(t)
}

func TestCreateEmergencyContact_Success(t *testing.T) {
	t.Parallel()

	mockRepo := new(repomocks.MockEmergencyContactRepository)

	guardianID := uuid.New()
	contactID := uuid.New()
	now := time.Now()

	expectedContact := &models.EmergencyContact{
		ID:          contactID,
		Name:        "Jane Doe",
		GuardianID:  guardianID,
		PhoneNumber: "+16462996961",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	mockRepo.On(
		"CreateEmergencyContact",
		mock.Anything,
		mock.AnythingOfType("*models.CreateEmergencyContactInput"),
	).Return(&models.CreateEmergencyContactOutput{Body: expectedContact}, nil)

	app, _ := setupEmergencyContactTestAPI(mockRepo)

	body, _ := json.Marshal(map[string]interface{}{
		"name":         "Jane Doe",
		"guardian_id":  guardianID.String(),
		"phone_number": "+16462996961",
	})

	req, err := http.NewRequest(
		http.MethodPost,
		"/api/v1/emergency-contact",
		bytes.NewReader(body),
	)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var decoded models.EmergencyContact
	err = json.NewDecoder(resp.Body).Decode(&decoded)
	assert.NoError(t, err)

	assert.Equal(t, expectedContact.GuardianID, decoded.GuardianID)
	assert.Equal(t, expectedContact.Name, decoded.Name)

	mockRepo.AssertExpectations(t)
}

func TestUpdateEmergencyContact_Success(t *testing.T) {
	t.Parallel()

	mockRepo := new(repomocks.MockEmergencyContactRepository)

	contactID := uuid.New()
	guardianID := uuid.New()
	now := time.Now()

	expectedContact := &models.EmergencyContact{
		ID:          contactID,
		Name:        "Updated Name",
		GuardianID:  guardianID,
		PhoneNumber: "+16462996962",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	mockRepo.On(
		"UpdateEmergencyContact",
		mock.Anything,
		mock.AnythingOfType("*models.UpdateEmergencyContactInput"),
	).Return(&models.UpdateEmergencyContactOutput{Body: expectedContact}, nil)

	app, _ := setupEmergencyContactTestAPI(mockRepo)

	body, _ := json.Marshal(map[string]interface{}{
		"name":         "Updated Name",
		"guardian_id":  guardianID.String(),
		"phone_number": "+16462996962",
	})

	req, err := http.NewRequest(
		http.MethodPatch,
		"/api/v1/emergency-contact/"+contactID.String(),
		bytes.NewReader(body),
	)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	mockRepo.AssertExpectations(t)
}

func TestDeleteEmergencyContact_Success(t *testing.T) {
	t.Parallel()

	mockRepo := new(repomocks.MockEmergencyContactRepository)

	contactID := uuid.New()

	mockRepo.On(
		"DeleteEmergencyContact",
		mock.Anything,
		contactID,
	).Return(&models.DeleteEmergencyContactOutput{Body: &models.DeleteEmergencyContactBody{SuccessMessage: "nice"}}, nil)

	app, _ := setupEmergencyContactTestAPI(mockRepo)

	req, err := http.NewRequest(
		http.MethodDelete,
		"/api/v1/emergency-contact/"+contactID.String(),
		nil,
	)
	assert.NoError(t, err)

	resp, err := app.Test(req)
	assert.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	mockRepo.AssertExpectations(t)
}

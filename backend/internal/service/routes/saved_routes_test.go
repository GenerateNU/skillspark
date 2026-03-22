package routes_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"skillspark/internal/models"
	"skillspark/internal/service/routes"
	"skillspark/internal/storage"
	repomocks "skillspark/internal/storage/repo-mocks"
	"skillspark/internal/utils"
	"testing"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupSavedTestAPI(
	savedRepo *repomocks.MockSavedRepository,
) (*fiber.App, huma.API) {
	app := fiber.New()
	api := humafiber.New(app, huma.DefaultConfig("Test Schools API", "1.0.0"))
	repo := &storage.Repository{
		Saved: savedRepo,
	}
	routes.SetUpSavedRoutes(api, repo)
	return app, api
}

func TestGetSavedByGuardianID_Success(t *testing.T) {
	t.Parallel()

	mockRepo := new(repomocks.MockSavedRepository)

	guardianID := uuid.New()
	now := time.Now()

	eight := 8
	twelve := 12
	jpg := "events/robotics_workshop.jpg"

	event := models.Event{
		ID:               uuid.MustParse("60000000-0000-0000-0000-000000000001"),
		Title:            "Junior Robotics Workshop",
		Description:      "Learn the basics of robotics with hands-on LEGO Mindstorms projects. Build and program your own robots!",
		OrganizationID:   uuid.MustParse("40000000-0000-0000-0000-000000000001"),
		AgeRangeMin:      &eight,
		AgeRangeMax:      &twelve,
		Category:         []string{"science", "technology"},
		HeaderImageS3Key: &jpg,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	expectedSaved := []models.Saved{
		{
			ID:         uuid.New(),
			GuardianID: guardianID,
			Event:      event,
			CreatedAt:  now,
		},
	}

	mockRepo.On(
		"GetByGuardianID",
		mock.Anything,
		guardianID,
		utils.Pagination{Page: 1, Limit: 10},
	).Return(expectedSaved, nil)

	app, _ := setupSavedTestAPI(mockRepo)

	req, err := http.NewRequest(
		http.MethodGet,
		"/api/v1/saved/"+guardianID.String(),
		nil,
	)
	assert.NoError(t, err)

	resp, err := app.Test(req)
	assert.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var decoded []models.Saved
	err = json.NewDecoder(resp.Body).Decode(&decoded)
	assert.NoError(t, err)

	assert.Len(t, decoded, 1)
	assert.Equal(t, expectedSaved[0].GuardianID, decoded[0].GuardianID)
	assert.Equal(t, expectedSaved[0].Event.ID, decoded[0].Event.ID)

	mockRepo.AssertExpectations(t)
}

func TestGetSavedByGuardianID_WithPagination(t *testing.T) {
	t.Parallel()

	mockRepo := new(repomocks.MockSavedRepository)

	guardianID := uuid.New()

	eight := 8
	twelve := 12
	jpg := "events/robotics_workshop.jpg"

	event := models.Event{
		ID:               uuid.MustParse("60000000-0000-0000-0000-000000000001"),
		Title:            "Junior Robotics Workshop",
		Description:      "Learn the basics of robotics with hands-on LEGO Mindstorms projects. Build and program your own robots!",
		OrganizationID:   uuid.MustParse("40000000-0000-0000-0000-000000000001"),
		AgeRangeMin:      &eight,
		AgeRangeMax:      &twelve,
		Category:         []string{"science", "technology"},
		HeaderImageS3Key: &jpg,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	expectedSaved := []models.Saved{
		{
			ID:         uuid.New(),
			GuardianID: guardianID,
			Event:      event,
		},
		{
			ID:         uuid.New(),
			GuardianID: guardianID,
			Event:      event,
		},
	}

	mockRepo.On(
		"GetByGuardianID",
		mock.Anything,
		guardianID,
		utils.Pagination{Page: 2, Limit: 10},
	).Return(expectedSaved, nil)

	app, _ := setupSavedTestAPI(mockRepo)

	req, err := http.NewRequest(
		http.MethodGet,
		"/api/v1/saved/"+guardianID.String()+"?page=2&limit=5",
		nil,
	)
	assert.NoError(t, err)

	resp, err := app.Test(req)
	assert.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var decoded []models.Saved
	err = json.NewDecoder(resp.Body).Decode(&decoded)
	assert.NoError(t, err)

	assert.Len(t, decoded, 2)

	mockRepo.AssertExpectations(t)
}

func TestDeleteSaved_Success(t *testing.T) {
	t.Parallel()

	mockRepo := new(repomocks.MockSavedRepository)

	savedID := uuid.New()

	mockRepo.On(
		"DeleteSaved",
		mock.Anything,
		savedID,
	).Return(nil)

	app, _ := setupSavedTestAPI(mockRepo)

	req, err := http.NewRequest(
		http.MethodDelete,
		"/api/v1/saved/"+savedID.String(),
		nil,
	)
	assert.NoError(t, err)

	resp, err := app.Test(req)
	assert.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var decoded struct {
		Message string `json:"message"`
	}

	err = json.NewDecoder(resp.Body).Decode(&decoded)
	assert.NoError(t, err)

	assert.NotEmpty(t, decoded.Message)

	mockRepo.AssertExpectations(t)
}

func TestCreateSaved_Success(t *testing.T) {
	t.Parallel()

	mockRepo := new(repomocks.MockSavedRepository)

	guardianID := uuid.New()
	eventID := uuid.New()

	eight := 8
	twelve := 12
	jpg := "events/robotics_workshop.jpg"
	event := models.Event{
		ID:               eventID,
		Title:            "Junior Robotics Workshop",
		Description:      "Learn the basics of robotics with hands-on LEGO Mindstorms projects. Build and program your own robots!",
		OrganizationID:   uuid.MustParse("40000000-0000-0000-0000-000000000001"),
		AgeRangeMin:      &eight,
		AgeRangeMax:      &twelve,
		Category:         []string{"science", "technology"},
		HeaderImageS3Key: &jpg,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	input := models.CreateSavedInput{
		Body: struct {
			GuardianID uuid.UUID `json:"guardian_id" db:"guardian_id" doc:"ID of the guardian that saved this."`
			EventID    uuid.UUID `json:"event_id" db:"event_id" doc:"ID of this saved event."`
		}{
			GuardianID: guardianID,
			EventID:    eventID,
		},
	}

	expectedSaved := &models.Saved{
		ID:         uuid.New(),
		GuardianID: guardianID,
		Event:      event,
	}

	mockRepo.On(
		"CreateSaved",
		mock.Anything,
		mock.Anything,
	).Return(expectedSaved, nil)

	app, _ := setupSavedTestAPI(mockRepo)

	body, _ := json.Marshal(input.Body)

	req, err := http.NewRequest(
		http.MethodPost,
		"/api/v1/saved",
		bytes.NewReader(body),
	)
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	defer func() { _ = resp.Body.Close() }()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var decoded models.Saved
	err = json.NewDecoder(resp.Body).Decode(&decoded)
	assert.NoError(t, err)

	assert.Equal(t, expectedSaved.GuardianID, decoded.GuardianID)
	assert.Equal(t, expectedSaved.Event.ID, decoded.Event.ID)

	mockRepo.AssertExpectations(t)
}

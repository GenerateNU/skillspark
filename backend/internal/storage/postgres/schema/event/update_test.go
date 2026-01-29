package event

import (
	"context"
	"testing"

	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// -------------------- Junior Robotics Workshop --------------------
func TestEventRepository_Update_JuniorRoboticsWorkshop(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewEventRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	createInput := func() *models.CreateEventInput {
		input := &models.CreateEventInput{}
		ageMin := 8
		ageMax := 12
		headerImage := "events/robotics_workshop.jpg"

		input.Body.Title = "Junior Robotics Workshop"
		input.Body.Description = "Learn the basics of robotics with hands-on LEGO Mindstorms projects. Build and program your own robots!"
		input.Body.OrganizationID = uuid.MustParse("40000000-0000-0000-0000-000000000001")
		input.Body.AgeRangeMin = &ageMin
		input.Body.AgeRangeMax = &ageMax
		input.Body.Category = []string{"science", "technology"}
		input.Body.HeaderImageS3Key = &headerImage

		return input
	}()

	createdEvent, err := repo.CreateEvent(ctx, createInput)
	assert.Nil(t, err)
	assert.NotNil(t, createdEvent)

	updateInput := &models.UpdateEventInput{}
	updateInput.ID = createdEvent.ID

	newTitle := "Advanced Robotics Workshop"
	newDescription := "Learn the basics of robotics."
	newOrgID := createdEvent.OrganizationID
	newAgeMin := 10
	newAgeMax := 14
	newCategory := []string{"science", "technology"}
	imageKey := "events/robotics_workshop.jpg"

	updateInput.Body.Title = &newTitle
	updateInput.Body.Description = &newDescription
	updateInput.Body.OrganizationID = &newOrgID
	updateInput.Body.AgeRangeMin = &newAgeMin
	updateInput.Body.AgeRangeMax = &newAgeMax
	updateInput.Body.Category = &newCategory
	updateInput.Body.HeaderImageS3Key = &imageKey

	updatedEvent, err := repo.UpdateEvent(ctx, updateInput)
	assert.Nil(t, err)
	assert.NotNil(t, updatedEvent)

	assert.Equal(t, createdEvent.ID, updatedEvent.ID)

	// Verify Changed Fields
	assert.Equal(t, "Advanced Robotics Workshop", updatedEvent.Title)
	assert.Equal(t, 10, *updatedEvent.AgeRangeMin)
	assert.Equal(t, 14, *updatedEvent.AgeRangeMax)

	// Verify Unchanged Fields
	assert.Equal(t, "Learn the basics of robotics.", updatedEvent.Description)
	assert.Equal(t, "events/robotics_workshop.jpg", *updatedEvent.HeaderImageS3Key)

	assert.True(t, updatedEvent.UpdatedAt.After(createdEvent.CreatedAt) || updatedEvent.UpdatedAt.Equal(createdEvent.CreatedAt))
}

func TestEventRepository_Update_NotFound(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewEventRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	updateInput := &models.UpdateEventInput{}
	updateInput.ID = uuid.New()

	title := "Non-existent Event"
	orgID := uuid.New()

	updateInput.Body.Title = &title
	updateInput.Body.OrganizationID = &orgID

	updatedEvent, err := repo.UpdateEvent(ctx, updateInput)
	assert.Nil(t, updatedEvent)
	assert.NotNil(t, err)
}

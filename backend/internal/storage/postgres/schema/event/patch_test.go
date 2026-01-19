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
func TestEventRepository_Patch_JuniorRoboticsWorkshop(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewEventRepository(testDB)
	ctx := context.Background()

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

	patchInput := &models.PatchEventInput{}
	patchInput.ID = createdEvent.ID

	newAgeMin := 10
	newAgeMax := 14
	imageKey := "events/robotics_workshop.jpg"

	patchInput.Body.Title = "Advanced Robotics Workshop"
	patchInput.Body.Description = "Learn the basics of robotics."
	patchInput.Body.OrganizationID = createdEvent.OrganizationID
	patchInput.Body.AgeRangeMin = &newAgeMin
	patchInput.Body.AgeRangeMax = &newAgeMax
	patchInput.Body.Category = []string{"science", "technology", "engineering"}
	patchInput.Body.HeaderImageS3Key = &imageKey

	updatedEvent, err := repo.PatchEvent(ctx, patchInput)
	assert.Nil(t, err)
	assert.NotNil(t, updatedEvent)

	assert.Equal(t, createdEvent.ID, updatedEvent.ID)

	// Verify Changed Fields
	assert.Equal(t, "Advanced Robotics Workshop", updatedEvent.Title)
	assert.Equal(t, 10, *updatedEvent.AgeRangeMin)
	assert.Equal(t, 14, *updatedEvent.AgeRangeMax)
	assert.Contains(t, updatedEvent.Category, "engineering")

	// Verify Unchanged Fields
	assert.Equal(t, "Learn the basics of robotics.", updatedEvent.Description)
	assert.Equal(t, "events/robotics_workshop.jpg", *updatedEvent.HeaderImageS3Key)

	assert.True(t, updatedEvent.UpdatedAt.After(createdEvent.CreatedAt) || updatedEvent.UpdatedAt.Equal(createdEvent.CreatedAt))
}

func TestEventRepository_Patch_NotFound(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewEventRepository(testDB)
	ctx := context.Background()

	patchInput := &models.PatchEventInput{}
	patchInput.ID = uuid.New()
	patchInput.Body.Title = "Non-existent Event"
	patchInput.Body.OrganizationID = uuid.New()

	updatedEvent, err := repo.PatchEvent(ctx, patchInput)
	assert.Nil(t, updatedEvent)
	assert.NotNil(t, err)
}

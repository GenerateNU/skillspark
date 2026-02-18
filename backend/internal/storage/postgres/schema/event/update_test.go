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

	createInput := func() *models.CreateEventDBInput {
		input := &models.CreateEventDBInput{}
		ageMin := 8
		ageMax := 12
		title_ptr := "เวิร์คช็อปหุ่นยนต์สำหรับเด็ก"
		desc_ptr := "เรียนรู้พื้นฐานหุ่นยนต์ด้วยโครงการ LEGO Mindstorms สร้างและเขียนโปรแกรมหุ่นยนต์ของคุณเอง!"

		input.Body.Title_EN = "Junior Robotics Workshop"
		input.Body.Description_EN = "Learn the basics of robotics with hands-on LEGO Mindstorms projects. Build and program your own robots!"
		input.Body.Title_TH = &title_ptr
		input.Body.Description_TH = &desc_ptr
		input.Body.OrganizationID = uuid.MustParse("40000000-0000-0000-0000-000000000001")
		input.Body.AgeRangeMin = &ageMin
		input.Body.AgeRangeMax = &ageMax
		input.Body.Category = []string{"science", "technology"}

		return input
	}()

	headerImage := "events/robotics_workshop.jpg"
	createdEvent, err := repo.CreateEvent(ctx, createInput, &headerImage)
	assert.Nil(t, err)
	assert.NotNil(t, createdEvent)

	updateInput := &models.UpdateEventDBInput{}
	updateInput.ID = createdEvent.ID

	newTitle := "Advanced Robotics Workshop"
	newDescription := "Learn the basics of robotics."
	newOrgID := createdEvent.OrganizationID
	newAgeMin := 10
	newAgeMax := 14
	newCategory := []string{"science", "technology"}

	updateInput.Body.Title_EN = &newTitle
	updateInput.Body.Description_EN = &newDescription
	updateInput.Body.OrganizationID = &newOrgID
	updateInput.Body.AgeRangeMin = &newAgeMin
	updateInput.Body.AgeRangeMax = &newAgeMax
	updateInput.Body.Category = &newCategory

	imageKey := "events/robotics_workshop.jpg"
	updatedEvent, err := repo.UpdateEvent(ctx, updateInput, &imageKey)
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

	updateInput := &models.UpdateEventDBInput{}
	updateInput.ID = uuid.New()

	title := "Non-existent Event"
	orgID := uuid.New()

	updateInput.Body.Title_EN = &title
	updateInput.Body.OrganizationID = &orgID

	updatedEvent, err := repo.UpdateEvent(ctx, updateInput, nil)
	assert.Nil(t, updatedEvent)
	assert.NotNil(t, err)
}

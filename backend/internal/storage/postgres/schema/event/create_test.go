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
func TestEventRepository_Create_JuniorRoboticsWorkshop(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewEventRepository(testDB)
	ctx := context.Background()

	eventInput := func() *models.CreateEventInput {
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

	event, err := repo.CreateEvent(ctx, eventInput)
	assert.Nil(t, err)
	assert.NotNil(t, event)
	assert.Equal(t, "Junior Robotics Workshop", event.Title)
	assert.Equal(t, "Learn the basics of robotics with hands-on LEGO Mindstorms projects. Build and program your own robots!", event.Description)
	assert.Equal(t, uuid.MustParse("40000000-0000-0000-0000-000000000001"), event.OrganizationID)
	assert.NotNil(t, event.AgeRangeMin)
	assert.Equal(t, 8, *event.AgeRangeMin)
	assert.NotNil(t, event.AgeRangeMax)
	assert.Equal(t, 12, *event.AgeRangeMax)
	assert.Equal(t, []string{"science", "technology"}, event.Category)
	assert.NotNil(t, event.HeaderImageS3Key)
	assert.Equal(t, "events/robotics_workshop.jpg", *event.HeaderImageS3Key)
}

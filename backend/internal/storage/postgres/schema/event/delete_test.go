package event

import (
	"context"
	"testing"

	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestEventRepository_Delete_JuniorRoboticsWorkshop(t *testing.T) {
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

		input.Body.Title = "Junior Robotics Workshop"
		input.Body.Description = "Learn the basics of robotics with hands-on LEGO Mindstorms projects. Build and program your own robots!"
		input.Body.OrganizationID = uuid.MustParse("40000000-0000-0000-0000-000000000001")
		input.Body.AgeRangeMin = &ageMin
		input.Body.AgeRangeMax = &ageMax
		input.Body.Category = []string{"science", "technology"}

		return input
	}()

	headerImage := "events/robotics_workshop.jpg"
	event, err := repo.CreateEvent(ctx, eventInput, &headerImage)
	assert.Nil(t, err)
	assert.NotNil(t, event)

	delErr := repo.DeleteEvent(ctx, event.ID)
	assert.Nil(t, delErr)

	delErr2 := repo.DeleteEvent(ctx, event.ID)
	assert.NotNil(t, delErr2, "Expected error when deleting an already deleted event")
}

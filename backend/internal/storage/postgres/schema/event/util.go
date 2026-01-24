package event

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/organization"
	"skillspark/internal/storage/postgres/testutil"
	"testing"
)

func CreateTestEvent(
	t *testing.T,
	ctx context.Context,
) *models.Event {
	t.Helper()

	testDB := testutil.SetupTestDB(t)
	repo := NewEventRepository(testDB)

	organization := organization.CreateTestOrganization(t, ctx)

	input := &models.CreateEventInput{}
	ageMin := 8
	ageMax := 12
	headerImage := "events/robotics_workshop.jpg"

	input.Body.Title = "Junior Robotics Workshop"
	input.Body.Description = "Learn the basics of robotics with hands-on LEGO Mindstorms projects. Build and program your own robots!"
	input.Body.OrganizationID = organization.ID
	input.Body.AgeRangeMin = &ageMin
	input.Body.AgeRangeMax = &ageMax
	input.Body.Category = []string{"science", "technology"}
	input.Body.HeaderImageS3Key = &headerImage

	event, _ := repo.CreateEvent(ctx, input)
	return event
}

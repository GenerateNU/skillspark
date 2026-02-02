package event

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/organization"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

func CreateTestEvent(
	t *testing.T,
	ctx context.Context,
	db *pgxpool.Pool,
) *models.Event {
	t.Helper()

	repo := NewEventRepository(db)

	organization := organization.CreateTestOrganization(t, ctx, db)

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

	event, err := repo.CreateEvent(ctx, input)

	require.NoError(t, err)
	require.NotNil(t, event)
	return event
}

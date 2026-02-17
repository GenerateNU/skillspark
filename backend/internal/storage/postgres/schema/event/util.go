package event

import (
	"context"
	"embed"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/organization"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

//go:embed sql/*.sql
var SqlEventFiles embed.FS

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

	input.Body.Title = "Junior Robotics Workshop"
	input.Body.Description = "Learn the basics of robotics with hands-on LEGO Mindstorms projects. Build and program your own robots!"
	input.Body.OrganizationID = organization.ID
	input.Body.AgeRangeMin = &ageMin
	input.Body.AgeRangeMax = &ageMax
	input.Body.Category = []string{"science", "technology"}

	event, err := repo.CreateEvent(ctx, input, nil)

	require.NoError(t, err)
	require.NotNil(t, event)
	return event
}

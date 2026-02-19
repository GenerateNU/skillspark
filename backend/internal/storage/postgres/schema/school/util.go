package school

import (
	"context"
	"embed"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/location"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

//go:embed sql/*.sql
var SqlSchoolFiles embed.FS

func CreateTestSchool(
	t *testing.T,
	ctx context.Context,
	db *pgxpool.Pool,
) *models.School {
	t.Helper()

	repo := NewSchoolRepository(db)

	dummyLocation := location.CreateTestLocation(t, ctx, db)

	input := &models.CreateSchoolInput{}
	input.Body.Name = "Monster High School"
	input.Body.LocationID = dummyLocation.ID

	school, err := repo.CreateSchool(ctx, input)

	require.NoError(t, err)
	require.NotNil(t, school)

	return school
}

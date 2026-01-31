package school

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/location"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

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

	school, _ := repo.CreateSchool(ctx, input)
	return school
}

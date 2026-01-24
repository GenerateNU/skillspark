package school

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/location"
	"skillspark/internal/storage/postgres/testutil"
	"testing"
)

func CreateTestSchool(
	t *testing.T,
	ctx context.Context,
) *models.School {
	t.Helper()

	testDB := testutil.SetupTestDB(t)
	repo := NewSchoolRepository(testDB)

	dummyLocation := location.CreateTestLocation(t, ctx)

	input := &models.CreateSchoolInput{}
	input.Body.Name = "Monster High School"
	input.Body.LocationID = dummyLocation.ID

	school, _ := repo.CreateSchool(ctx, input)
	return school
}

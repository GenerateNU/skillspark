package school

import (
	"context"
	"skillspark/internal/storage/postgres/testutil"
	"skillspark/internal/utils"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSchoolRepository_GetAllSchools(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewSchoolRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	// Use existing seeded locations for referential integrity
	newYorkLocationID := uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	bostonLocationID := uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a19")

	// Insert some test schools
	nySchoolID := uuid.New()
	bostonSchoolID := uuid.New()
	_, err := testDB.Exec(ctx, `
		INSERT INTO school (id, name, location_id, created_at, updated_at)
		VALUES 
			($1, $2, $3, NOW(), NOW()),
			($4, $5, $6, NOW(), NOW());
	`, nySchoolID, "NY High School", newYorkLocationID,
		bostonSchoolID, "Boston High School", bostonLocationID)
	if err != nil {
		t.Fatalf("failed to insert test schools: %v", err)
	}

	// Test with default pagination
	pagination := utils.NewPagination()
	schools, httpErr := repo.GetAllSchools(ctx, pagination)

	assert.Nil(t, httpErr)
	assert.NotNil(t, schools)
	assert.GreaterOrEqual(t, len(schools), 2)

	// Map by ID to avoid relying on ordering
	found := make(map[uuid.UUID]struct{})
	for _, s := range schools {
		found[s.ID] = struct{}{}
		if s.ID == nySchoolID {
			assert.Equal(t, "NY High School", s.Name)
			assert.Equal(t, newYorkLocationID, s.LocationID)
		}
		if s.ID == bostonSchoolID {
			assert.Equal(t, "Boston High School", s.Name)
			assert.Equal(t, bostonLocationID, s.LocationID)
		}
	}

	_, hasNY := found[nySchoolID]
	_, hasBoston := found[bostonSchoolID]
	assert.True(t, hasNY, "expected NY High School to be returned")
	assert.True(t, hasBoston, "expected Boston High School to be returned")
}

func TestSchoolRepository_GetAllSchools_Pagination(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewSchoolRepository(testDB)
	ctx := context.Background()

	t.Parallel()

	newYorkLocationID := uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")

	// Insert 5 test schools
	schoolIDs := make([]uuid.UUID, 5)
	for i := 0; i < 5; i++ {
		schoolIDs[i] = uuid.New()
		_, err := testDB.Exec(ctx, `
			INSERT INTO school (id, name, location_id, created_at, updated_at)
			VALUES ($1, $2, $3, NOW(), NOW());
		`, schoolIDs[i], "Test School "+string(rune('A'+i)), newYorkLocationID)
		if err != nil {
			t.Fatalf("failed to insert test school: %v", err)
		}
	}

	// Test page 1 with limit 2
	pagination := utils.Pagination{Page: 1, Limit: 2}
	schools, httpErr := repo.GetAllSchools(ctx, pagination)
	assert.Nil(t, httpErr)
	assert.Equal(t, 2, len(schools))

	// Test page 2 with limit 2
	pagination = utils.Pagination{Page: 2, Limit: 2}
	schools, httpErr = repo.GetAllSchools(ctx, pagination)
	assert.Nil(t, httpErr)
	assert.Equal(t, 2, len(schools))

	// Test page 3 with limit 2 (should have remaining schools)
	pagination = utils.Pagination{Page: 3, Limit: 2}
	schools, httpErr = repo.GetAllSchools(ctx, pagination)
	assert.Nil(t, httpErr)
	assert.GreaterOrEqual(t, len(schools), 1)
}

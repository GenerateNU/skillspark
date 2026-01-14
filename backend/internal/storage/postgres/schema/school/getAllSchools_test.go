package school

import (
	"context"
	"testing"

	"skillspark/internal/storage/postgres/testutil"

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

	schools, httpErr := repo.GetAllSchools(ctx)

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

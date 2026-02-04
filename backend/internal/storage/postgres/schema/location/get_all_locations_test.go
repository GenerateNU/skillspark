package location

import (
	"context"
	"skillspark/internal/storage/postgres/testutil"
	"skillspark/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocationRepository_GetAllLocations(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewLocationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	t.Run("get all locations with default pagination", func(t *testing.T) {
		pagination := utils.Pagination{
			Limit: 10,
			Page:  1,
		}

		locations, err := repo.GetAllLocations(ctx, pagination)

		assert.Nil(t, err)
		assert.NotNil(t, locations)
		assert.GreaterOrEqual(t, len(locations), 1) // Expect at least one seeded location

		// Verify fields of the first location (assuming consistent seeding order or just checking one)
		// We just check required fields are populated
		for _, loc := range locations {
			assert.NotEmpty(t, loc.ID)
			assert.NotEmpty(t, loc.District)
			assert.NotEmpty(t, loc.Country)
		}
	})

	t.Run("get all locations with limit 1", func(t *testing.T) {
		pagination := utils.Pagination{
			Limit: 1,
			Page:  1,
		}

		locations, err := repo.GetAllLocations(ctx, pagination)

		assert.Nil(t, err)
		assert.Len(t, locations, 1)
	})
}

package location

import (
	"context"
	"testing"

	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"

	"github.com/stretchr/testify/assert"
)

// -------------------- New York --------------------
func TestLocationRepository_Create_NewYork(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewLocationRepository(testDB)
	ctx := context.Background()

	input := &models.CreateLocationInput{}
	input.Body.Latitude = 40.7128
	input.Body.Longitude = -74.0060
	input.Body.Address = "123 Broadway"
	input.Body.City = "New York"
	input.Body.State = "NY"
	input.Body.ZipCode = "10001"
	input.Body.Country = "USA"

	loc, err := repo.CreateLocation(ctx, input)
	assert.Nil(t, err)
	assert.NotNil(t, loc)
	assert.Equal(t, "New York", loc.City)
	assert.Equal(t, "NY", loc.State)
	assert.Equal(t, "10001", loc.ZipCode)
	assert.Equal(t, "USA", loc.Country)
	assert.Equal(t, "123 Broadway", loc.Address)

	// Verify retrievability
	retrieved, err := repo.GetLocationByID(ctx, loc.ID)
	assert.Nil(t, err)
	assert.Equal(t, loc.ID, retrieved.ID)
}

// -------------------- Boston --------------------
func TestLocationRepository_Create_Boston(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewLocationRepository(testDB)
	ctx := context.Background()

	input := &models.CreateLocationInput{}
	input.Body.Latitude = 42.3601
	input.Body.Longitude = -71.0589
	input.Body.Address = "600 Boylston Street"
	input.Body.City = "Boston"
	input.Body.State = "MA"
	input.Body.ZipCode = "02116"
	input.Body.Country = "USA"

	loc, err := repo.CreateLocation(ctx, input)
	assert.Nil(t, err)
	assert.NotNil(t, loc)
	assert.Equal(t, "Boston", loc.City)
	assert.Equal(t, "MA", loc.State)
	assert.Equal(t, "02116", loc.ZipCode)
	assert.Equal(t, "USA", loc.Country)
	assert.Equal(t, "600 Boylston Street", loc.Address)

	retrieved, err := repo.GetLocationByID(ctx, loc.ID)
	assert.Nil(t, err)
	assert.Equal(t, loc.ID, retrieved.ID)
}

// -------------------- San Francisco --------------------
func TestLocationRepository_Create_SF(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewLocationRepository(testDB)
	ctx := context.Background()

	input := &models.CreateLocationInput{}
	input.Body.Latitude = 37.7749
	input.Body.Longitude = -122.4194
	input.Body.Address = "700 Market Street"
	input.Body.City = "San Francisco"
	input.Body.State = "CA"
	input.Body.ZipCode = "94102"
	input.Body.Country = "USA"

	loc, err := repo.CreateLocation(ctx, input)
	assert.Nil(t, err)
	assert.NotNil(t, loc)
	assert.Equal(t, "San Francisco", loc.City)
	assert.Equal(t, "CA", loc.State)
	assert.Equal(t, "94102", loc.ZipCode)
	assert.Equal(t, "USA", loc.Country)
	assert.Equal(t, "700 Market Street", loc.Address)

	retrieved, err := repo.GetLocationByID(ctx, loc.ID)
	assert.Nil(t, err)
	assert.Equal(t, loc.ID, retrieved.ID)
}

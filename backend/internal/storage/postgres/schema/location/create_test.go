package location

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocationRepository_Create(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestWithCleanup(t)
	repo := NewLocationRepository(testDB)
	ctx := context.Background()

	locationInput := func() *models.CreateLocationInput {
		input := &models.CreateLocationInput{}
		input.Body.Latitude = 40.7128
		input.Body.Longitude = -74.0060
		input.Body.Address = "123 Broadway"
		input.Body.City = "New York"
		input.Body.State = "NY"
		input.Body.ZipCode = "10001"
		input.Body.Country = "USA"
		return input
	}()

	location, err := repo.CreateLocation(ctx, locationInput)

	assert.Nil(t, err)
	assert.NotNil(t, location)
	assert.Equal(t, "New York", location.City)
	assert.Equal(t, "NY", location.State)
	assert.Equal(t, "10001", location.ZipCode)
	assert.Equal(t, "USA", location.Country)
	assert.Equal(t, "123 Broadway", location.Address)

	id := location.ID

	// Verify we can retrieve the created location
	retrievedLocation, err := repo.GetLocationByID(ctx, id)

	assert.Nil(t, err)
	assert.NotNil(t, retrievedLocation)
	assert.Equal(t, location.ID, retrievedLocation.ID)
	assert.Equal(t, location.City, retrievedLocation.City)
	assert.Equal(t, location.State, retrievedLocation.State)
	assert.Equal(t, location.ZipCode, retrievedLocation.ZipCode)
	assert.Equal(t, location.Country, retrievedLocation.Country)
	assert.Equal(t, location.Address, retrievedLocation.Address)
}

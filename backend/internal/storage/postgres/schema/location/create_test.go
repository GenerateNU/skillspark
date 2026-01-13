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
		input.Body.StreetNumber = "123"
		input.Body.StreetName = "Broadway"
		input.Body.SecondaryAddress = nil
		input.Body.City = "New York"
		input.Body.State = "NY"
		input.Body.PostalCode = "10001"
		input.Body.Country = "USA"
		return input
	}()

	location, err := repo.CreateLocation(ctx, locationInput)

	assert.Nil(t, err)
	assert.NotNil(t, location)
	assert.Equal(t, "New York", location.City)
	assert.Equal(t, "NY", location.State)
	assert.Equal(t, "10001", location.PostalCode)
	assert.Equal(t, "USA", location.Country)
	assert.Equal(t, "123", location.StreetNumber)
	assert.Equal(t, "Broadway", location.StreetName)
	assert.Equal(t, nil, location.SecondaryAddress)

	id := location.ID

	// Verify we can retrieve the created location
	retrievedLocation, err := repo.GetLocationByID(ctx, id)

	assert.Nil(t, err)
	assert.NotNil(t, retrievedLocation)
	assert.Equal(t, location.ID, retrievedLocation.ID)
	assert.Equal(t, location.City, retrievedLocation.City)
	assert.Equal(t, location.State, retrievedLocation.State)
	assert.Equal(t, location.PostalCode, retrievedLocation.PostalCode)
	assert.Equal(t, location.Country, retrievedLocation.Country)
	assert.Equal(t, location.StreetNumber, retrievedLocation.StreetNumber)
	assert.Equal(t, location.StreetName, retrievedLocation.StreetName)
	assert.Equal(t, location.SecondaryAddress, retrievedLocation.SecondaryAddress)
}

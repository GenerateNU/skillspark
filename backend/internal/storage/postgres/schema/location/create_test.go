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

	locationInput := func() *models.CreateLocationInput {
		input := &models.CreateLocationInput{}
		input.Body.Latitude = 40.7128
		input.Body.Longitude = -74.0060
		input.Body.AddressLine1 = "123 Broadway"
		input.Body.AddressLine2 = nil
		input.Body.Subdistrict = "Manhattan"
		input.Body.District = "New York County"
		input.Body.Province = "NY"
		input.Body.PostalCode = "10001"
		input.Body.Country = "USA"
		return input
	}()

	location, err := repo.CreateLocation(ctx, locationInput)
	assert.Nil(t, err)
	assert.NotNil(t, location)
	assert.Equal(t, "10001", location.PostalCode)
	assert.Equal(t, "USA", location.Country)
	assert.Equal(t, "123 Broadway", location.AddressLine1)
	assert.Nil(t, location.AddressLine2)

	id := location.ID

	// Verify we can retrieve the created location
	retrievedLocation, err := repo.GetLocationByID(ctx, id)
	assert.Nil(t, err)
	assert.NotNil(t, retrievedLocation)
	assert.Equal(t, location.ID, retrievedLocation.ID)
	assert.Equal(t, location.Subdistrict, retrievedLocation.Subdistrict)
	assert.Equal(t, location.District, retrievedLocation.District)
	assert.Equal(t, location.PostalCode, retrievedLocation.PostalCode)
	assert.Equal(t, location.Country, retrievedLocation.Country)
	assert.Equal(t, location.AddressLine1, retrievedLocation.AddressLine1)
	assert.Nil(t, retrievedLocation.AddressLine2)
}

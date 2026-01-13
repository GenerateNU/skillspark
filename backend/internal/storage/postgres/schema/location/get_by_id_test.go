package location

import (
	"context"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestLocationRepository_GetLocationByID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestWithCleanup(t)
	repo := NewLocationRepository(testDB)
	ctx := context.Background()

	location, err := repo.GetLocationByID(ctx, uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"))

	assert.Nil(t, err)
	assert.NotNil(t, location)
	assert.Equal(t, "New York", location.City)
	assert.Equal(t, "NY", location.State)
	assert.Equal(t, "10001", location.PostalCode)
	assert.Equal(t, "USA", location.Country)
	assert.Equal(t, "123", location.StreetNumber)
	assert.Equal(t, "Broadway", location.StreetName)

	assert.Nil(t, location.SecondaryAddress)
	assert.Equal(t, "New York", location.City)
	assert.Equal(t, "NY", location.State)
	assert.Equal(t, "10001", location.PostalCode)
	assert.Equal(t, "USA", location.Country)

	location2, err := repo.GetLocationByID(ctx, uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a19"))

	assert.Nil(t, err)
	assert.NotNil(t, location2)
	assert.Equal(t, "Boston", location2.City)
	assert.Equal(t, "MA", location2.State)
	assert.Equal(t, "02116", location2.PostalCode)
	assert.Equal(t, "USA", location2.Country)
	assert.Equal(t, "600", location2.StreetNumber)
	assert.Equal(t, "Boylston Street", location2.StreetName)
	assert.Nil(t, location2.SecondaryAddress)

	location3, err := repo.GetLocationByID(ctx, uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a20"))

	assert.Nil(t, err)
	assert.NotNil(t, location3)
	assert.Equal(t, "San Francisco", location3.City)
	assert.Equal(t, "CA", location3.State)
	assert.Equal(t, "94102", location3.PostalCode)
	assert.Equal(t, "USA", location3.Country)
	assert.Equal(t, "700", location3.StreetNumber)
	assert.Equal(t, "Market Street", location3.StreetName)
}

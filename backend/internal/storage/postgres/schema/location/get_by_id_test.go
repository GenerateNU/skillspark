package location

import (
	"context"
	"testing"

	"skillspark/internal/storage/postgres/testutil"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestLocationRepository_GetLocationByID_NewYork(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewLocationRepository(testDB)
	ctx := context.Background()

	location, err := repo.GetLocationByID(ctx, uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"))

	assert.Nil(t, err)
	assert.NotNil(t, location)
	assert.Equal(t, "New York", location.City)
	assert.Equal(t, "NY", location.State)
	assert.Equal(t, "10001", location.ZipCode)
	assert.Equal(t, "USA", location.Country)
	assert.Equal(t, "123 Broadway", location.Address)
}

func TestLocationRepository_GetLocationByID_Boston(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewLocationRepository(testDB)
	ctx := context.Background()

	location, err := repo.GetLocationByID(ctx, uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a19"))

	assert.Nil(t, err)
	assert.NotNil(t, location)
	assert.Equal(t, "Boston", location.City)
	assert.Equal(t, "MA", location.State)
	assert.Equal(t, "02116", location.ZipCode)
	assert.Equal(t, "USA", location.Country)
	assert.Equal(t, "600 Boylston Street", location.Address)
}

func TestLocationRepository_GetLocationByID_SF(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewLocationRepository(testDB)
	ctx := context.Background()

	location, err := repo.GetLocationByID(ctx, uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a20"))

	assert.Nil(t, err)
	assert.NotNil(t, location)
	assert.Equal(t, "San Francisco", location.City)
	assert.Equal(t, "CA", location.State)
	assert.Equal(t, "94102", location.ZipCode)
	assert.Equal(t, "USA", location.Country)
	assert.Equal(t, "700 Market Street", location.Address)
}

func TestLocationRepository_GetLocationByID_NotFound(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewLocationRepository(testDB)
	ctx := context.Background()

	invalidID := uuid.New()
	location, err := repo.GetLocationByID(ctx, invalidID)

	assert.Nil(t, location)
	assert.NotNil(t, err)
	assert.Equal(t, "http error: 404 Location with id='"+invalidID.String()+"' not found", err.Error()) // adapt if your NotFound message differs
}

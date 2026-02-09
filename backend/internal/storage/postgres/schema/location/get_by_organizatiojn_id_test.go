package location

import (
	"context"
	"testing"

	"skillspark/internal/storage/postgres/testutil"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestLocationRepository_GetLocationByOrganizationID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewLocationRepository(testDB)
	ctx := context.Background()

	// get location for Science Academy Bangkok
	location, err := repo.GetLocationByOrganizationID(ctx, uuid.MustParse("40000000-0000-0000-0000-000000000001"))
	if err != nil {
		t.Fatalf("Failed to get location by organization id: %v", err)
	}

	assert.NotNil(t, location)
	assert.Equal(t, uuid.MustParse("10000000-0000-0000-0000-000000000004"), location.ID)
	assert.Equal(t, "10400", location.PostalCode)
	assert.Equal(t, "Thailand", location.Country)
	assert.Equal(t, "321 Phetchaburi Road", location.AddressLine1)
	assert.Equal(t, "Suite 15", *location.AddressLine2)
	assert.Equal(t, "Ratchathewi", location.District)
	assert.NotNil(t, location.CreatedAt)
	assert.NotNil(t, location.UpdatedAt)

	// get location for Champions Sports Center
	location2, err := repo.GetLocationByOrganizationID(ctx, uuid.MustParse("40000000-0000-0000-0000-000000000002"))
	if err != nil {
		t.Fatalf("Failed to get location by organization id: %v", err)
	}

	assert.NotNil(t, location2)
	assert.Equal(t, uuid.MustParse("10000000-0000-0000-0000-000000000005"), location2.ID)
	assert.Equal(t, "10120", location2.PostalCode)
	assert.Equal(t, "Thailand", location2.Country)
	assert.Equal(t, "654 Sathorn Road", location2.AddressLine1)
	assert.Nil(t, location2.AddressLine2)
	assert.NotNil(t, location2.CreatedAt)
	assert.NotNil(t, location2.UpdatedAt)

	// get location for Creative Arts Studio
	location3, err := repo.GetLocationByOrganizationID(ctx, uuid.MustParse("40000000-0000-0000-0000-000000000003"))
	if err != nil {
		t.Fatalf("Failed to get location by organization id: %v", err)
	}

	assert.NotNil(t, location3)
	assert.Equal(t, uuid.MustParse("10000000-0000-0000-0000-000000000006"), location3.ID)
	assert.Equal(t, "10230", location3.PostalCode)
	assert.Equal(t, "Thailand", location3.Country)
	assert.Equal(t, "147 Lat Phrao Road", location3.AddressLine1)
	assert.Equal(t, "Building C", *location3.AddressLine2)

	// get location for organization that does not exist
	location4, err := repo.GetLocationByOrganizationID(ctx, uuid.MustParse("40000000-0000-0000-0000-000000000099"))
	assert.Nil(t, location4)
	assert.NotNil(t, err)
}
package location

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"skillspark/internal/errs"
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
	assert.Equal(t, "10001", location.PostalCode)
	assert.Equal(t, "USA", location.Country)
	assert.Equal(t, "123 Broadway", location.AddressLine1)
	assert.Nil(t, location.AddressLine2)
	assert.Equal(t, "10001", location.PostalCode)
	assert.Equal(t, "USA", location.Country)

	location2, err := repo.GetLocationByID(ctx, uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a19"))

	assert.Nil(t, err)
	assert.NotNil(t, location2)
	assert.Equal(t, "02116", location2.PostalCode)
	assert.Equal(t, "USA", location2.Country)
	assert.Equal(t, "456 Boylston St", location2.AddressLine1)
	assert.Nil(t, location2.AddressLine2)

	location3, err := repo.GetLocationByID(ctx, uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a20"))

	assert.Nil(t, err)
	assert.NotNil(t, location3)
	assert.Equal(t, "94102", location3.PostalCode)
	assert.Equal(t, "USA", location3.Country)
	assert.Equal(t, "789 Market St", location3.AddressLine1)
	assert.Nil(t, location3.AddressLine2)

	location4, err := repo.GetLocationByID(ctx, uuid.MustParse("a0111c99-9c0b-4ef8-bb6d-6bb9bd380a20"))

	assert.NotNil(t, err)
	assert.Nil(t, location4)
	fmt.Println(err)
	fmt.Println(errs.NotFound("Location", "id", uuid.MustParse("a0111c99-9c0b-4ef8-bb6d-6bb9bd380a20")))
	assert.Equal(t, http.StatusNotFound, err.(*errs.HTTPError).Code)
	assert.Contains(t, err.(*errs.HTTPError).Message, "Location")
	assert.Contains(t, err.(*errs.HTTPError).Message, "a0111c99-9c0b-4ef8-bb6d-6bb9bd380a20")
}

package location

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"
	"testing"
)

func CreateTestLocation(
	t *testing.T,
	ctx context.Context,
) *models.Location {
	t.Helper()

	testDB := testutil.SetupTestDB(t)
	repo := NewLocationRepository(testDB)

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

	location, _ := repo.CreateLocation(ctx, input)
	return location
}

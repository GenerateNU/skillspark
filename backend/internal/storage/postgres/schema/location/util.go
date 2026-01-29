package location

import (
	"context"
	"skillspark/internal/models"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateTestLocation(
	t *testing.T,
	ctx context.Context,
	db *pgxpool.Pool,
) *models.Location {
	t.Helper()

	repo := NewLocationRepository(db)

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

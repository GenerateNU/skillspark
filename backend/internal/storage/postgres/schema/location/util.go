package location

import (
	"context"
	"embed"
	"skillspark/internal/models"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

//go:embed sql/*.sql
var SqlLocationFiles embed.FS

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

	location, err := repo.CreateLocation(ctx, input)

	require.NoError(t, err)
	require.NotNil(t, location)

	return location

}

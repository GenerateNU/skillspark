package organization

import (
	"context"
	"embed"
	"encoding/json"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/location"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

//go:embed sql/*.sql
var SqlOrganizationFiles embed.FS

func CreateTestOrganization(
	t *testing.T,
	ctx context.Context,
	db *pgxpool.Pool,
) *models.Organization {
	t.Helper()

	repo := NewOrganizationRepository(db)
	testLocation := location.CreateTestLocation(t, ctx, db)

	active := true
	locationID := testLocation.ID
	i := &models.CreateOrganizationDBInput{
		AcceptLanguage: "en-US",
		Body: models.CreateOrgDBBody{
			Name:       "Test Corp",
			Active:     &active,
			LocationID: &locationID,
		},
	}

	organization, err := repo.CreateOrganization(ctx, i, nil)
	require.NoError(t, err)
	require.NotNil(t, organization)

	return organization
}

func scanLinks(raw []byte) ([]models.OrgLink, error) {
	var links []models.OrgLink
	if raw == nil {
		return []models.OrgLink{}, nil
	}
	if err := json.Unmarshal(raw, &links); err != nil {
		return nil, err
	}
	return links, nil
}

func byteSliceLinks(links []models.OrgLink) ([]byte, error) {
	if links == nil {
		return []byte("[]"), nil
	}
	return json.Marshal(links)
}

// pickAbout returns the about column matching AcceptLanguage, falling back
// to the other column if the requested one is NULL. Returns nil if both are NULL.
func pickAbout(AcceptLanguage string, aboutEN, aboutTH *string) *string {
	switch AcceptLanguage {
	case "th-TH":
		if aboutTH != nil {
			return aboutTH
		}
		return aboutEN
	default:
		if aboutEN != nil {
			return aboutEN
		}
		return aboutTH
	}
}

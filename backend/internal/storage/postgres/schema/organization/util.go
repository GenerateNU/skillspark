package organization

import (
	"context"
	"embed"
	"encoding/json"
	"skillspark/internal/models"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
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

	active := true
	i := &models.CreateOrganizationInput{}
	i.Body.Name = "Test Corp"
	i.Body.Active = &active

	organization, _ := repo.CreateOrganization(ctx, i, nil)

	// t.Logf("error returned: %v", err)

	// require.NoError(t, err)
	// require.NotNil(t, organization)

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

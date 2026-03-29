package organization

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateOrganization_WithLinks(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	active := true
	links := []models.OrgLink{
		{Href: "https://example.com", Label: "Website"},
		{Href: "https://instagram.com/test", Label: "Instagram"},
	}
	input := &models.CreateOrganizationInput{}
	input.Body.Name = "Links Corp"
	input.Body.Active = &active
	input.Body.Links = links

	created, err := repo.CreateOrganization(ctx, input, nil)

	require.Nil(t, err)
	require.NotNil(t, created)
	assert.Equal(t, 2, len(created.Links))
	assert.Equal(t, "https://example.com", created.Links[0].Href)
	assert.Equal(t, "Website", created.Links[0].Label)
}

func TestCreateOrganization_DefaultsToEmptyLinks(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	active := true
	input := &models.CreateOrganizationInput{}
	input.Body.Name = "No Links Corp"
	input.Body.Active = &active

	created, err := repo.CreateOrganization(ctx, input, nil)

	require.Nil(t, err)
	require.NotNil(t, created)
	assert.NotNil(t, created.Links)
	assert.Equal(t, 0, len(created.Links))
}

func TestGetOrganizationByID_ReturnsLinks(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	active := true
	links := []models.OrgLink{
		{Href: "https://example.com", Label: "Website"},
	}
	input := &models.CreateOrganizationInput{}
	input.Body.Name = "Get Links Corp"
	input.Body.Active = &active
	input.Body.Links = links

	created, err := repo.CreateOrganization(ctx, input, nil)
	require.Nil(t, err)

	fetched, err := repo.GetOrganizationByID(ctx, created.ID)

	require.Nil(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, 1, len(fetched.Links))
	assert.Equal(t, "https://example.com", fetched.Links[0].Href)
}

func TestUpdateOrganization_UpdatesLinks(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	org := CreateTestOrganization(t, ctx, testDB)

	newLinks := []models.OrgLink{
		{Href: "https://new.com", Label: "New Site"},
	}
	input := &models.UpdateOrganizationInput{}
	input.ID = org.ID
	input.Body.Links = &newLinks

	updated, err := repo.UpdateOrganization(ctx, input, nil)

	require.Nil(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, 1, len(updated.Links))
	assert.Equal(t, "https://new.com", updated.Links[0].Href)
}

func TestGetEventOccurrencesByOrganizationID_IncludesLinks(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	// Use seeded org that already has event occurrences
	orgID := uuid.MustParse("40000000-0000-0000-0000-000000000001")

	// First set links on the org via update
	links := []models.OrgLink{
		{Href: "https://scienceacademy.com", Label: "Website"},
	}
	input := &models.UpdateOrganizationInput{}
	input.ID = orgID
	input.Body.Links = &links
	_, err := repo.UpdateOrganization(ctx, input, nil)
	require.Nil(t, err)

	// Now fetch event occurrences and verify links flow through
	occurrences, err := repo.GetEventOccurrencesByOrganizationID(ctx, orgID, "en-US")

	require.Nil(t, err)
	require.NotEmpty(t, occurrences)
	assert.Equal(t, 1, len(occurrences[0].OrgLinks))
	assert.Equal(t, "https://scienceacademy.com", occurrences[0].OrgLinks[0].Href)
}

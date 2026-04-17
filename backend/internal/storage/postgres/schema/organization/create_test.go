package organization

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/location"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateOrganization(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	active := true
	locationID := location.CreateTestLocation(t, ctx, testDB).ID

	input := &models.CreateOrganizationDBInput{
		AcceptLanguage: "en-US",
		Body: models.CreateOrgDBBody{
			Name:       "Test Corp",
			Active:     &active,
			LocationID: &locationID,
		},
	}

	created, err := repo.CreateOrganization(ctx, input, nil)
	require.Nil(t, err)
	require.NotNil(t, created)
	assert.Equal(t, "Test Corp", created.Name)
	assert.True(t, created.Active)
	assert.NotEqual(t, uuid.Nil, created.ID)
	require.NotNil(t, created.LocationID)
	assert.Equal(t, locationID, *created.LocationID)
	// Verify Stripe fields default correctly
	assert.Nil(t, created.StripeAccountID)
	assert.False(t, created.StripeAccountActivated)
}

func TestCreateOrganization_WithLocation(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	active := true
	locationID := location.CreateTestLocation(t, ctx, testDB).ID

	input := &models.CreateOrganizationDBInput{
		AcceptLanguage: "en-US",
		Body: models.CreateOrgDBBody{
			Name:       "Test Corp with Location",
			Active:     &active,
			LocationID: &locationID,
		},
	}

	created, err := repo.CreateOrganization(ctx, input, nil)
	require.Nil(t, err)
	require.NotNil(t, created)
	assert.Equal(t, "Test Corp with Location", created.Name)
	assert.True(t, created.Active)
	require.NotNil(t, created.LocationID)
	assert.Equal(t, locationID, *created.LocationID)
	assert.Nil(t, created.StripeAccountID)
	assert.False(t, created.StripeAccountActivated)
}

func TestCreateOrganization_WithPfp(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	active := true
	pfpKey := "orgs/test_corp.jpg"
	locationID := location.CreateTestLocation(t, ctx, testDB).ID

	input := &models.CreateOrganizationDBInput{
		AcceptLanguage: "en-US",
		Body: models.CreateOrgDBBody{
			Name:       "Test Corp with Profile",
			Active:     &active,
			LocationID: &locationID,
		},
	}

	created, err := repo.CreateOrganization(ctx, input, &pfpKey)
	require.Nil(t, err)
	require.NotNil(t, created)
	assert.Equal(t, "Test Corp with Profile", created.Name)
	assert.Equal(t, &pfpKey, created.PfpS3Key)
	assert.Nil(t, created.StripeAccountID)
	assert.False(t, created.StripeAccountActivated)
}

func TestCreateOrganization_Inactive(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	active := false
	locationID := location.CreateTestLocation(t, ctx, testDB).ID

	input := &models.CreateOrganizationDBInput{
		AcceptLanguage: "en-US",
		Body: models.CreateOrgDBBody{
			Name:       "Inactive Corp",
			Active:     &active,
			LocationID: &locationID,
		},
	}

	created, err := repo.CreateOrganization(ctx, input, nil)
	require.Nil(t, err)
	require.NotNil(t, created)
	assert.Equal(t, "Inactive Corp", created.Name)
	assert.False(t, created.Active)
	assert.Nil(t, created.StripeAccountID)
	assert.False(t, created.StripeAccountActivated)
}

func TestCreateOrganization_FullDetails(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	active := true
	locationID := location.CreateTestLocation(t, ctx, testDB).ID
	pfpKey := "orgs/full_corp.jpg"

	input := &models.CreateOrganizationDBInput{
		AcceptLanguage: "en-US",
		Body: models.CreateOrgDBBody{
			Name:       "Full Details Corp",
			Active:     &active,
			LocationID: &locationID,
		},
	}

	created, err := repo.CreateOrganization(ctx, input, &pfpKey)
	require.Nil(t, err)
	require.NotNil(t, created)
	assert.Equal(t, "Full Details Corp", created.Name)
	assert.True(t, created.Active)
	assert.Equal(t, &pfpKey, created.PfpS3Key)
	require.NotNil(t, created.LocationID)
	assert.Equal(t, locationID, *created.LocationID)
	assert.Nil(t, created.StripeAccountID)
	assert.False(t, created.StripeAccountActivated)
}

func TestCreateOrganization_BilingualAbout(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	active := true
	locationID := location.CreateTestLocation(t, ctx, testDB).ID
	aboutEN := "A leading tech company"
	aboutTH := "บริษัทเทคโนโลยีชั้นนำ"

	input := &models.CreateOrganizationDBInput{
		AcceptLanguage: "en-US",
		Body: models.CreateOrgDBBody{
			Name:       "Bilingual Corp",
			Active:     &active,
			LocationID: &locationID,
			AboutEN:    &aboutEN,
			AboutTH:    &aboutTH,
		},
	}

	created, err := repo.CreateOrganization(ctx, input, nil)
	require.Nil(t, err)
	require.NotNil(t, created)
	require.NotNil(t, created.About)
	assert.Equal(t, aboutEN, *created.About)

	// Same row, fetched as Thai
	fetched, err := repo.GetOrganizationByID(ctx, created.ID, "th-TH")
	require.Nil(t, err)
	require.NotNil(t, fetched.About)
	assert.Equal(t, aboutTH, *fetched.About)
}

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
	input := func() *models.CreateOrganizationInput {
		i := &models.CreateOrganizationInput{}
		i.Body.Name = "Test Corp"
		i.Body.Active = &active
		return i
	}()

	created, err := repo.CreateOrganization(ctx, input, nil)

	require.Nil(t, err)
	require.NotNil(t, created)
	assert.Equal(t, "Test Corp", created.Name)
	assert.True(t, created.Active)
	assert.NotEqual(t, uuid.Nil, created.ID)
	
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
	input := func() *models.CreateOrganizationInput {
		i := &models.CreateOrganizationInput{}
		i.Body.Name = "Test Corp with Location"
		i.Body.Active = &active
		i.Body.LocationID = &locationID
		return i
	}()

	created, err := repo.CreateOrganization(ctx, input, nil)

	require.Nil(t, err)
	require.NotNil(t, created)
	assert.Equal(t, "Test Corp with Location", created.Name)
	assert.True(t, created.Active)
	assert.Equal(t, &locationID, created.LocationID)
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
	input := func() *models.CreateOrganizationInput {
		i := &models.CreateOrganizationInput{}
		i.Body.Name = "Test Corp with Profile"
		i.Body.Active = &active
		return i
	}()

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
	input := func() *models.CreateOrganizationInput {
		i := &models.CreateOrganizationInput{}
		i.Body.Name = "Inactive Corp"
		i.Body.Active = &active
		return i
	}()

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
	input := func() *models.CreateOrganizationInput {
		i := &models.CreateOrganizationInput{}
		i.Body.Name = "Full Details Corp"
		i.Body.Active = &active
		i.Body.LocationID = &locationID
		return i
	}()

	created, err := repo.CreateOrganization(ctx, input, &pfpKey)

	require.Nil(t, err)
	require.NotNil(t, created)
	assert.Equal(t, "Full Details Corp", created.Name)
	assert.True(t, created.Active)
	assert.Equal(t, &pfpKey, created.PfpS3Key)
	assert.Equal(t, &locationID, created.LocationID)
	assert.Nil(t, created.StripeAccountID)
	assert.False(t, created.StripeAccountActivated)
}
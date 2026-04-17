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

func TestUpdateOrganization(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	active := true
	locationID := location.CreateTestLocation(t, ctx, testDB).ID

	createInput := &models.CreateOrganizationDBInput{
		AcceptLanguage: "en-US",
		Body: models.CreateOrgDBBody{
			Name:       "Original Name",
			Active:     &active,
			LocationID: &locationID,
		},
	}

	created, createErr := repo.CreateOrganization(ctx, createInput, nil)
	require.Nil(t, createErr)
	require.NotNil(t, created)

	newName := "Updated Name"
	newActive := false
	updateInput := &models.UpdateOrganizationDBInput{
		AcceptLanguage: "en-US",
		ID:             created.ID,
		Body: models.UpdateOrgDBBody{
			Name:   &newName,
			Active: &newActive,
		},
	}

	updated, updateErr := repo.UpdateOrganization(ctx, updateInput, nil)
	require.Nil(t, updateErr)
	require.NotNil(t, updated)
	assert.Equal(t, "Updated Name", updated.Name)
	assert.False(t, updated.Active)
	assert.Nil(t, updated.StripeAccountID)
	assert.False(t, updated.StripeAccountActivated)

	fetched, getErr := repo.GetOrganizationByID(ctx, created.ID, "en-US")
	require.Nil(t, getErr)
	assert.Equal(t, "Updated Name", fetched.Name)
	assert.False(t, fetched.Active)
	assert.Nil(t, fetched.StripeAccountID)
	assert.False(t, fetched.StripeAccountActivated)
}

func TestUpdateOrganization_WithLocation(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	active := true
	locationID := location.CreateTestLocation(t, ctx, testDB).ID

	createInput := &models.CreateOrganizationDBInput{
		AcceptLanguage: "en-US",
		Body: models.CreateOrgDBBody{
			Name:       "Test Org",
			Active:     &active,
			LocationID: &locationID,
		},
	}

	created, createErr := repo.CreateOrganization(ctx, createInput, nil)
	require.Nil(t, createErr)
	require.NotNil(t, created)

	newLocationID := location.CreateTestLocation(t, ctx, testDB).ID
	newName := "Test Org with Location"
	updateInput := &models.UpdateOrganizationDBInput{
		AcceptLanguage: "en-US",
		ID:             created.ID,
		Body: models.UpdateOrgDBBody{
			Name:       &newName,
			LocationID: &newLocationID,
		},
	}

	updated, updateErr := repo.UpdateOrganization(ctx, updateInput, nil)
	require.Nil(t, updateErr)
	require.NotNil(t, updated)
	assert.Equal(t, "Test Org with Location", updated.Name)
	require.NotNil(t, updated.LocationID)
	assert.Equal(t, newLocationID, *updated.LocationID)
	assert.Nil(t, updated.StripeAccountID)
	assert.False(t, updated.StripeAccountActivated)
}

func TestUpdateOrganization_NotFound(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	nonExistentID := uuid.New()
	newName := "Does Not Exist"
	updateInput := &models.UpdateOrganizationDBInput{
		AcceptLanguage: "en-US",
		ID:             nonExistentID,
		Body: models.UpdateOrgDBBody{
			Name: &newName,
		},
	}

	updated, err := repo.UpdateOrganization(ctx, updateInput, nil)
	require.NotNil(t, err)
	assert.Nil(t, updated)
}

func TestUpdateOrganization_DoesNotModifyStripeFields(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	active := true
	locationID := location.CreateTestLocation(t, ctx, testDB).ID

	createInput := &models.CreateOrganizationDBInput{
		AcceptLanguage: "en-US",
		Body: models.CreateOrgDBBody{
			Name:       "Stripe Test Org",
			Active:     &active,
			LocationID: &locationID,
		},
	}

	created, createErr := repo.CreateOrganization(ctx, createInput, nil)
	require.Nil(t, createErr)

	newName := "Updated Stripe Org"
	updateInput := &models.UpdateOrganizationDBInput{
		AcceptLanguage: "en-US",
		ID:             created.ID,
		Body: models.UpdateOrgDBBody{
			Name: &newName,
		},
	}

	updated, updateErr := repo.UpdateOrganization(ctx, updateInput, nil)
	require.Nil(t, updateErr)
	assert.Equal(t, "Updated Stripe Org", updated.Name)
	assert.Nil(t, updated.StripeAccountID)
	assert.False(t, updated.StripeAccountActivated)
}

func TestUpdateOrganization_PreservesAboutWhenNotProvided(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	active := true
	locationID := location.CreateTestLocation(t, ctx, testDB).ID
	aboutEN := "Original english"
	aboutTH := "ข้อความไทยต้นฉบับ"

	created, createErr := repo.CreateOrganization(ctx, &models.CreateOrganizationDBInput{
		AcceptLanguage: "en-US",
		Body: models.CreateOrgDBBody{
			Name:       "Preserve About",
			Active:     &active,
			LocationID: &locationID,
			AboutEN:    &aboutEN,
			AboutTH:    &aboutTH,
		},
	}, nil)
	require.Nil(t, createErr)

	newName := "Preserve About Renamed"
	_, updateErr := repo.UpdateOrganization(ctx, &models.UpdateOrganizationDBInput{
		AcceptLanguage: "en-US",
		ID:             created.ID,
		Body: models.UpdateOrgDBBody{
			Name: &newName,
		},
	}, nil)
	require.Nil(t, updateErr)

	// Both language variants should still be present
	fetchedEN, _ := repo.GetOrganizationByID(ctx, created.ID, "en-US")
	fetchedTH, _ := repo.GetOrganizationByID(ctx, created.ID, "th-TH")
	require.NotNil(t, fetchedEN.About)
	require.NotNil(t, fetchedTH.About)
	assert.Equal(t, aboutEN, *fetchedEN.About)
	assert.Equal(t, aboutTH, *fetchedTH.About)
}

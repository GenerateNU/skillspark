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
	createInput := func() *models.CreateOrganizationInput {
		i := &models.CreateOrganizationInput{}
		i.Body.Name = "Original Name"
		i.Body.Active = &active
		return i
	}()

	created, createErr := repo.CreateOrganization(ctx, createInput, nil)
	require.Nil(t, createErr)
	require.NotNil(t, created)

	newName := "Updated Name"
	newActive := false
	updateInput := &models.UpdateOrganizationInput{
		ID: created.ID,
		Body: models.UpdateOrganizationBody{
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

	fetched, getErr := repo.GetOrganizationByID(ctx, created.ID)
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
	createInput := func() *models.CreateOrganizationInput {
		i := &models.CreateOrganizationInput{}
		i.Body.Name = "Test Org"
		i.Body.Active = &active
		return i
	}()

	created, createErr := repo.CreateOrganization(ctx, createInput, nil)
	require.Nil(t, createErr)
	require.NotNil(t, created)

	locationID := location.CreateTestLocation(t, ctx, testDB).ID
	newName := "Test Org with Location"
	updateInput := &models.UpdateOrganizationInput{
		ID: created.ID,
		Body: models.UpdateOrganizationBody{
			Name:       &newName,
			LocationID: &locationID,
		},
	}

	updated, updateErr := repo.UpdateOrganization(ctx, updateInput, nil)
	require.Nil(t, updateErr)
	require.NotNil(t, updated)
	assert.Equal(t, "Test Org with Location", updated.Name)
	assert.Equal(t, &locationID, updated.LocationID)
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
	updateInput := &models.UpdateOrganizationInput{
		ID: nonExistentID,
		Body: models.UpdateOrganizationBody{
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
	createInput := func() *models.CreateOrganizationInput {
		i := &models.CreateOrganizationInput{}
		i.Body.Name = "Stripe Test Org"
		i.Body.Active = &active
		return i
	}()

	created, createErr := repo.CreateOrganization(ctx, createInput, nil)
	require.Nil(t, createErr)

	newName := "Updated Stripe Org"
	updateInput := &models.UpdateOrganizationInput{
		ID: created.ID,
		Body: models.UpdateOrganizationBody{
			Name: &newName,
		},
	}

	updated, updateErr := repo.UpdateOrganization(ctx, updateInput, nil)
	require.Nil(t, updateErr)
	assert.Equal(t, "Updated Stripe Org", updated.Name)
	assert.Nil(t, updated.StripeAccountID)
	assert.False(t, updated.StripeAccountActivated)
}
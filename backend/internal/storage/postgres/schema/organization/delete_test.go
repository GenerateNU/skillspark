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

func TestDeleteOrganization(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	active := true
	input := func() *models.CreateOrganizationInput {
		i := &models.CreateOrganizationInput{}
		i.Body.Name = "To Be Deleted"
		i.Body.Active = &active
		return i
	}()

	created, err := repo.CreateOrganization(ctx, input, nil)
	require.NoError(t, err)
	require.NotNil(t, created)

	deleted, err := repo.DeleteOrganization(ctx, created.ID)
	require.NoError(t, err)
	require.NotNil(t, deleted)
	assert.Equal(t, created.ID, deleted.ID)
	assert.Equal(t, "To Be Deleted", deleted.Name)
	assert.Nil(t, deleted.StripeAccountID)
	assert.False(t, deleted.StripeAccountActivated)

	_, err = repo.GetOrganizationByID(ctx, created.ID)
	assert.Error(t, err)
}

func TestDeleteOrganization_NotFound(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	deleted, err := repo.DeleteOrganization(ctx, uuid.New())

	require.Error(t, err)
	assert.Nil(t, deleted)
}

func TestDeleteOrganization_AlreadyDeleted(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	active := true
	input := func() *models.CreateOrganizationInput {
		i := &models.CreateOrganizationInput{}
		i.Body.Name = "Delete Twice"
		i.Body.Active = &active
		return i
	}()

	created, err := repo.CreateOrganization(ctx, input, nil)
	require.NoError(t, err)
	require.NotNil(t, created)

	_, err = repo.DeleteOrganization(ctx, created.ID)
	require.NoError(t, err)

	deleted2, err := repo.DeleteOrganization(ctx, created.ID)
	require.Error(t, err)
	assert.Nil(t, deleted2)
}

func TestDeleteOrganization_WithStripeAccount(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	testOrg := CreateTestOrganization(t, ctx, testDB)
	stripeAccountID := "acct_delete_test123"

	_, err := repo.SetStripeAccountID(ctx, testOrg.ID, stripeAccountID)
	require.NoError(t, err)
	_, err = repo.SetStripeAccountActivated(ctx, stripeAccountID, true)
	require.NoError(t, err)

	deleted, err := repo.DeleteOrganization(ctx, testOrg.ID)
	require.NoError(t, err)
	require.NotNil(t, deleted)
	assert.Equal(t, stripeAccountID, *deleted.StripeAccountID)
	assert.True(t, deleted.StripeAccountActivated)
}
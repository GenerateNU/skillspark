package organization

import (
	"context"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSetStripeAccountActivated(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	testOrg := CreateTestOrganization(t, ctx, testDB)
	stripeAccountID := "acct_1SwvOB2Sjs4wsi8o"
	
	orgWithAccount, err := repo.SetStripeAccountID(ctx, testOrg.ID, stripeAccountID)
	require.Nil(t, err)
	require.False(t, orgWithAccount.StripeAccountActivated)

	updated, err := repo.SetStripeAccountActivated(ctx, stripeAccountID, true)

	require.Nil(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, testOrg.ID, updated.ID)
	assert.Equal(t, stripeAccountID, *updated.StripeAccountID)
	assert.True(t, updated.StripeAccountActivated)

	fetched, getErr := repo.GetOrganizationByID(ctx, testOrg.ID)
	require.Nil(t, getErr)
	assert.True(t, fetched.StripeAccountActivated)
}

func TestSetStripeAccountActivated_Deactivate(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	testOrg := CreateTestOrganization(t, ctx, testDB)
	stripeAccountID := "acct_deactivate123"
	
	repo.SetStripeAccountID(ctx, testOrg.ID, stripeAccountID)
	repo.SetStripeAccountActivated(ctx, stripeAccountID, true)

	deactivated, err := repo.SetStripeAccountActivated(ctx, stripeAccountID, false)

	require.Nil(t, err)
	require.NotNil(t, deactivated)
	assert.False(t, deactivated.StripeAccountActivated)
}

func TestSetStripeAccountActivated_NotFound(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	nonExistentAccountID := "acct_doesnotexist123"

	updated, err := repo.SetStripeAccountActivated(ctx, nonExistentAccountID, true)

	require.NotNil(t, err)
	assert.Nil(t, updated)
}

func TestSetStripeAccountActivated_DoesNotModifyOtherFields(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	testOrg := CreateTestOrganization(t, ctx, testDB)
	stripeAccountID := "acct_fieldstest123"
	
	orgWithAccount, _ := repo.SetStripeAccountID(ctx, testOrg.ID, stripeAccountID)
	originalName := orgWithAccount.Name
	originalActive := orgWithAccount.Active

	updated, err := repo.SetStripeAccountActivated(ctx, stripeAccountID, true)

	require.Nil(t, err)
	assert.Equal(t, originalName, updated.Name)
	assert.Equal(t, originalActive, updated.Active)
	assert.Equal(t, stripeAccountID, *updated.StripeAccountID)
	assert.True(t, updated.StripeAccountActivated)
}

func TestSetStripeAccountActivated_MultipleToggle(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	testOrg := CreateTestOrganization(t, ctx, testDB)
	stripeAccountID := "acct_toggle123"
	
	repo.SetStripeAccountID(ctx, testOrg.ID, stripeAccountID)

	activated, err := repo.SetStripeAccountActivated(ctx, stripeAccountID, true)
	require.Nil(t, err)
	assert.True(t, activated.StripeAccountActivated)

	deactivated, err := repo.SetStripeAccountActivated(ctx, stripeAccountID, false)
	require.Nil(t, err)
	assert.False(t, deactivated.StripeAccountActivated)

	reactivated, err := repo.SetStripeAccountActivated(ctx, stripeAccountID, true)
	require.Nil(t, err)
	assert.True(t, reactivated.StripeAccountActivated)
}
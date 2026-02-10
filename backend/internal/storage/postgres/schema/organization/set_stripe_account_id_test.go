package organization

import (
	"context"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSetStripeAccountID(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	testOrg := CreateTestOrganization(t, ctx, testDB)
	require.NotNil(t, testOrg)
	require.Nil(t, testOrg.StripeAccountID)

	stripeAccountID := "acct_1SwvOB2Sjs4wsi8o"
	updated, err := repo.SetStripeAccountID(ctx, testOrg.ID, stripeAccountID)

	require.Nil(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, testOrg.ID, updated.ID)
	assert.NotNil(t, updated.StripeAccountID)
	assert.Equal(t, stripeAccountID, *updated.StripeAccountID)
	assert.False(t, updated.StripeAccountActivated)

	fetched, getErr := repo.GetOrganizationByID(ctx, testOrg.ID)
	require.Nil(t, getErr)
	assert.Equal(t, stripeAccountID, *fetched.StripeAccountID)
	assert.False(t, fetched.StripeAccountActivated)
}

func TestSetStripeAccountID_MultipleTimes(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	testOrg := CreateTestOrganization(t, ctx, testDB)

	firstAccountID := "acct_first123"
	updated1, err := repo.SetStripeAccountID(ctx, testOrg.ID, firstAccountID)
	require.Nil(t, err)
	assert.Equal(t, firstAccountID, *updated1.StripeAccountID)

	secondAccountID := "acct_second456"
	updated2, err := repo.SetStripeAccountID(ctx, testOrg.ID, secondAccountID)
	require.Nil(t, err)
	assert.Equal(t, secondAccountID, *updated2.StripeAccountID)
	assert.False(t, updated2.StripeAccountActivated)
}

func TestSetStripeAccountID_NotFound(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	nonExistentID := uuid.New()
	stripeAccountID := "acct_1SwvOB2Sjs4wsi8o"

	updated, err := repo.SetStripeAccountID(ctx, nonExistentID, stripeAccountID)

	require.NotNil(t, err)
	assert.Nil(t, updated)
}

func TestSetStripeAccountID_DoesNotModifyOtherFields(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	testOrg := CreateTestOrganization(t, ctx, testDB)
	originalName := testOrg.Name
	originalActive := testOrg.Active

	stripeAccountID := "acct_1SwvOB2Sjs4wsi8o"
	updated, err := repo.SetStripeAccountID(ctx, testOrg.ID, stripeAccountID)

	require.Nil(t, err)
	assert.Equal(t, originalName, updated.Name)
	assert.Equal(t, originalActive, updated.Active)
	assert.Equal(t, stripeAccountID, *updated.StripeAccountID)
	assert.False(t, updated.StripeAccountActivated)
}
package organization

import (
	"context"
	"skillspark/internal/storage/postgres/testutil"
	"skillspark/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAllOrganizations_BasicPagination(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	pagination := utils.Pagination{Page: 1, Limit: 10}
	orgs, err := repo.GetAllOrganizations(ctx, pagination)

	require.Nil(t, err)
	require.NotNil(t, orgs)
	assert.GreaterOrEqual(t, len(orgs), 3)
	assert.LessOrEqual(t, len(orgs), 10)
	
	if len(orgs) > 0 {
		assert.Nil(t, orgs[0].StripeAccountID)
		assert.False(t, orgs[0].StripeAccountActivated)
	}
}

func TestGetAllOrganizations_SecondPage(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	firstPage := utils.Pagination{Page: 1, Limit: 2}
	firstPageOrgs, err := repo.GetAllOrganizations(ctx, firstPage)
	require.Nil(t, err)
	assert.Equal(t, 2, len(firstPageOrgs))

	secondPage := utils.Pagination{Page: 2, Limit: 2}
	secondPageOrgs, err := repo.GetAllOrganizations(ctx, secondPage)

	require.Nil(t, err)
	require.NotNil(t, secondPageOrgs)
	assert.GreaterOrEqual(t, len(secondPageOrgs), 1)

	if len(secondPageOrgs) > 0 && len(firstPageOrgs) > 0 {
		assert.NotEqual(t, firstPageOrgs[0].ID, secondPageOrgs[0].ID)
		assert.False(t, secondPageOrgs[0].StripeAccountActivated)
	}
}

func TestGetAllOrganizations_SmallPageSize(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	pagination := utils.Pagination{Page: 1, Limit: 2}
	orgs, err := repo.GetAllOrganizations(ctx, pagination)

	require.Nil(t, err)
	require.NotNil(t, orgs)
	assert.Equal(t, 2, len(orgs))
	
	for _, org := range orgs {
		assert.False(t, org.StripeAccountActivated)
	}
}

func TestGetAllOrganizations_SingleItemPerPage(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	pagination := utils.Pagination{Page: 1, Limit: 1}
	orgs, err := repo.GetAllOrganizations(ctx, pagination)

	require.Nil(t, err)
	require.NotNil(t, orgs)
	assert.Equal(t, 1, len(orgs))
	
	assert.Nil(t, orgs[0].StripeAccountID)
	assert.False(t, orgs[0].StripeAccountActivated)
}

func TestGetAllOrganizations_PageBeyondData(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	pagination := utils.Pagination{Page: 100, Limit: 10}
	orgs, err := repo.GetAllOrganizations(ctx, pagination)

	require.Nil(t, err)
	require.NotNil(t, orgs)
	assert.Equal(t, 0, len(orgs))
}

func TestGetAllOrganizations_AllDataOnePage(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	pagination := utils.Pagination{Page: 1, Limit: 100}
	orgs, err := repo.GetAllOrganizations(ctx, pagination)

	require.Nil(t, err)
	require.NotNil(t, orgs)
	assert.GreaterOrEqual(t, len(orgs), 3)
	
	for _, org := range orgs {
		assert.False(t, org.StripeAccountActivated)
	}
}

func TestGetAllOrganizations_OrderByCreatedAt(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	pagination := utils.Pagination{Page: 1, Limit: 10}
	orgs, err := repo.GetAllOrganizations(ctx, pagination)

	require.Nil(t, err)
	assert.GreaterOrEqual(t, len(orgs), 2)

	for i := 0; i < len(orgs)-1; i++ {
		assert.True(t,
			orgs[i].CreatedAt.After(orgs[i+1].CreatedAt) || orgs[i].CreatedAt.Equal(orgs[i+1].CreatedAt),
			"Organizations should be ordered by created_at DESC",
		)
		assert.False(t, orgs[i].StripeAccountActivated)
	}
}

func TestGetAllOrganizations_ZeroOffset(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	repo := NewOrganizationRepository(testDB)
	ctx := context.Background()
	t.Parallel()

	pagination := utils.Pagination{Page: 1, Limit: 3}
	orgs, err := repo.GetAllOrganizations(ctx, pagination)

	require.Nil(t, err)
	require.NotNil(t, orgs)
	assert.Equal(t, 3, len(orgs))
	
	for _, org := range orgs {
		assert.Nil(t, org.StripeAccountID)
		assert.False(t, org.StripeAccountActivated)
	}
}
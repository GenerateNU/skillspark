package getallorganizationspaginated

import (
	"context"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecute_BasicPagination(t *testing.T) {
	testDB := testutil.SetupTestWithCleanup(t)
	ctx := context.Background()

	orgs, totalCount, err := Execute(ctx, testDB, 0, 10, nil, nil)

	assert.Nil(t, err)
	assert.NotNil(t, orgs)
	assert.GreaterOrEqual(t, totalCount, 3) 
	assert.LessOrEqual(t, len(orgs), 10)
}

func TestExecute_SecondPage(t *testing.T) {
	testDB := testutil.SetupTestWithCleanup(t)
	ctx := context.Background()

	firstPageOrgs, totalCount, err := Execute(ctx, testDB, 0, 2, nil, nil)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(firstPageOrgs))

	secondPageOrgs, totalCount2, err := Execute(ctx, testDB, 2, 2, nil, nil)

	assert.Nil(t, err)
	assert.NotNil(t, secondPageOrgs)
	assert.Equal(t, totalCount, totalCount2) 
	assert.GreaterOrEqual(t, len(secondPageOrgs), 1) 
	

	if len(secondPageOrgs) > 0 && len(firstPageOrgs) > 0 {
		assert.NotEqual(t, firstPageOrgs[0].ID, secondPageOrgs[0].ID)
	}
}

func TestExecute_SmallPageSize(t *testing.T) {
	testDB := testutil.SetupTestWithCleanup(t)
	ctx := context.Background()

	orgs, totalCount, err := Execute(ctx, testDB, 0, 2, nil, nil)

	assert.Nil(t, err)
	assert.NotNil(t, orgs)
	assert.Equal(t, 2, len(orgs))
	assert.GreaterOrEqual(t, totalCount, 3)
}

func TestExecute_SingleItemPerPage(t *testing.T) {
	testDB := testutil.SetupTestWithCleanup(t)
	ctx := context.Background()

	orgs, totalCount, err := Execute(ctx, testDB, 0, 1, nil, nil)

	assert.Nil(t, err)
	assert.NotNil(t, orgs)
	assert.Equal(t, 1, len(orgs))
	assert.GreaterOrEqual(t, totalCount, 3)
}

func TestExecute_PageBeyondData(t *testing.T) {
	testDB := testutil.SetupTestWithCleanup(t)
	ctx := context.Background()

	
	orgs, totalCount, err := Execute(ctx, testDB, 1000, 10, nil, nil)

	assert.Nil(t, err)
	assert.NotNil(t, orgs)
	assert.Equal(t, 0, len(orgs))
	assert.GreaterOrEqual(t, totalCount, 3) 
}

func TestExecute_AllDataOnePage(t *testing.T) {
	testDB := testutil.SetupTestWithCleanup(t)
	ctx := context.Background()

	orgs, totalCount, err := Execute(ctx, testDB, 0, 100, nil, nil)

	assert.Nil(t, err)
	assert.NotNil(t, orgs)
	assert.Equal(t, totalCount, len(orgs)) 
	assert.GreaterOrEqual(t, totalCount, 3)
}

func TestExecute_OrderByCreatedAt(t *testing.T) {
	testDB := testutil.SetupTestWithCleanup(t)
	ctx := context.Background()

	orgs, _, err := Execute(ctx, testDB, 0, 10, nil, nil)

	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(orgs), 2)

	for i := 0; i < len(orgs)-1; i++ {
		assert.True(t, 
			orgs[i].CreatedAt.After(orgs[i+1].CreatedAt) || orgs[i].CreatedAt.Equal(orgs[i+1].CreatedAt),
			"Organizations should be ordered by created_at DESC",
		)
	}
}

func TestExecute_ZeroOffset(t *testing.T) {
	testDB := testutil.SetupTestWithCleanup(t)
	ctx := context.Background()

	orgs, totalCount, err := Execute(ctx, testDB, 0, 3, nil, nil)

	assert.Nil(t, err)
	assert.NotNil(t, orgs)
	assert.Equal(t, 3, len(orgs))
	assert.GreaterOrEqual(t, totalCount, 3)
}
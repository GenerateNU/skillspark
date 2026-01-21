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
	assert.Greater(t, totalCount, 0)
	assert.LessOrEqual(t, len(orgs), 10)
}

func TestExecute_SecondPage(t *testing.T) {
	testDB := testutil.SetupTestWithCleanup(t)
	ctx := context.Background()

	orgs, totalCount, err := Execute(ctx, testDB, 10, 10, nil, nil)

	assert.Nil(t, err)
	assert.NotNil(t, orgs)
	assert.Greater(t, totalCount, 0)
}

func TestExecute_SmallPageSize(t *testing.T) {
	testDB := testutil.SetupTestWithCleanup(t)
	ctx := context.Background()

	orgs, totalCount, err := Execute(ctx, testDB, 0, 2, nil, nil)

	assert.Nil(t, err)
	assert.NotNil(t, orgs)
	assert.Equal(t, 2, len(orgs))
	assert.Greater(t, totalCount, 2)
}

func TestExecute_EmptyResult(t *testing.T) {
	testDB := testutil.SetupTestWithCleanup(t)
	ctx := context.Background()

	// Request page far beyond available data
	orgs, totalCount, err := Execute(ctx, testDB, 1000, 10, nil, nil)

	assert.Nil(t, err)
	assert.NotNil(t, orgs)
	assert.Equal(t, 0, len(orgs))
	assert.Greater(t, totalCount, 0)
}
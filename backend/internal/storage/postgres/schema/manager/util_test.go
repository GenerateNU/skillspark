package manager

import (
	"context"
	"skillspark/internal/storage/postgres/testutil"
	"testing"
)

func Test_CreateTestManager(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	ctx := context.Background()
	t.Parallel()
	CreateTestManager(t, ctx, testDB)
}

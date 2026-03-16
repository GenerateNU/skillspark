package saved

import (
	"context"
	"skillspark/internal/storage/postgres/testutil"
	"testing"
)

func Test_CreateTestChild(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	ctx := context.Background()
	t.Parallel()
	CreateTestSaved(t, ctx, testDB)
}

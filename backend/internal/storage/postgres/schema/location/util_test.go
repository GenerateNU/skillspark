package location

import (
	"context"
	"skillspark/internal/storage/postgres/testutil"
	"testing"
)

func Test_CreateTestLocation(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	ctx := context.Background()
	t.Parallel()
	CreateTestLocation(t, ctx, testDB)
}

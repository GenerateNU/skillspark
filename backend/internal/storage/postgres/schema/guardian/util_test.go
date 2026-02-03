package guardian

import (
	"context"
	"skillspark/internal/storage/postgres/testutil"
	"testing"
)

func Test_CreateTestGuardian(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	ctx := context.Background()
	t.Parallel()
	CreateTestGuardian(t, ctx, testDB)
}

package organization

import (
	"context"
	"skillspark/internal/storage/postgres/testutil"
	"testing"
)

func Test_CreateTestOrganization(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	ctx := context.Background()
	t.Parallel()
	CreateTestOrganization(t, ctx, testDB)
}

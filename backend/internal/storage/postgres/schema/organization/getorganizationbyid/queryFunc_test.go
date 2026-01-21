package getorganizationbyid
import (
	"context"
	"skillspark/internal/storage/postgres/testutil"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	ctx := context.Background()

	org, err := Execute(ctx, testDB, uuid.MustParse("00000000-0900-0000-0000-000000000004"))

	assert.Nil(t, err)
	assert.NotNil(t, org)
	assert.Equal(t, "Babel Street", org.Name)
	assert.Equal(t, true, org.Active)
}

func TestExecute_SecondOrganization(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	ctx := context.Background()

	org, err := Execute(ctx, testDB, uuid.MustParse("30000000-0030-0000-0000-000000000001"))

	assert.Nil(t, err)
	assert.NotNil(t, org)
	assert.Equal(t, "Tech Innovations", org.Name)
	assert.Equal(t, true, org.Active)
}

func TestExecute_NotFound(t *testing.T) {
	testDB := testutil.SetupTestDB(t)
	ctx := context.Background()

	org, err := Execute(ctx, testDB, uuid.New())

	assert.NotNil(t, err)
	assert.Nil(t, org)
}
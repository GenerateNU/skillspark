package emergencycontact

import (
	"context"
	"embed"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/guardian"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

//go:embed sql/*.sql
var SqlSavedFiles embed.FS

func CreateTestEmergencyContact(
	t *testing.T,
	ctx context.Context,
	db *pgxpool.Pool,
) *models.EmergencyContact {

	t.Helper()

	repo := NewEmergencyContactRepository(db)

	g := guardian.CreateTestGuardian(t, ctx, db)

	number := "+16462996961"

	input := &models.CreateEmergencyContactInput{}
	input.Body.Name = g.Name
	input.Body.GuardianID = g.ID
	input.Body.PhoneNumber = number

	c, err := repo.CreateEmergencyContact(ctx, input)

	require.NoError(t, err)
	require.NotNil(t, c)

	return c.Body
}

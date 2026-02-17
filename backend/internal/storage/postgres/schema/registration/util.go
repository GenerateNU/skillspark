package registration

import (
	"context"
	"embed"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/child"
	eventoccurrence "skillspark/internal/storage/postgres/schema/event-occurrence"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

//go:embed sql/*.sql
var SqlRegistrationFiles embed.FS

func CreateTestRegistration(
	t *testing.T,
	ctx context.Context,
	db *pgxpool.Pool,
) *models.Registration {
	t.Helper()

	repo := NewRegistrationRepository(db)

	child := child.CreateTestChild(t, ctx, db)
	occurrence := eventoccurrence.CreateTestEventOccurrence(t, ctx, db)

	input := func() *models.CreateRegistrationInput {
		i := &models.CreateRegistrationInput{}
		i.Body.ChildID = child.ID
		i.Body.GuardianID = child.GuardianID
		i.Body.EventOccurrenceID = occurrence.ID
		i.Body.Status = models.RegistrationStatusRegistered
		return i
	}()

	registration, err := repo.CreateRegistration(ctx, input)

	require.NoError(t, err)
	require.NotNil(t, registration.Body)

	return &registration.Body
}

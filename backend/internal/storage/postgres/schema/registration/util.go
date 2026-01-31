package registration

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/child"
	eventoccurrence "skillspark/internal/storage/postgres/schema/event-occurrence"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

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

	registration, _ := repo.CreateRegistration(ctx, input)

	return &registration.Body
}

package eventoccurrence

import (
	"context"
	"log"
	"skillspark/internal/errs"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *EventOccurrenceRepository) CancelEventOccurrence(ctx context.Context, id uuid.UUID) error {

	eo, err := r.GetEventOccurrenceByID(ctx, id)

	if err != nil {
		err := errs.InternalServerError("Failed to find event occurrence with given ID: ", err.Error())
		return &err
	}

	now := time.Now()
	if eo.StartTime.After(now) && eo.StartTime.Before(now.Add(24*time.Hour)) {
		err := errs.InternalServerError("Cannot delete event happening within the next 24 hours.")
		return &err
	}

	tx, err := r.db.Begin(ctx)
	if err != nil {
		e := errs.InternalServerError("Failed to begin transaction", err.Error())
		return &e
	}

	defer func() {
		if rerr := tx.Rollback(ctx); rerr != nil && rerr != pgx.ErrTxClosed {
			log.Printf("rollback failed: %v", rerr)
		}
	}()

	_, err = tx.Exec(ctx, `
		UPDATE registration
		SET status = 'cancelled'
		WHERE event_occurrence_id = $1
	`, id)
	if err != nil {
		e := errs.InternalServerError("Failed to cancel registrations", err.Error())
		return &e
	}

	commandTag, err := tx.Exec(ctx, `
		UPDATE event_occurrence
		SET status = 'cancelled',
			updated_at = NOW()
		WHERE id = $1
	`, id)
	if err != nil {
		e := errs.InternalServerError("Failed to cancel event occurrence", err.Error())
		return &e
	}

	if commandTag.RowsAffected() == 0 {
		e := errs.NotFound("Event", "id", id)
		return &e
	}

	if err := tx.Commit(ctx); err != nil {
		e := errs.InternalServerError("Failed to commit transaction", err.Error())
		return &e
	}

	return nil
}

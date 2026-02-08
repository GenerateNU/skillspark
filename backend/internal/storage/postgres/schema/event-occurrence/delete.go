package eventoccurrence

import (
	"context"
	"log"
	"skillspark/internal/errs"
	"skillspark/internal/storage/postgres/schema"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *EventOccurrenceRepository) CancelEventOccurrence(ctx context.Context, id uuid.UUID) error {
	eo, err := r.GetEventOccurrenceByID(ctx, id)
	if err != nil {
		e := errs.InternalServerError(
			"Failed to find event occurrence with given ID: ",
			err.Error(),
		)
		return &e
	}

	now := time.Now()
	if eo.StartTime.After(now) && eo.StartTime.Before(now.Add(24*time.Hour)) {
		e := errs.InternalServerError(
			"Cannot delete event happening within the next 24 hours.",
		)
		return &e
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

	// cancel associated registrations
	cancelRegistrationsQuery, err := schema.ReadSQLBaseScript(
		"event-occurrence/sql/cancel_registrations.sql",
	)
	if err != nil {
		e := errs.InternalServerError("Failed to read cancel registrations SQL", err.Error())
		return &e
	}

	if _, err := tx.Exec(ctx, cancelRegistrationsQuery, id); err != nil {
		e := errs.InternalServerError("Failed to cancel registrations", err.Error())
		return &e
	}

	// cancel the event occurrence
	cancelEventOccurrenceQuery, err := schema.ReadSQLBaseScript(
		"event-occurrence/sql/cancel_eventoccurrence.sql",
	)
	if err != nil {
		e := errs.InternalServerError("Failed to read cancel event occurrence SQL", err.Error())
		return &e
	}

	commandTag, err := tx.Exec(ctx, cancelEventOccurrenceQuery, id)
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

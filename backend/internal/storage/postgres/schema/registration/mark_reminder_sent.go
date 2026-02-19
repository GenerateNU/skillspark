package registration

import (
	"context"
	_ "embed"

	"skillspark/internal/errs"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

//go:embed sql/mark_reminder_sent.sql
var markReminderSentSQL string

// MarkReminderSent updates the reminder_sent status of a registration.
func (r *RegistrationRepository) MarkReminderSent(ctx context.Context, tx pgx.Tx, id uuid.UUID, sent bool) error {
	var row pgx.Row
	if tx != nil {
		row = tx.QueryRow(ctx, markReminderSentSQL, id, sent)
	} else {
		row = r.db.QueryRow(ctx, markReminderSentSQL, id, sent)
	}

	var returnedID uuid.UUID
	err := row.Scan(&returnedID)
	if err != nil {
		if err == pgx.ErrNoRows {
			err := errs.NotFound("Registration", "id", id)
			return &err
		}
		err := errs.InternalServerError("Failed to mark reminder sent: ", err.Error())
		return &err
	}

	return nil
}

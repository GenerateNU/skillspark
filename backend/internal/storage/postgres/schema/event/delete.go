package event

import (
	"context"
	"github.com/google/uuid"
	"skillspark/internal/errs"
	"skillspark/internal/storage/postgres/schema"
)

func (r *EventRepository) DeleteEvent(ctx context.Context, id uuid.UUID) (*struct{}, error) {
	query, err := schema.ReadSQLBaseScript("event/sql/delete.sql")
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	commandTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		err := errs.InternalServerError("Failed to delete event: ", err.Error())
		return nil, &err
	}

	if commandTag.RowsAffected() == 0 {
		err := errs.NotFound("Event", "id", id)
		return nil, &err
	}

	return nil, nil
}

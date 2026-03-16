package saved

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
)

func (r *SavedRepository) DeleteSaved(ctx context.Context, id uuid.UUID) error {

	query, err := schema.ReadSQLBaseScript("delete.sql", SqlSavedFiles)

	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return &err
	}

	commandTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		err := errs.InternalServerError("Failed to delete saved: ", err.Error())
		return &err
	}

	if commandTag.RowsAffected() == 0 {
		err := errs.NotFound("Saved", "id", id)
		return &err
	}

	return nil
}

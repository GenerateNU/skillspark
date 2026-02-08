package review

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
)

func (r *ReviewRepository) DeleteReview(ctx context.Context, id uuid.UUID) error {

	query, err := schema.ReadSQLBaseScript("review/sql/delete.sql")
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return &err
	}

	commandTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		err := errs.InternalServerError("Failed to delete review: ", err.Error())
		return &err
	}

	if commandTag.RowsAffected() == 0 {
		err := errs.NotFound("Review", "id", id)
		return &err
	}

	return nil
}

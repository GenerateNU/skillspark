package review

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
	"skillspark/internal/utils"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *ReviewRepository) GetReviewsByEventID(ctx context.Context, id uuid.UUID, AcceptLanguage string, pagination utils.Pagination) ([]models.Review, error) {

	language = AcceptLanguage
	baseQuery, err := schema.ReadSQLBaseScript("review/sql/get_by_event_id.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	rows, err := r.db.Query(ctx, baseQuery, id, pagination.Limit, pagination.GetOffset())
	if err != nil {
		errr := errs.InternalServerError("Failed to get reviews: ", err.Error())
		return nil, &errr
	}
	defer rows.Close()

	reviews, err := pgx.CollectRows(rows, ScanReviews)
	if err != nil {
		errr := errs.InternalServerError("Failed to collect reviews: ", err.Error())
		return nil, &errr
	}

	return reviews, nil
}

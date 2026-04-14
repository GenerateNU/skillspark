package review

import (
	"context"
	"fmt"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
	"skillspark/internal/utils"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *ReviewRepository) GetReviewsByEventID(ctx context.Context, id uuid.UUID, AcceptLanguage string, pagination utils.Pagination, sortBy string) ([]models.Review, error) {

	language = AcceptLanguage

	baseQuery, err := schema.ReadSQLBaseScript("get_by_event_id.sql", SqlReviewFiles)

	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	var orderByClause string
	switch sortBy {
	case "highest":
		orderByClause = "r.rating DESC, r.created_at DESC"
	case "lowest":
		orderByClause = "r.rating ASC, r.created_at DESC"
	default:
		orderByClause = "r.created_at DESC"
	}

	query := fmt.Sprintf(baseQuery, orderByClause)

	rows, err := r.db.Query(ctx, query, id, pagination.Limit, pagination.GetOffset())

	if err != nil {
		errr := errs.InternalServerError("Failed to get reviews: ", err.Error())
		return nil, &errr
	}
	defer rows.Close()

	reviews, err := pgx.CollectRows(rows, r.ScanReviews)
	if err != nil {
		errr := errs.InternalServerError("Failed to collect reviews: ", err.Error())
		return nil, &errr
	}

	return reviews, nil
}

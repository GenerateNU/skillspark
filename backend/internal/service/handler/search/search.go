package search

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/utils"
)

func (h *Handler) SearchEvents(ctx context.Context, query string, acceptLanguage string, pagination utils.Pagination) ([]models.Event, error) {
	from := (pagination.Page - 1) * pagination.Limit
	return h.OpenSearchClient.FuzzySearch(ctx, query, acceptLanguage, from, pagination.Limit)
}

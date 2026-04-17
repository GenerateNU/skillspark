package search

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/utils"
	"time"
)

func (h *Handler) SearchEvents(ctx context.Context, query string, acceptLanguage string, pagination utils.Pagination) ([]models.Event, error) {
	if h.OpenSearchClient == nil {
		return h.EventRepo.GetAllEvents(ctx, pagination, acceptLanguage, models.GetAllEventsFilter{Search: &query})
	}
	from := (pagination.Page - 1) * pagination.Limit
	events, err := h.OpenSearchClient.FuzzySearch(ctx, query, acceptLanguage, from, pagination.Limit)
	if err != nil {
		return nil, err
	}

	for i := range events {
		if events[i].HeaderImageS3Key != nil {
			url, err := h.S3Client.GeneratePresignedURL(ctx, *events[i].HeaderImageS3Key, time.Hour)
			if err == nil {
				events[i].PresignedURL = &url
			}
		}
	}

	return events, nil
}

package event

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/utils"
	"time"
)

func (h *Handler) GetAllEvents(ctx context.Context, pagination utils.Pagination, acceptLanguage string) ([]models.Event, error) {
	events, err := h.EventRepository.GetAllEvents(ctx, pagination, acceptLanguage)
	if err != nil {
		return nil, err
	}

	for idx := range events {
		if events[idx].HeaderImageS3Key != nil {
			url, err := h.s3client.GeneratePresignedURL(ctx, *events[idx].HeaderImageS3Key, time.Hour)
			if err != nil {
				return nil, err
			}
			events[idx].PresignedURL = &url
		}
	}

	return events, nil
}

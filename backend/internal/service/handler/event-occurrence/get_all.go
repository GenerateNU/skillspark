package eventoccurrence

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/utils"
	"time"
)

func (h *Handler) GetAllEventOccurrences(ctx context.Context, pagination utils.Pagination, AcceptLanguage string, filters models.GetAllEventOccurrencesFilter) ([]models.EventOccurrence, error) {

	if AcceptLanguage != "en-US" && AcceptLanguage != "th-TH" {
		e := errs.BadRequest("Invalid AcceptLanguage parameter: language does not exist")
		return nil, &e
	}

	eventOccurrence, err := h.EventOccurrenceRepository.GetAllEventOccurrences(ctx, pagination, AcceptLanguage, filters)
	if err != nil {
		return nil, err
	}

	for idx := range eventOccurrence {
		err = h.AssignURLS(ctx, eventOccurrence, idx)
		if err != nil {
			return nil, err
		}
	}

	return eventOccurrence, nil
}

func (h *Handler) AssignURLS(ctx context.Context, occurrences []models.EventOccurrence, idx int) error {
	var url string
	var err error
	key := occurrences[idx].Event.HeaderImageS3Key
	if key != nil {
		url, err = h.s3Client.GeneratePresignedURL(ctx, *key, time.Hour)
		if err != nil {
			return err
		}
	}

	occurrences[idx].Event.PresignedURL = &url
	return nil

}

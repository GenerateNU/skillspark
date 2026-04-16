package saved

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/utils"
	"time"

	"github.com/google/uuid"
)

func (h *Handler) GetByGuardianID(
	ctx context.Context,
	id uuid.UUID,
	pagination utils.Pagination,
	AcceptLanguage string,
) ([]models.Saved, error) {

	_, err := h.GuardianRepository.GetGuardianByID(ctx, id)
	if err != nil {
		return nil, errs.BadRequest("Invalid guardian_id: guardian does not exist")
	}

	saved, httpErr := h.SavedRepository.GetByGuardianID(ctx, id, pagination, AcceptLanguage)
	if httpErr != nil {
		return nil, httpErr
	}

	for idx := range saved {

		key := saved[idx].Event.HeaderImageS3Key
		if key != nil {

			presignedURL, err := h.s3Client.GeneratePresignedURL(ctx, *key, time.Hour)
			if err != nil {
				return nil, err
			}

			saved[idx].Event.PresignedURL = &presignedURL
		}
	}

	return saved, nil
}

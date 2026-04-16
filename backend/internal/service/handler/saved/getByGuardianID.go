package saved

import (
	"context"
	"fmt"
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

	fmt.Println("➡️ ENTER GetByGuardianID")
	fmt.Println("guardianID:", id)

	// 1. Guardian lookup (VERY likely crash point)
	fmt.Println("➡️ calling GuardianRepository.GetGuardianByID")

	if h.GuardianRepository == nil {
		fmt.Println("💥 GuardianRepository is NIL")
	}

	_, err := h.GuardianRepository.GetGuardianByID(ctx, id)
	if err != nil {
		fmt.Println("💥 Guardian lookup failed:", err)
		return nil, errs.BadRequest("Invalid guardian_id: guardian does not exist")
	}

	fmt.Println("➡️ guardian exists")

	// 2. Saved query
	fmt.Println("➡️ calling SavedRepository.GetByGuardianID")

	if h.SavedRepository == nil {
		fmt.Println("💥 SavedRepository is NIL")
	}

	saved, httpErr := h.SavedRepository.GetByGuardianID(ctx, id, pagination, AcceptLanguage)
	if httpErr != nil {
		fmt.Println("💥 saved repo error:", httpErr)
		return nil, httpErr
	}

	fmt.Println("➡️ saved rows returned:", len(saved))

	// 3. S3 loop (second likely crash point)
	for idx := range saved {
		fmt.Println("➡️ processing index:", idx)

		if h.s3Client == nil {
			fmt.Println("💥 s3Client is NIL")
		}

		key := saved[idx].Event.HeaderImageS3Key
		if key != nil {
			fmt.Println("➡️ generating presigned URL for:", *key)

			presignedURL, err := h.s3Client.GeneratePresignedURL(ctx, *key, time.Hour)
			if err != nil {
				fmt.Println("💥 S3 error:", err)
				return nil, err
			}

			saved[idx].Event.PresignedURL = &presignedURL
		}
	}

	fmt.Println("✅ EXIT GetByGuardianID OK")
	return saved, nil
}

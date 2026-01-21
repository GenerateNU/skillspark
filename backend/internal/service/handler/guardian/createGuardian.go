package guardian

import (
	"context"
	"errors"
	"net/http"
	"skillspark/internal/errs"
	"skillspark/internal/models"
)

func (h *Handler) CreateGuardian(ctx context.Context, input *models.CreateGuardianInput) (*models.Guardian, error) {

	guardian, err := h.GuardianRepository.GetGuardianByUserID(ctx, input.Body.UserID)

	if err != nil {
		var httpErr *errs.HTTPError
		if errors.As(err, &httpErr) && httpErr.Code == http.StatusNotFound {
			// proceed
		} else {
			return nil, err
		}
	}

	if guardian != nil {
		return nil, errs.BadRequest("Guardian already exists for this user")
	}

	guardian, err = h.GuardianRepository.CreateGuardian(ctx, input)
	if err != nil {
		return nil, err
	}

	return guardian, nil
}

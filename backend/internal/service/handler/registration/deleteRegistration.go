package registration

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

func (h *Handler) DeleteRegistration(ctx context.Context, input *models.DeleteRegistrationInput) (*models.DeleteRegistrationOutput, error) {
	id, err := uuid.Parse(input.ID.String())
	if err != nil {
		return nil, errs.BadRequest("Invalid ID format")
	}

	deleted, httpErr := h.RegistrationRepository.DeleteRegistration(ctx, &models.DeleteRegistrationInput{ID: id})
	if httpErr != nil {
		return nil, httpErr
	}

	return deleted, nil
}
package review

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
)

func (h *Handler) CreateReview(ctx context.Context, input *models.CreateReviewInput) (*models.CreateReviewOutput, *errs.HTTPError) {

	description, err := h.CallTranslateAPI(ctx, &input.Body.Description)
	if err != nil {
		e := errs.BadRequest("Invalid registration_id: registration does not exist")
		return nil, &e
	}
	CreateReviewInput := h.CreateTranslateStruct(ctx, input, description)

	if _, err := h.RegistrationRepository.GetRegistrationByID(ctx, &models.GetRegistrationByIDInput{
		ID: input.Body.RegistrationID,
	}, nil); err != nil {
		e := errs.BadRequest("Invalid registration_id: registration does not exist" + err.Error())
		return nil, &e
	}

	if _, err := h.GuardianRepository.GetGuardianByID(ctx, input.Body.GuardianID); err != nil {
		e := errs.BadRequest("Invalid guardian_id: guardian does not exist")
		return nil, &e
	}

	review, err := h.ReviewRepository.CreateReview(ctx, CreateReviewInput)
	if err != nil {
		if httpErr, ok := err.(*errs.HTTPError); ok {
			return nil, httpErr
		}
		e := errs.InternalServerError("Failed to create review: " + err.Error())
		return nil, &e
	}

	return &models.CreateReviewOutput{
		Body: *review,
	}, nil
}

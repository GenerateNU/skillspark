package review

import (
	"context"
	"fmt"
	"skillspark/internal/errs"
	"skillspark/internal/models"
)

func (h *Handler) CreateReview(ctx context.Context, input *models.CreateReviewInput) (*models.CreateReviewOutput, *errs.HTTPError) {

	fmt.Println("Input")
	fmt.Printf("%+v\n", input)

	translateInput := []*string{&input.Body.Description}

	description, err := h.TranslateClient.CallTranslateAPI(ctx, translateInput, input.AcceptLanguage)
	if err != nil {
		fmt.Println("Translation error:", err)
		e := errs.BadRequest("Invalid registration_id: registration does not exist")
		return nil, &e
	}
	CreateReviewInput := h.CreateTranslateStruct(ctx, input, description[input.Body.Description])

	fmt.Println("CreateReviewInput")
	fmt.Printf("%+v\n", CreateReviewInput)

	if _, err := h.RegistrationRepository.GetRegistrationByID(ctx, &models.GetRegistrationByIDInput{
		ID: input.Body.RegistrationID,
	}, nil); err != nil {
		fmt.Println("Registration error:", err)
		e := errs.BadRequest("Invalid registration_id: registration does not exist" + err.Error())
		return nil, &e
	}

	if input.Body.GuardianID != nil {
		if _, err := h.GuardianRepository.GetGuardianByID(ctx, *input.Body.GuardianID); err != nil {
			fmt.Println("Guardian error:", err)
			e := errs.BadRequest("Invalid guardian_id: guardian does not exist")
			return nil, &e
		}
	}

	review, err := h.ReviewRepository.CreateReview(ctx, CreateReviewInput)
	if err != nil {
		fmt.Println("Review error:", err)

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

package event

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/s3_client"
)

func (h *Handler) CreateEvent(ctx context.Context, input *models.CreateEventInput, updateBody *models.UpdateEventBody, imageData *[]byte, s3Client s3_client.S3Interface) (*models.Event, error) {
	var key *string
	var url *string

	initInput := h.CreateTranslateStruct(ctx, input, nil, nil)
	event, err := h.EventRepository.CreateEvent(ctx, initInput, key)
	if err != nil {
		e := errs.InternalServerError("Invalid creation" + err.Error())
		return nil, e
	}

	updateInput := &models.UpdateEventInput{
		ID:   event.ID,
		Body: *updateBody,
	}

	translationsReinsertion, err := h.TranslationHelper(ctx, event, updateInput)
	if err != nil {
		e := errs.BadRequest("Invalid translation call" + err.Error())
		return nil, e
	}

	url, key, err = h.CreateEventS3Helper(ctx, s3Client, event, updateInput, imageData)
	if err != nil {
		e := errs.InternalServerError("Something went wrong when inserting into S3" + err.Error())
		return nil, e
	}
	_, err = h.EventRepository.UpdateEvent(ctx, translationsReinsertion, key)

	if err != nil {
		e := errs.InternalServerError("Invalid Update" + err.Error())
		return nil, e
	}
	event.PresignedURL = url
	event.HeaderImageS3Key = key
	return event, nil
}

// helper for uploading image to s3
func (h *Handler) CreateEventS3Helper(ctx context.Context, s3Client s3_client.S3Interface, event *models.Event,
	updateInput *models.UpdateEventInput, imageData *[]byte) (*string, *string, error) {

	if imageData != nil {

		key, err := h.generateS3Key(event.ID)
		if err != nil {
			return nil, nil, err
		}

		url, errr := s3Client.UploadImage(ctx, key, *imageData)
		if errr != nil {
			return nil, nil, errr
		}

		return url, key, nil

	}

	return nil, nil, nil

}

func (h *Handler) TranslationHelper(ctx context.Context, event *models.Event, updateInput *models.UpdateEventInput) (*models.UpdateEventDBInput, error) {
	translationResp, err := h.CallTranslateAPI(ctx, &event.Title, &event.Description)
	if err != nil {
		return nil, err
	}
	translationsReinsertion := h.UpdateTranslateStruct(ctx, updateInput, translationResp.Title_TH, translationResp.Description_TH)
	_, err = h.EventRepository.UpdateEvent(ctx, translationsReinsertion, nil)
	if err != nil {
		return nil, err
	}

	return translationsReinsertion, nil
}

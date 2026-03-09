package event

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/s3_client"
)

func (h *Handler) UpdateEvent(ctx context.Context, input *models.UpdateEventInput, image_data *[]byte, s3Client s3_client.S3Interface) (*models.Event, error) {

	var key *string
	var url *string

	if input.AcceptLanguage != "en-US" && input.AcceptLanguage != "th-TH" {
		e := errs.BadRequest("Invalid AcceptLanguage parameter: language does not exist")
		return nil, &e
	}

	translateInput := []*string{input.Body.Title, input.Body.Description}

	translationResp, err := h.TranslateClient.CallTranslateAPI(ctx, translateInput, input.AcceptLanguage)
	if err != nil {
		e := errs.InternalServerError("Translation failed", err.Error())
		return nil, e
	}

	translatedTitle := translationResp[*input.Body.Title]
	translatedDescription := translationResp[*input.Body.Description]

	updateInput := h.UpdateTranslateStruct(ctx, input, translatedTitle, translatedDescription)

	if image_data != nil {
		var err error
		url, key, err = h.UpdateEventS3Helper(ctx, s3Client, input, image_data)
		if err != nil {
			e := errs.InternalServerError("S3 upload failed", err.Error())
			return nil, e
		}
	}

	event, err := h.EventRepository.UpdateEvent(ctx, updateInput, key)
	if err != nil {
		return nil, err
	}

	event.PresignedURL = url

	return event, nil
}

func (h *Handler) UpdateEventS3Helper(ctx context.Context, s3Client s3_client.S3Interface, input *models.UpdateEventInput, image_data *[]byte) (*string, *string, error) {
	key, genErr := h.generateS3Key(input.ID)
	if genErr != nil {
		return nil, nil, genErr
	}
	url, errr := s3Client.UploadImage(ctx, key, *image_data)
	if errr != nil {
		return nil, nil, errr
	}

	return url, key, nil
}

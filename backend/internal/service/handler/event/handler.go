package event

import (
	"skillspark/internal/s3_client"
	"skillspark/internal/storage"
	translations "skillspark/internal/translation"
)

type Handler struct {
	EventRepository storage.EventRepository
	s3client        s3_client.S3Interface
	TranslateClient translations.TranslationInterface
}

func NewHandler(eventRepository storage.EventRepository, s3client s3_client.S3Interface, translateClient translations.TranslationInterface) *Handler {
	return &Handler{
		EventRepository: eventRepository,
		s3client:        s3client,
		TranslateClient: translateClient,
	}
}

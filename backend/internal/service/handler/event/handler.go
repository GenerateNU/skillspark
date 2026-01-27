package event

import (
	"skillspark/internal/s3_client"
	"skillspark/internal/storage"
)

type Handler struct {
	EventRepository storage.EventRepository
	s3client        s3_client.Client
}

func NewHandler(eventRepository storage.EventRepository, s3client *s3_client.Client) *Handler {
	return &Handler{
		EventRepository: eventRepository,
		s3client:        *s3client,
	}
}

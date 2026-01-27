package event

import (
	"skillspark/internal/errs"
	"skillspark/internal/s3_client"
	"skillspark/internal/storage"

	"github.com/google/uuid"
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

func (h *Handler) generateS3Key(id uuid.UUID) (string, error) {
	if id == uuid.Nil {
		err := errs.InternalServerError("Failed to create location: invalid UUID")
		return "", &err
	}

	id_string := id.String()
	res := "events" + "/" + id_string
	return res, nil

}

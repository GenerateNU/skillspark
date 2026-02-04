package event

import (
	"skillspark/internal/errs"

	"github.com/google/uuid"
)

func (h *Handler) generateS3Key(id uuid.UUID) (*string, error) {
	if id == uuid.Nil {
		err := errs.InternalServerError("Failed to create location: invalid UUID")
		return nil, &err
	}

	res := "events/header-image/" + id.String()
	return &res, nil

}

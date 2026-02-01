package organization

import (
	"skillspark/internal/errs"

	"github.com/google/uuid"
)

func (h *Handler) generateS3Key(id uuid.UUID) (*string, error) {
	if id == uuid.Nil {
		err := errs.InternalServerError("Failed to generate S3 key: invalid UUID")
		return nil, &err
	}

	res := "orgs/header-image/" + id.String()
	return &res, nil
}

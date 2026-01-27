package organization

import (
	"skillspark/internal/errs"
	"skillspark/internal/s3_client"
	"skillspark/internal/storage"

	"github.com/google/uuid"
)

type Handler struct {
	OrganizationRepository storage.OrganizationRepository
	LocationRepository     storage.LocationRepository
	s3client               s3_client.Client
}

func NewHandler(orgRepo storage.OrganizationRepository, locRepo storage.LocationRepository, s3client *s3_client.Client) *Handler {
	return &Handler{
		OrganizationRepository: orgRepo,
		LocationRepository:     locRepo,
		s3client:               *s3client,
	}
}

func (h *Handler) generateS3Key(id uuid.UUID) (string, error) {
	if id == uuid.Nil {
		err := errs.InternalServerError("Failed to create location: invalid UUID")
		return "", &err
	}

	id_string := id.String()
	res := "orgs" + "/" + id_string
	return res, nil

}

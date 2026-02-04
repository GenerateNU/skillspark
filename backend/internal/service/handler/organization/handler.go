package organization

import (
	"skillspark/internal/s3_client"
	"skillspark/internal/storage"
)

type Handler struct {
	OrganizationRepository storage.OrganizationRepository
	LocationRepository     storage.LocationRepository
	s3client               s3_client.S3Interface
}

func NewHandler(orgRepo storage.OrganizationRepository, locRepo storage.LocationRepository, s3client s3_client.S3Interface) *Handler {
	return &Handler{
		OrganizationRepository: orgRepo,
		LocationRepository:     locRepo,
		s3client:               s3client,
	}
}

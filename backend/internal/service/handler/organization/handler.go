package organization

import (
	"skillspark/internal/s3_client"
	"skillspark/internal/storage"
	translations "skillspark/internal/translation"
)

type Handler struct {
	OrganizationRepository storage.OrganizationRepository
	LocationRepository     storage.LocationRepository
	ReviewRepository       storage.ReviewRepository
	s3client               s3_client.S3Interface
	TranslateClient        translations.TranslationInterface
}

func NewHandler(orgRepo storage.OrganizationRepository, locRepo storage.LocationRepository, reviewRepo storage.ReviewRepository, s3client s3_client.S3Interface, translateClient translations.TranslationInterface) *Handler {
	return &Handler{
		OrganizationRepository: orgRepo,
		LocationRepository:     locRepo,
		ReviewRepository:       reviewRepo,
		s3client:               s3client,
		TranslateClient:        translateClient,
	}
}

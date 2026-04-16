package saved

import (
	"skillspark/internal/s3_client"
	"skillspark/internal/storage"
)

type Handler struct {
	SavedRepository    storage.SavedRepository
	GuardianRepository storage.GuardianRepository
	s3Client           s3_client.S3Interface
}

func NewHandler(savedRepository storage.SavedRepository, guardianRepository storage.GuardianRepository, s3client s3_client.S3Interface) *Handler {
	return &Handler{
		SavedRepository:    savedRepository,
		GuardianRepository: guardianRepository,
		s3Client:           s3client,
	}
}

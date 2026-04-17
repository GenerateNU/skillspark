package search

import (
	"skillspark/internal/opensearch"
	"skillspark/internal/s3_client"
	"skillspark/internal/storage"
)

type Handler struct {
	OpenSearchClient *opensearch.Client
	S3Client         s3_client.S3Interface
	EventRepo        storage.EventRepository
}

func NewHandler(osClient *opensearch.Client, s3 s3_client.S3Interface, eventRepo storage.EventRepository) *Handler {
	return &Handler{
		OpenSearchClient: osClient,
		S3Client:         s3,
		EventRepo:        eventRepo,
	}
}

package search

import (
	"skillspark/internal/opensearch"
	"skillspark/internal/s3_client"
)

type Handler struct {
	OpenSearchClient *opensearch.Client
	S3Client         s3_client.S3Interface
}

func NewHandler(osClient *opensearch.Client, s3 s3_client.S3Interface) *Handler {
	return &Handler{
		OpenSearchClient: osClient,
		S3Client:         s3,
	}
}

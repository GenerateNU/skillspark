package search

import (
	"skillspark/internal/opensearch"
)

type Handler struct {
	OpenSearchClient *opensearch.Client
}

func NewHandler(osClient *opensearch.Client) *Handler {
	return &Handler{
		OpenSearchClient: osClient,
	}
}

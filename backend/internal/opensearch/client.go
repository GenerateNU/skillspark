package opensearch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"skillspark/internal/config"
	"skillspark/internal/models"

	"github.com/opensearch-project/opensearch-go/v4"
	"github.com/opensearch-project/opensearch-go/v4/opensearchapi"
)

const index = "events"

type Client struct {
	api *opensearchapi.Client
}

func NewClient(cfg config.OpenSearch) (*Client, error) {
	client, err := opensearchapi.NewClient(opensearchapi.Config{
		Client: opensearch.Config{
			Addresses: []string{cfg.URL},
			Username:  cfg.Username,
			Password:  cfg.Password,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("opensearch: failed to create client: %w", err)
	}
	return &Client{api: client}, nil
}

func (c *Client) FuzzySearch(ctx context.Context, query string, acceptLanguage string, from, size int) ([]models.Event, error) {
	titleField := "title_en"
	descField := "description_en"
	if acceptLanguage == "th-TH" {
		titleField = "title_th"
		descField = "description_th"
	}

	body := map[string]any{
		"from": from,
		"size": size,
		"query": map[string]any{
			"multi_match": map[string]any{
				"query":     query,
				"fields":    []string{titleField + "^2", descField, "category"},
				"fuzziness": "AUTO",
			},
		},
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("opensearch: failed to marshal query: %w", err)
	}

	resp, err := c.api.Search(ctx, &opensearchapi.SearchReq{
		Indices: []string{index},
		Body:    bytes.NewReader(bodyBytes),
	})
	if err != nil {
		return nil, fmt.Errorf("opensearch: search failed: %w", err)
	}

	var events []models.Event
	for _, hit := range resp.Hits.Hits {
		var event models.Event
		if err := json.Unmarshal(hit.Source, &event); err != nil {
			return nil, fmt.Errorf("opensearch: failed to unmarshal hit: %w", err)
		}
		events = append(events, event)
	}
	return events, nil
}

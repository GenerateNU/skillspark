package opensearch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"skillspark/internal/config"
	"skillspark/internal/models"

	"github.com/google/uuid"
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
			"bool": map[string]any{
				"should": []any{
					map[string]any{
						"multi_match": map[string]any{
							"query":     query,
							"fields":    []string{titleField + "^2", descField},
							"fuzziness": "AUTO",
						},
					},
					map[string]any{
						"term": map[string]any{
							"category": query,
						},
					},
				},
				"minimum_should_match": 1,
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

	type osEvent struct {
		ID               string   `json:"id"`
		TitleEN          string   `json:"title_en"`
		TitleTH          string   `json:"title_th"`
		DescriptionEN    string   `json:"description_en"`
		DescriptionTH    string   `json:"description_th"`
		Category         []string `json:"category"`
		HeaderImageS3Key *string  `json:"header_image_s3_key"`
		AgeRangeMin      *int     `json:"age_range_min"`
		AgeRangeMax      *int     `json:"age_range_max"`
	}

	events := make([]models.Event, 0)
	for _, hit := range resp.Hits.Hits {
		var src osEvent
		if err := json.Unmarshal(hit.Source, &src); err != nil {
			return nil, fmt.Errorf("opensearch: failed to unmarshal hit: %w", err)
		}
		id, err := uuid.Parse(src.ID)
		if err != nil {
			continue
		}
		title, description := src.TitleEN, src.DescriptionEN
		if acceptLanguage == "th-TH" && src.TitleTH != "" {
			title, description = src.TitleTH, src.DescriptionTH
		}
		events = append(events, models.Event{
			ID:               id,
			Title:            title,
			Description:      description,
			Category:         src.Category,
			HeaderImageS3Key: src.HeaderImageS3Key,
			AgeRangeMin:      src.AgeRangeMin,
			AgeRangeMax:      src.AgeRangeMax,
		})
	}
	return events, nil
}

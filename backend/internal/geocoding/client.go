package geocoding

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const defaultMinConfidence = 7

// ErrInvalidAddress is returned when OpenCage cannot geocode an address with sufficient confidence.
var ErrInvalidAddress = errors.New("address is invalid or could not be geocoded with sufficient confidence")

type opencageResponse struct {
	Results []struct {
		Geometry struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"geometry"`
		Confidence int `json:"confidence"`
	} `json:"results"`
}

type Client struct {
	apiKey        string
	minConfidence int
	httpClient    *http.Client
}

// NewClient creates an OpenCage client reading OPENCAGE_API_KEY and
// OPENCAGE_MIN_CONFIDENCE (default 7) from the environment.
func NewClient() (*Client, error) {
	apiKey := os.Getenv("OPENCAGE_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("OPENCAGE_API_KEY is not set")
	}

	minConfidence := defaultMinConfidence
	if v := os.Getenv("OPENCAGE_MIN_CONFIDENCE"); v != "" {
		parsed, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("invalid OPENCAGE_MIN_CONFIDENCE value %q: %w", v, err)
		}
		minConfidence = parsed
	}

	return &Client{
		apiKey:        apiKey,
		minConfidence: minConfidence,
		httpClient:    &http.Client{},
	}, nil
}

// Geocode calls the OpenCage API and returns lat/lng for the given address
// Returns ErrInvalidAddress when there are no results or confidence is too low
func (c *Client) Geocode(ctx context.Context, address string) (lat, lng *float64, err error) {
	endpoint := fmt.Sprintf(
		"https://api.opencagedata.com/geocode/v1/json?q=%s&key=%s&limit=1",
		url.QueryEscape(address), c.apiKey,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build opencage request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("opencage request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return  nil, nil, fmt.Errorf("opencage returned unexpected status %d", resp.StatusCode)
	}

	var result opencageResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, nil, fmt.Errorf("failed to decode opencage response: %w", err)
	}

	if len(result.Results) == 0 || result.Results[0].Confidence <= c.minConfidence {
		return nil, nil, ErrInvalidAddress
	}

	r := result.Results[0]
	return &r.Geometry.Lat, &r.Geometry.Lng, nil
}

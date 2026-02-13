package auth

import (
	"net/http"
	authLib "skillspark/internal/auth"
	"testing"

	"bytes"
	"encoding/json"
	"io"
	"time"
)

// MockRoundTripper is a helper to mock http.Client via its Transport
type MockRoundTripper struct {
	RoundTripFunc func(req *http.Request) *http.Response
}

func (m *MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.RoundTripFunc(req), nil
}

func SetupMockAuthClient(t *testing.T, responseBody interface{}, statusCode int) {
	originalClient := authLib.Client
	t.Cleanup(func() {
		authLib.Client = originalClient
	})

	mockTransport := &MockRoundTripper{
		RoundTripFunc: func(req *http.Request) *http.Response {
			respBytes, _ := json.Marshal(responseBody)
			return &http.Response{
				StatusCode: statusCode,
				Body:       io.NopCloser(bytes.NewBuffer(respBytes)),
				Header:     make(http.Header),
			}
		},
	}

	authLib.Client = &http.Client{
		Transport: mockTransport,
		Timeout:   1 * time.Second,
	}
}
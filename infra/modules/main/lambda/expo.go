package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

const (
	expoAPIURL = "https://exp.host/--/api/v2/push/send"
)

// ExpoClient handles push notification sending via Expo API
type ExpoClient struct {
	accessToken string
	client      *http.Client
}

// NewExpoClient creates a new Expo client
func NewExpoClient() *ExpoClient {
	accessToken := os.Getenv("EXPO_ACCESS_TOKEN")
	
	return &ExpoClient{
		accessToken: accessToken,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// ExpoPushMessage represents a single push notification message
type ExpoPushMessage struct {
	To       string          `json:"to"`
	Title    string          `json:"title,omitempty"`
	Body     string          `json:"body"`
	Data     json.RawMessage `json:"data,omitempty"`
	Sound    string          `json:"sound,omitempty"`
	Priority string          `json:"priority,omitempty"`
}

// ExpoPushRequest represents the request payload for Expo API
type ExpoPushRequest struct {
	To    []string         `json:"to"`
	Title string           `json:"title,omitempty"`
	Body  string           `json:"body"`
	Data  json.RawMessage  `json:"data,omitempty"`
	Sound string           `json:"sound,omitempty"`
}

// ExpoPushResponse represents the response from Expo API
type ExpoPushResponse struct {
	Data []ExpoPushTicket `json:"data"`
}

// ExpoPushTicket represents a ticket for a push notification
type ExpoPushTicket struct {
	Status string `json:"status"`
	ID     string `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// SendPushNotification sends a push notification via Expo API
func (c *ExpoClient) SendPushNotification(ctx context.Context, token string, body string, metadata json.RawMessage) error {
	if token == "" {
		return fmt.Errorf("push token is required")
	}

	// Create request payload
	reqBody := ExpoPushRequest{
		To:   []string{token},
		Body: body,
		Data: metadata,
		Sound: "default",
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create HTTP request
	req, err := http.NewRequestWithContext(ctx, "POST", expoAPIURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if c.accessToken != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
	}

	// Send request
	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var errorBody bytes.Buffer
		errorBody.ReadFrom(resp.Body)
		
		slog.Error("Expo API error",
			"status", resp.StatusCode,
			"response", errorBody.String(),
		)
		
		return fmt.Errorf("expo API returned status %d: %s", resp.StatusCode, errorBody.String())
	}

	// Parse response
	var pushResp ExpoPushResponse
	if err := json.NewDecoder(resp.Body).Decode(&pushResp); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	// Check ticket status
	if len(pushResp.Data) > 0 {
		ticket := pushResp.Data[0]
		if ticket.Status != "ok" {
			return fmt.Errorf("expo push notification failed: %s", ticket.Message)
		}
		
		slog.Info("Push notification sent successfully",
			"token", token,
			"ticket_id", ticket.ID,
		)
	}

	return nil
}


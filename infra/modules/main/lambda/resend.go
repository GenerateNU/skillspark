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
	resendAPIURL = "https://api.resend.com/emails"
)

// ResendClient handles email sending via Resend API
type ResendClient struct {
	apiKey  string
	client  *http.Client
	from    string
}

// NewResendClient creates a new Resend client
func NewResendClient() (*ResendClient, error) {
	apiKey := os.Getenv("RESEND_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("RESEND_API_KEY environment variable is required")
	}

	// Default from address - can be overridden via environment variable
	from := os.Getenv("RESEND_FROM_EMAIL")
	if from == "" {
		from = "notifications@skillspark.app" // Default, should be configured
	}

	return &ResendClient{
		apiKey: apiKey,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		from: from,
	}, nil
}

// ResendEmailRequest represents the request payload for Resend API
type ResendEmailRequest struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Text    string   `json:"text"`
	HTML    string   `json:"html,omitempty"`
}

// ResendEmailResponse represents the response from Resend API
type ResendEmailResponse struct {
	ID string `json:"id"`
}

// SendEmail sends an email via Resend API
func (c *ResendClient) SendEmail(ctx context.Context, recipient string, subject string, body string) error {
	if recipient == "" {
		return fmt.Errorf("recipient email is required")
	}

	// Create request payload
	reqBody := ResendEmailRequest{
		From:    c.from,
		To:      []string{recipient},
		Subject: subject,
		Text:    body,
		HTML:    fmt.Sprintf("<p>%s</p>", body), // Simple HTML version
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create HTTP request
	req, err := http.NewRequestWithContext(ctx, "POST", resendAPIURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	// Send request
	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		// Try to read error response
		var errorBody bytes.Buffer
		errorBody.ReadFrom(resp.Body)
		
		slog.Error("Resend API error",
			"status", resp.StatusCode,
			"response", errorBody.String(),
		)
		
		return fmt.Errorf("resend API returned status %d: %s", resp.StatusCode, errorBody.String())
	}

	// Parse successful response
	var emailResp ResendEmailResponse
	if err := json.NewDecoder(resp.Body).Decode(&emailResp); err != nil {
		slog.Warn("Failed to decode Resend response", "error", err)
		// Don't fail if we can't decode response, email was likely sent
	}

	slog.Info("Email sent successfully",
		"recipient", recipient,
		"subject", subject,
		"email_id", emailResp.ID,
	)

	return nil
}


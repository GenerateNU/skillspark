package main

import (
	"context"
	"fmt"
	"log/slog"
)

// NotificationProcessor processes notification messages
type NotificationProcessor struct {
	resendClient *ResendClient
	expoClient   *ExpoClient
}

// NewNotificationProcessor creates a new notification processor
func NewNotificationProcessor(resendClient *ResendClient, expoClient *ExpoClient) *NotificationProcessor {
	return &NotificationProcessor{
		resendClient: resendClient,
		expoClient:   expoClient,
	}
}

// ProcessNotification processes a single notification message
func (p *NotificationProcessor) ProcessNotification(ctx context.Context, message NotificationMessage) error {
	// Validate notification type
	if message.NotificationType != NotificationTypeEmail &&
		message.NotificationType != NotificationTypePush &&
		message.NotificationType != NotificationTypeBoth {
		return fmt.Errorf("invalid notification type: %s", message.NotificationType)
	}

	// Validate body is not empty
	if message.Body == "" {
		return fmt.Errorf("notification body is required")
	}

	// Determine subject for email notifications
	subject := "Notification"
	if message.Subject != nil && *message.Subject != "" {
		subject = *message.Subject
	}

	// Process email notification
	if message.NotificationType == NotificationTypeEmail || message.NotificationType == NotificationTypeBoth {
		if message.RecipientEmail == nil || *message.RecipientEmail == "" {
			return fmt.Errorf("recipient email is required for email notification")
		}

		if err := p.resendClient.SendEmail(ctx, *message.RecipientEmail, subject, message.Body); err != nil {
			return fmt.Errorf("failed to send email: %w", err)
		}

		slog.Info("Email notification sent",
			"recipient", *message.RecipientEmail,
			"subject", subject,
		)
	}

	// Process push notification
	if message.NotificationType == NotificationTypePush || message.NotificationType == NotificationTypeBoth {
		if message.RecipientPushToken == nil || *message.RecipientPushToken == "" {
			return fmt.Errorf("recipient push token is required for push notification")
		}

		if err := p.expoClient.SendPushNotification(ctx, *message.RecipientPushToken, message.Body, message.Metadata); err != nil {
			return fmt.Errorf("failed to send push notification: %w", err)
		}

		slog.Info("Push notification sent",
			"token", *message.RecipientPushToken,
		)
	}

	return nil
}


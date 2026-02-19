package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	// Initialize logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	// Initialize Resend client
	resendClient, err := NewResendClient()
	if err != nil {
		slog.Error("Failed to initialize Resend client", "error", err)
		os.Exit(1)
	}

	// Initialize Expo client (optional access token)
	expoClient := NewExpoClient()

	// Initialize notification processor
	processor := NewNotificationProcessor(resendClient, expoClient)

	// Initialize handler
	handler := NewHandler(processor)

	// Wrap handler for Lambda runtime
	lambdaHandler := func(ctx context.Context, event SQSEvent) (SQSEventResponse, error) {
		return handler.Handle(ctx, event)
	}

	// Start Lambda runtime
	slog.Info("Starting Lambda function")
	lambda.Start(lambdaHandler)
}


#!/bin/bash

# Build script for Lambda function
# This builds the Go binary for Linux/AMD64 architecture

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
LAMBDA_DIR="${SCRIPT_DIR}/lambda"
OUTPUT_DIR="${SCRIPT_DIR}/lambda"
BINARY_NAME="bootstrap"

echo "Building Lambda function..."
echo "Source directory: ${LAMBDA_DIR}"
echo "Output binary: ${OUTPUT_DIR}/${BINARY_NAME}"

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed. Please install Go to build the Lambda function."
    exit 1
fi

# Change to Lambda directory
cd "${LAMBDA_DIR}"

# Download dependencies
echo "Downloading dependencies..."
go mod download

# Build for Linux/AMD64 (required for Lambda)
echo "Building binary for Linux/AMD64..."
GOOS=linux GOARCH=amd64 go build -o "${OUTPUT_DIR}/${BINARY_NAME}" .

if [ $? -eq 0 ]; then
    echo "✓ Build successful!"
    echo "Binary location: ${OUTPUT_DIR}/${BINARY_NAME}"
    ls -lh "${OUTPUT_DIR}/${BINARY_NAME}"
else
    echo "✗ Build failed!"
    exit 1
fi


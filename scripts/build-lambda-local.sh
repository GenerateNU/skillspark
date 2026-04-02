#!/bin/sh
# Builds the notification Lambda binary for LocalStack.
# Runs via: docker run golang:1.24-alpine (see make up-localstack)
# /src   = infra/modules/main/lambda  (read-only)
# /output = infra/modules/main        (writes lambda_function.zip here)
set -e

apk add --no-cache zip > /dev/null 2>&1

echo "Downloading Lambda dependencies..."
cd /src
go mod download

echo "Building Lambda binary (linux/amd64)..."
GOOS=linux GOARCH=amd64 go build -o /tmp/bootstrap .

echo "Zipping Lambda binary..."
zip /output/lambda_function.zip -j /tmp/bootstrap

echo "Lambda build complete: /output/lambda_function.zip ($(du -sh /output/lambda_function.zip | cut -f1))"

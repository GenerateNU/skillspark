#!/bin/bash

# Set PATH to include common locations
export PATH="/opt/homebrew/bin:/usr/local/bin:$(go env GOPATH)/bin:$PATH"

echo "🔍 Running pre-commit checks..."

# Check for staged Go files
STAGED_GO_FILES=$(git diff --cached --name-only --diff-filter=ACM | grep '\.go$')

# Check for staged frontend files
STAGED_FRONTEND_FILES=$(git diff --cached --name-only --diff-filter=ACM | grep -E '^frontend/web/.*\.(ts|tsx|js|jsx)$')

# Run Go checks if there are staged Go files
if [ -n "$STAGED_GO_FILES" ]; then
  echo "📝 Running go fmt..."
  gofmt -w $STAGED_GO_FILES
  git add $STAGED_GO_FILES

  # Change to backend directory for go commands
  cd backend || exit 1

  # Check and install golangci-lint if needed
  GOLANGCI_LINT_VERSION="v1.64.2"
  if ! golangci-lint version 2>&1 | grep -q "$GOLANGCI_LINT_VERSION"; then
    echo "📦 Installing golangci-lint $GOLANGCI_LINT_VERSION..."
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@$GOLANGCI_LINT_VERSION
  fi

  # Run golangci-lint
  echo "🔍 Running golangci-lint..."
  if ! command -v golangci-lint &> /dev/null; then
    echo "⚠️  golangci-lint not found. Install it with:"
    echo "   brew install golangci-lint"
    echo "   OR go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"
    exit 1
  fi

  golangci-lint run
  if [ $? -ne 0 ]; then
    echo "❌ golangci-lint failed"
    exit 1
  fi

  # Run go mod tidy
  echo "🧹 Running go mod tidy..."
  go mod tidy
  git add go.mod go.sum
  
  cd ..
else
  echo "✅ No Go files to check"
fi

# Run frontend checks if there are staged frontend files
if [ -n "$STAGED_FRONTEND_FILES" ]; then
  echo "🎨 Running frontend linting..."
  cd frontend/web || exit 1
  
  # Run ESLint if configured
  if [ -f "package.json" ] && grep -q "\"lint\"" package.json; then
    bun run lint
    if [ $? -ne 0 ]; then
      echo "❌ Frontend linting failed"
      exit 1
    fi
  fi
  
  cd ..
else
  echo "✅ No frontend files to check"
fi

echo "✅ All checks passed!"
exit 0
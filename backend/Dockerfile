# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the executable
# -o specifies output location
# CGO_ENABLED=0 creates a static binary (better for alpine)
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/skillspark ./cmd

# Run stage
FROM alpine:latest

WORKDIR /app

# Copy only the compiled binary from builder stage
COPY --from=builder /app/bin/skillspark .

# Expose port if needed (adjust based on your app)
# EXPOSE 8080

# Run the executable
CMD ["./skillspark"]

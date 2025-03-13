# Use the official lightweight Go image
FROM golang:1.22.2 AS builder

WORKDIR /app

# Copy the Go module files and download dependencies
COPY . .
WORKDIR /app/site
RUN go mod tidy

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./site/cmd/main.go

# Use a minimal runtime image
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/site/server /app/site/server
COPY --from=builder /app/site/db_config.yaml /app/site/db_config.yaml

# Expose port 8080 (or the default port you set)
EXPOSE 8080

# Run the binary with default flags (can be overridden)
CMD ["/app/site/server", "-p", "8080", "-d", "/app/site/db_config.yaml"]



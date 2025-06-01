# Build stage
FROM golang:1.21.13-alpine AS builder

# Add build labels
LABEL maintainer="Arpit Srivastava <arpitsrivastava529@gmail.com>"
LABEL description="Go Backend Service"

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage
FROM alpine:3.19

# Add runtime labels
LABEL maintainer="Arpit Srivastava <arpitsrivastava529@gmail.com>"
LABEL description="Go Backend Service"

# Create non-root user
RUN adduser -D -g '' appuser

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/main .

# Copy .env file if it exists
RUN --mount=type=bind,from=builder,target=/builder \
    if [ -f /builder/.env ]; then cp /builder/.env .; fi

# Set proper permissions
RUN chown -R appuser:appuser /app

# Switch to non-root user
USER appuser

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["./main"]
# done

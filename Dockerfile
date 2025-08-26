# Build stage
FROM golang:1.24-alpine AS builder

# Install git and ca-certificates (needed for go modules)
RUN apk add --no-cache git ca-certificates

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/nadia

# Final stage
FROM alpine:latest

# Install ca-certificates and sqlite
RUN apk --no-cache add ca-certificates sqlite

# Set working directory
WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Copy static files
COPY --from=builder /app/*.html .
COPY --from=builder /app/favicon.ico .

# Create directory for database
RUN mkdir -p /root/data

# Expose port
EXPOSE 8080

# Set environment variables
ENV ENVIRONMENT=production
ENV DB_PATH=/root/data/nadia_transactions.db

# Run the application
CMD ["./main"]
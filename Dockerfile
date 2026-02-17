# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go mod files (if you have them)
COPY go.mod go.sum* ./

# Download dependencies
RUN go mod download || true

# Copy source code
COPY . .

# Build the application from cmd/web directory
RUN go build -o main ./cmd/web

# Runtime stage
FROM alpine:latest

WORKDIR /app

# Copy the binary
COPY --from=builder /app/main .

# Copy templates directory
COPY templates/ /app/templates/

EXPOSE 8080

CMD ["./main"]

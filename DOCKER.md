# Docker Best Practices

This file provides guidance for building Dockerfiles in this repository.

## Multi-Stage Builds

Use multi-stage builds to minimize final image size:

```dockerfile
# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /build
COPY . .
RUN go build -o app main.go

# Runtime stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /build/app .
CMD ["./app"]
```

## Base Image Selection

- Use specific version tags (e.g., `golang:1.21-alpine`) instead of `latest`
- Prefer Alpine Linux variants for smaller image sizes
- Use distroless images for enhanced security when possible

## Layer Optimization

- Order instructions from least to most frequently changing
- Combine RUN commands with `&&` to reduce layers
- Copy dependency files before source code to leverage cache:

```dockerfile
# Good: dependencies cached separately
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build
```

## Security

- Run containers as non-root user:
```dockerfile
RUN adduser -D -u 1001 appuser
USER appuser
```

- Scan images for vulnerabilities using `docker scan` or `trivy`
- Don't include secrets in image layers; use build arguments or runtime secrets

## Go-Specific Optimizations

- Use `CGO_ENABLED=0` for static binaries:
```dockerfile
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o app
```

- The `-ldflags="-w -s"` flags strip debug information for smaller binaries

## Image Tagging

- Tag images with semantic versions and git commit SHAs
- Use meaningful tags: `nova:v1.2.3`, `nova:abc123`, `nova:latest`

## .dockerignore

Create a `.dockerignore` file to exclude unnecessary files:
```
.git
*.md
.gitignore
.dockerignore
```

## Health Checks

Add health checks for container orchestration:
```dockerfile
HEALTHCHECK --interval=30s --timeout=3s \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1
```

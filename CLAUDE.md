# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Nova is a simple Go web service that serves a "Hello, World!" HTTP endpoint on port 8080. The project is containerized using Docker with multi-stage builds for optimized image size.

## Build and Run Commands

### Local Development

Build and run locally:
```bash
go build -o main .
./main
```

Run without building:
```bash
go run main.go
```

### Docker

Build Docker image locally:
```bash
./build.sh
```

Build and push multi-architecture image (ARM64 + AMD64) to registry:
```bash
export DOCKER_USERNAME=your-dockerhub-username
./build.sh --push
```

Or manually:
```bash
docker build -t nova .
```

Run the container:
```bash
docker run -p 8080:8080 nova
```

### Kubernetes (k3s on Raspberry Pi)

Deploy to Kubernetes cluster:
```bash
# 1. Build and push multi-arch image to registry
export DOCKER_USERNAME=your-dockerhub-username
./build.sh --push

# 2. Update k8s/deployment.yaml with your image name

# 3. Deploy to cluster
kubectl apply -k k8s/

# 4. Verify deployment
kubectl get pods -l app=nova
kubectl get svc nova
```

Access the service (internal):
```bash
# Port-forward for testing
kubectl port-forward svc/nova 8080:80

# Or from within cluster
curl http://nova.default.svc.cluster.local
```

For external access, configure your CloudFlare Tunnel to point to `http://nova.default.svc.cluster.local:80`

Delete deployment:
```bash
kubectl delete -k k8s/
```

### Testing the Service

Once running (locally or in Docker), test the endpoint:
```bash
curl http://localhost:8080
```

Expected response: `Hello, World!`

## Architecture

This is a single-file Go application (`main.go`) with:
- One HTTP handler serving "/" endpoint
- Server running on port 8080
- Uses only Go standard library (`net/http`)

The Dockerfile implements a multi-stage build pattern:
1. **Build stage**: Uses `golang:1.21-alpine` to compile the Go binary
2. **Runtime stage**: Uses `alpine:latest` with only the compiled binary for minimal image size

Kubernetes manifests in `k8s/`:
- **deployment.yaml**: Runs 2 replicas with health checks, optimized resource limits for Raspberry Pi
- **service.yaml**: ClusterIP service exposing port 80 (maps to container port 8080)
- **kustomization.yaml**: Kustomize configuration for easy deployment management

The build script supports multi-architecture builds (AMD64 + ARM64) for compatibility with Raspberry Pi clusters.

## Docker Best Practices

When modifying Dockerfiles in this repository, refer to `DOCKER.md` for detailed guidance on:
- Multi-stage builds for size optimization
- Base image selection (prefer Alpine variants)
- Layer caching strategies
- Security practices (non-root users, vulnerability scanning)
- Go-specific optimizations (CGO_ENABLED=0, build flags)
- Health checks for container orchestration

Key optimization for Go builds:
```dockerfile
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o app
```

## Dependencies

Go version: 1.25.6 (specified in go.mod)
Module path: github.com/cpwu/nova

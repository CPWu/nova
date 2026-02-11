# Nova

A lightweight Go web service that serves a simple HTTP endpoint. Designed for containerized deployment on Kubernetes clusters, with optimized builds for ARM64 (Raspberry Pi) and AMD64 architectures.

## Features

- Minimal HTTP server using Go standard library
- Multi-stage Docker builds for optimized image size
- Multi-architecture support (ARM64 + AMD64)
- Kubernetes-ready with health checks and resource limits
- Optimized for k3s on Raspberry Pi clusters

## Prerequisites

- Go 1.23 or later
- Docker (with buildx for multi-arch builds)
- kubectl (for Kubernetes deployment)
- A Kubernetes cluster (tested on k3s on Raspberry Pi)
- Docker Hub account (for registry push)

## Quick Start

### Local Development

Run directly with Go:
```bash
go run main.go
```

Or build and run:
```bash
go build -o main .
./main
```

Test the endpoint:
```bash
curl http://localhost:8080
# Expected: Hello, World!
```

### Docker

Build and run locally:
```bash
./build.sh
docker run -p 8080:8080 nova
```

Build and push to registry (multi-arch):
```bash
export DOCKER_USERNAME=your-dockerhub-username
docker login
./build.sh --push
```

## Kubernetes Deployment

### Prerequisites
1. Build and push the image (deployment already configured for chunw208/nova):
   ```bash
   export DOCKER_USERNAME=your-username
   ./build.sh --push
   ```

### Deploy
```bash
kubectl apply -k k8s/
```

### Verify
```bash
kubectl get pods -l app=nova
kubectl get svc nova
```

### Test
```bash
# Port-forward for local testing
kubectl port-forward svc/nova 8080:80
curl http://localhost:8080

# Or test from within cluster
kubectl run -it --rm debug --image=alpine --restart=Never -- wget -qO- http://nova
```

### Cleanup
```bash
kubectl delete -k k8s/
```

## Project Structure

```
nova/
├── main.go              # Go web server
├── go.mod               # Go module definition
├── Dockerfile           # Multi-stage Docker build
├── build.sh             # Build script with multi-arch support
├── DOCKER.md            # Docker best practices guide
├── CLAUDE.md            # AI assistant guidance
├── k8s/                 # Kubernetes manifests
│   ├── deployment.yaml  # Deployment configuration
│   ├── service.yaml     # Service definition
│   ├── kustomization.yaml
│   └── README.md
└── README.md            # This file
```

## Configuration

### Resource Limits
Default limits are optimized for Raspberry Pi:
- Memory: 32Mi request, 64Mi limit
- CPU: 50m request, 100m limit

Adjust in `k8s/deployment.yaml` if needed.

### Replicas
Default: 2 replicas for high availability

Scale manually:
```bash
kubectl scale deployment nova --replicas=3
```

## Development

See [CLAUDE.md](CLAUDE.md) for detailed development instructions and architecture overview.

See [DOCKER.md](DOCKER.md) for Docker best practices and optimization techniques.

## External Access

The service uses a ClusterIP service by default. For external access:

- **CloudFlare Tunnel**: Point to `http://nova.default.svc.cluster.local:80`
- **Ingress**: Add an ingress resource
- **LoadBalancer**: Change service type in `k8s/service.yaml`
- **NodePort**: Change service type for direct node access

## Contributing

This is a personal project. Feel free to fork and modify for your own use.

## License

MIT

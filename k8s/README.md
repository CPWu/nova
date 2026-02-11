# Kubernetes Manifests

This directory contains Kubernetes manifests for deploying Nova to a k3s cluster.

## Files

- **deployment.yaml**: Deployment with 2 replicas, health checks, and resource limits optimized for Raspberry Pi
- **service.yaml**: ClusterIP service for internal cluster access
- **kustomization.yaml**: Kustomize configuration for managing deployments

## Quick Deploy

```bash
# Apply all manifests
kubectl apply -k .

# Or from project root
kubectl apply -k k8s/
```

## Configuration

### Update Image

Edit `deployment.yaml` and replace `YOUR_DOCKERHUB_USERNAME` with your Docker Hub username:

```yaml
image: YOUR_DOCKERHUB_USERNAME/nova:latest
```

Or use environment variable substitution with kustomize:

```bash
cd k8s
kustomize edit set image nova=yourusername/nova:v1.0.0
kubectl apply -k .
```

### Adjust Replicas

```bash
kubectl scale deployment nova --replicas=3
```

### Resource Limits

Current settings are optimized for Raspberry Pi:
- Memory: 32Mi request, 64Mi limit
- CPU: 50m request, 100m limit

Adjust in `deployment.yaml` if needed.

## Verification

```bash
# Check pods
kubectl get pods -l app=nova

# Check service
kubectl get svc nova

# View logs
kubectl logs -l app=nova --tail=50

# Test service internally
kubectl run -it --rm debug --image=alpine --restart=Never -- wget -qO- http://nova
```

## Cleanup

```bash
kubectl delete -k .
```
